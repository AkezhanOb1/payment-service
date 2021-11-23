package job

import (
	"encoding/json"
	"os"

	"github.com/AkezhanOb1/payment/repositories"

	"github.com/AkezhanOb1/payment/models"

	"github.com/AkezhanOb1/payment/configs/log"
	"github.com/streadway/amqp"
)

func ConsumeFromPaymentResultJob() {
	var rabbitURI = os.Getenv("RabbitURI")
	var paymentResultQueue = os.Getenv("PaymentResultQueue")
	conn, err := amqp.Dial(rabbitURI)
	if err != nil {
		log.Logger.Warn(err)
		return
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Logger.Warn(err)
		return
	}
	defer ch.Close()

	messages, err := ch.Consume(
		paymentResultQueue, // queue
		"",                 // consumer
		true,               // auto-ack
		false,              // exclusive
		false,              // no-local
		false,              // no-wait
		nil,                // args
	)
	if err != nil {
		log.Logger.Warn(err)
		return
	}

	forever := make(chan bool)

	go func() {
		for message := range messages {
			var payment models.Payment
			json.Unmarshal(message.Body, &payment)
			err = repositories.SendResultToMerchantRepository(payment)
			if err != nil {
				log.Logger.Warn(err, payment)
			}
		}
	}()

	<-forever
}
