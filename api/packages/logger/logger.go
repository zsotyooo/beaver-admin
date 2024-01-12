package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func Init() {
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)
	log.SetOutput(os.Stdout)
}

func Debug(fields map[string]interface{}, message string) {
	log.WithFields(fields).Debug(message)
}
