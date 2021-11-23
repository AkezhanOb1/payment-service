package repositories

import (
	"encoding/json"
	"os"

	"github.com/AkezhanOb1/payment/models"
	"github.com/streadway/amqp"
)

func PutPaymentResultInQueue(payment models.Payment) error {
	var rabbitURI = os.Getenv("RabbitURI")
	var paymentResultQueue = os.Getenv("PaymentResultQueue")

	conn, err := amqp.Dial(rabbitURI)
	if err != nil {
		return err
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		return err
	}
	defer ch.Close()

	body, _ := json.Marshal(payment)
	err = ch.Publish(
		"",                 // exchange
		paymentResultQueue, // routing key
		false,              // mandatory
		false,              // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(body),
		},
	)

	if err != nil {
		return err
	}

	return nil
}
