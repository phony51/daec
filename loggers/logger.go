package loggers

import (
	log "github.com/sirupsen/logrus"
	"io"
)

func ConfigureLogger(out io.Writer, fullTimestamp bool, level log.Level) {
	log.SetOutput(out)
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: fullTimestamp,
	})
	log.SetLevel(level)
}
