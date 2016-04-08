package main

import (
	"github.com/spf13/viper"
	"log"
	"github.com/aroslov/abc/util"
	"github.com/aroslov/abc/xray"
)

func main() {
	util.ReadConfig("consumer_config")

	rabbit_endpoint := viper.GetString("rabbit_endpoint")
	rabbit_exchange := viper.GetString("rabbit_exchange")
	conn, ch := util.RabbitMQConnect(rabbit_endpoint, rabbit_exchange)
	defer conn.Close()
	defer ch.Close()

	fetch_count := viper.GetInt("fetch_count")

	ch.Qos(fetch_count, 0, false)
	msgs := util.RabbitMQGetMessages(rabbit_exchange, ch)

	xray := xray.GetJournalService(viper.GetString("xray_base"))
	journal_id := viper.GetString("journal_id")
	check_timestamp := viper.GetBool("check_timestamp")
	count := 0
	messages := make([]([]byte), fetch_count)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			messages[count] = d.Body
			count++
			log.Printf("Messages received: %d", count)
			if (count == fetch_count) {
				log.Printf("Processing the batch")
				processMessages(xray, journal_id, messages, check_timestamp)
				d.Ack(true)
				log.Printf("Ack")
				count = 0
			}

		}
	}()
	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
}

func processMessages(xray *xray.XRayService, journal_id string, messages []([]byte), check_timestamp bool) {
	for _, message := range messages {
		record_id := xray.CreateRecord(message)
		log.Printf("Created record ID %s", record_id)
		fp := xray.NewFingerprint(message)
		fp.CreateRecordFingerprint(record_id)
		log.Printf("Created record ID %s fingerprint", record_id)
		xray.CommitJournalRecord(journal_id, record_id)
		log.Printf("Commited record ID %s to journal", record_id)
	}
	if check_timestamp {
		j := xray.GetJournal(journal_id)
		if !j.HasTimestamp() {
			util.Fail("No timestamp")
		}
	}
}

