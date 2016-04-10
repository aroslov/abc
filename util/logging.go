package util

import (
	log "github.com/Sirupsen/logrus"
	"os"
)

func SetupLogging(log_level_name, log_file_name string) {
	log.SetFormatter(&log.TextFormatter{FullTimestamp: true, DisableColors: true})
	log_level, err := log.ParseLevel(log_level_name)
	if (err != nil) {
		log.Warn("Unable to parse log level from config. Using default.")
	} else {
		log.SetLevel(log_level)
	}
	if (log_file_name != "") {
		f, err := os.OpenFile(log_file_name, os.O_WRONLY | os.O_CREATE, 0755)
		if (err != nil) {
			log.Warnf("Unable to create a log file for %s", log_file_name)
		} else {
			log.Infof("Logging to %s", log_file_name)
			log.SetOutput(f)
		}
	}
}

