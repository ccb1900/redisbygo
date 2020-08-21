package log

import (
	log "github.com/sirupsen/logrus"
)

type Log struct {
}

func NewLog() *Log {
	customFormatter := new(log.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "1970-01-01 15:04:05"
	log.SetFormatter(customFormatter)
	return &Log{}
}

func (l *Log) Info(message string) {
	log.Info(message)
	//log.Warn(message)
}
func (l *Log) Log(message string, level int) {
	//log.Info(message)
	log.Warn(message)
}
