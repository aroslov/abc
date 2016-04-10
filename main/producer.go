package main

import (
	log "github.com/Sirupsen/logrus"
	"os"
	"github.com/streadway/amqp"
	"github.com/spf13/viper"
	"github.com/aroslov/abc/util"
	"strconv"
)

func main() {
	util.ReadConfig("generator_config")
	util.SetupLogging("info", "")
	number_of_messages := getNumberOfMessages()

	rabbit_endpoint := viper.GetString("rabbit_endpoint")
	rabbit_exchange := viper.GetString("rabbit_exchange")
	conn, ch := util.RabbitMQConnect(rabbit_endpoint, rabbit_exchange)
	defer conn.Close()
	defer ch.Close()

	template := util.LoadTemplate(viper.GetString("message_template_file"))

	log.Printf("Sending messages")
	for i := 0; i < number_of_messages; i++ {
		message := util.SubstituteTemplate(template);
		err := ch.Publish(
			rabbit_exchange, // exchange
			"", // routing key
			false, // mandatory
			false, // immediate
			amqp.Publishing{
				ContentType: "application/json",
				Body: message,
			})
		util.PanicOnError(err, "Failed to publish a message")
	}
	log.Printf("Sent %d messages", number_of_messages)
}

func getNumberOfMessages() int {
	if (len(os.Args) < 2) || os.Args[1] == "" {
		log.Panicf("Usage: go run producer.go <number of messages to generate>")
		return 0
	} else {
		n, err := strconv.Atoi(os.Args[1])
		util.PanicOnError(err, "Unable to parse the parameter")
		return n
	}
}

