package main

import (
	"github.com/spf13/viper"
	log "github.com/Sirupsen/logrus"
	"github.com/aroslov/abc/util"
	"github.com/aroslov/abc/xray"
	"time"
)

func main() {
	util.ReadConfig("consumer_config")
	util.SetupLogging(viper.GetString("log_level"), viper.GetString("log_output"))

	rabbit_endpoint := viper.GetString("rabbit_endpoint")
	rabbit_exchange := viper.GetString("rabbit_exchange")
	conn, ch := util.RabbitMQConnect(rabbit_endpoint, rabbit_exchange)
	defer conn.Close()
	defer ch.Close()

	fetch_count := viper.GetInt("fetch_count")

	ch.Qos(fetch_count, 0, false)
	msgs := util.RabbitMQGetMessages(rabbit_exchange, ch)

	xray_service := xray.GetXRayService(viper.GetString("xray_base"),
		viper.GetString("xray_access_key"), viper.GetString("xray_access_secret"))
	journal_id := viper.GetString("journal_id")
	check_timestamp := viper.GetBool("check_timestamp")
	count := 0
	messages := make([]([]byte), fetch_count)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			messages[count] = d.Body
			count++
			log.Debugf("Messages received: %d", count)
			if (count == fetch_count) {
				log.Infof("Processing the batch")
				processMessages(xray_service, journal_id, messages, check_timestamp)
				d.Ack(true)
				log.Printf("Ack")
				count = 0
			}

		}
	}()
	log.Infof("Waiting for messages. To exit press CTRL+C")
	<-forever
}

func processMessages(xray_service *xray.XRayService, journal_id string, messages []([]byte), check_timestamp bool) {
	for _, message := range messages {
		record := xray_service.CreateRecord(message)
		log.Infof("Created record ID %s", record.Id)
		fp := xray_service.NewFingerprint(message)
		fp.CreateRecordFingerprint(record.Id)
		log.Infof("Created record ID %s fingerprint", record.Id)
		record.CommitToJournal(journal_id)
		log.Infof("Commited record ID %s to journal", record.Id)
	}
	if check_timestamp {
		j := xray_service.GetJournal(journal_id)
		n1 := j.NumeberOfTimestamps()
		j.Timestamp()
		j = xray_service.GetJournal(journal_id)
		n2 := j.NumeberOfTimestamps()
		if (n2 != n1 + 1) {
			log.Warnf("Unable to confirm timestamp creation, retrying in 10 seconds")
			time.Sleep(time.Duration(10)*time.Second)
			j = xray_service.GetJournal(journal_id)
			n2 := j.NumeberOfTimestamps()
			if (n2 != n1 + 1) {
				log.Panicf("Unable to confirm timestamp creation after second attempt")
			}
		}
		log.Infof("Created timestamp, txid: %s", j.Timestamps[n1].Txid)
	}
}

