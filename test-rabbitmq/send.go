package main

import (
	"fmt"
	"log"

	"github.com/streadway/amqp"
)

func failOnError1(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func main() {

	Confamqp := amqp.URI{
		Scheme:   "amqp",
		//Host:     "13.250.228.156",
		Host:     "rabbitmq.pede.id",
		Port:     5672,
		Username: "loanApiSdk",
		Password: "92Lo4nSdK01",
		Vhost:    "/LoanApiSdk",
	}.String()
	//conn, err := amqp.Dial("amqp://loanApiSdk:92Lo4nSdK01@https://rabbitmq.pede.id")
	//conn, err := amqp.Dial("amqp://test:test@dev.pede.id:2222")
	fmt.Println(Confamqp)
	conn, err := amqp.Dial(Confamqp)
	failOnError1(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError1(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError1(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	log.Printf(" [x] Sent %s", body)
	failOnError1(err, "Failed to publish a message")
}
