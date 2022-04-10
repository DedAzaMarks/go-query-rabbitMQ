package main

import (
	"bytes"
	"encoding/gob"
	"github.com/google/uuid"
	"github.com/streadway/amqp"
	"hw05-query/config"
	"log"
	"net/http"
	"net/rpc"
	"sync"
)

func newPathFinder() *Pathfinder {
	res := Pathfinder{
		make(map[uuid.UUID]chan *config.Respond),
		config.NewQueues("amqp://guest:guest@localhost:5672/"),
		sync.RWMutex{},
	}
	go res.respondGetter()
	return &res
}

type Pathfinder struct {
	channels map[uuid.UUID]chan *config.Respond
	queues   config.Queues
	mu       sync.RWMutex
}

func (pf *Pathfinder) respondGetter() {
	messages, err := pf.queues.Ch.Consume(
		config.QueueRespond, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	config.FailOnError(err, "Failed to register respond consumer")
	var buf bytes.Buffer
	dec := gob.NewDecoder(&buf)
	for d := range messages {
		var resp config.Respond
		buf.Reset()
		buf.Write(d.Body)
		config.FailOnError(dec.Decode(&resp), "can't decode respond")
		pf.mu.Lock()
		pf.channels[resp.Id] <- &resp
		pf.mu.Unlock()
	}
}

func (pf *Pathfinder) GetPath(request *config.Request, respond *config.Respond) error {
	log.Printf("%+v", request)
	id := uuid.New()
	ch := make(chan *config.Respond)

	pf.mu.Lock()
	pf.channels[id] = ch
	pf.mu.Unlock()

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	request.Id = id
	config.FailOnError(enc.Encode(request), "can't encode request")

	err := pf.queues.Ch.Publish(
		"",                  // exchange
		config.QueueRequest, // routing key
		false,               // mandatory
		false,               // immediate
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "binary/gob",
			Body:         buf.Bytes(),
		},
	)
	config.FailOnError(err, "can't publish request")
	*respond = *<-ch
	log.Printf("%+v", respond)
	return nil
}

func main() {
	config.FailOnError(rpc.Register(newPathFinder()), "can't register pathfinder handler")
	rpc.HandleHTTP()
	config.FailOnError(http.ListenAndServe(config.ADDRESS, nil), "can't listen http")
}
