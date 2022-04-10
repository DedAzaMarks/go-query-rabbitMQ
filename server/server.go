package main

import (
	"bytes"
	"encoding/gob"
	"github.com/streadway/amqp"
	"hw05-query/config"
)

func main() {
	queues := config.NewQueues("amqp://guest:guest@localhost:5672/")

	msgs, err := queues.Ch.Consume(
		config.QueueRequest, // queue
		"",                  // consumer
		true,                // auto-ack
		false,               // exclusive
		false,               // no-local
		false,               // no-wait
		nil,                 // args
	)
	config.FailOnError(err, "Failed to register request consumer")
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	for d := range msgs {
		var req config.Request
		buf.Reset()
		buf.Write(d.Body)
		config.FailOnError(dec.Decode(&req), "can't decode request")

		length, path, err := bfs(&req)
		if err != nil {
			err = queues.Ch.Publish(
				"",                  // exchange
				config.QueueRespond, // routing key
				false,               // mandatory
				false,               // immediate
				amqp.Publishing{
					DeliveryMode: amqp.Persistent,
					ContentType:  "text/plain",
					Body:         []byte(err.Error()),
				},
			)
			config.FailOnError(err, "can't publish respond")
			continue
		}

		var resp config.Respond
		resp.Path = path
		resp.PathLength = length
		resp.Id = req.Id

		buf.Reset()
		config.FailOnError(enc.Encode(resp), "can't encoder respond")

		err = queues.Ch.Publish(
			"",                  // exchange
			config.QueueRespond, // routing key
			false,               // mandatory
			false,               // immediate
			amqp.Publishing{
				DeliveryMode: amqp.Persistent,
				ContentType:  "binary/gob",
				Body:         buf.Bytes(),
			},
		)
		config.FailOnError(err, "can't publish respond")
	}
}
