package main

import (
	"context"
	"log"
	"strconv"

	"github.com/resc/gobit/claims"
	"github.com/streadway/amqp"
)

func main() {
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"rpc_queue", // name
		false,       // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	failOnError(err, "Failed to declare a queue")

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	failOnError(err, "Failed to set QoS")

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			n, err := strconv.Atoi(string(d.Body))
			failOnError(err, "Failed to convert body to integer")

			log.Printf(" [.] fib(%d)", n)
			response := 0

			err = ch.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          []byte(strconv.Itoa(response)),
				})
			failOnError(err, "Failed to publish a message")

			d.Ack(false)
		}
	}()

	<-forever
}

func failOnError(err error, msg string) {
	panic(msg)
}

// concerns:
//   - statistics
//   - user tracking (amqp AppID to be used as `app user`)
//   - database transactions
//   - logging
//   - request tracing
//   -

type (

	// Headers contains the message metadata
	Headers interface {
		MessageId() MsgID
		CorrelationId() MsgID
		MimeType() string
	}

	// Body contains the raw bytes of the body
	Body []byte

	// MsgID is a client side generated uuid for the message
	MsgID [16]byte

	// MessageSender can send messages via rabbitmq.
	MessageSender interface {
		Send(headers Headers, body Body)
	}

	// MessageContext contains the processing environment for a message.
	MessageContext interface {
		context.Context
		MessageSender

		Sender() claims.Principal
		Headers() Headers
		Body() Body

		Commit() error
		Rollback() error
	}

	// MessageHandler is a marker interface for handlers
	MessageHandler interface {
		Handle(ctx MessageContext)
	}
)

type (
	// access_token
	contextKey int
)

const (
	// ReceiveTimestamp container the message receive timestamp
	ReceiveTimestamp contextKey = iota
)
