package config

import "github.com/streadway/amqp"

type Queues struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func NewQueues(addr string) Queues {
	conn, err := amqp.Dial(addr)
	FailOnError(err, "Failed to connect to RabbitMQ")

	ch, err := conn.Channel()
	FailOnError(err, "Failed to open a channel")

	// task
	_, err = ch.QueueDeclare(
		QueueRequest, // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	// result
	_, err = ch.QueueDeclare(
		QueueRespond, // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	FailOnError(err, "Failed to declare a queue")

	return Queues{
		Conn: conn,
		Ch:   ch,
	}
}

func (q *Queues) Deconstruct() {
	q.Ch.Close()
	q.Conn.Close()
}
