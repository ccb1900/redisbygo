package log

import (
	"github.com/sirupsen/logrus"
	"os"
	"reflect"
)

type ILog interface {
	Trace(messages ...interface{})
	Debug(message ...interface{})
	Info(message ...interface{})
	Warn(message ...interface{})
	Error(message ...interface{})
	Fatal(message ...interface{})
}

type Log struct {
	name     string
	instance *logrus.Logger
}

func NewLog(name interface{}) ILog {
	l := new(Log)
	l.name = reflect.TypeOf(name).PkgPath()
	l.instance = logrus.New()
	l.init()
	return l
}
func (l *Log) init() {
	customFormatter := new(logrus.TextFormatter)
	customFormatter.FullTimestamp = true
	customFormatter.TimestampFormat = "2006-01-02 15:04:05"
	l.instance.SetFormatter(customFormatter)
	l.instance.SetOutput(os.Stdout)

}
func (l *Log) Info(messages ...interface{}) {
	l.log().Info(messages...)
}
func (l *Log) log() *logrus.Entry {
	return l.instance.WithField("package", l.name)
}
func (l *Log) Error(messages ...interface{}) {
	l.log().Error(messages...)
}
func (l *Log) Warn(messages ...interface{}) {
	l.log().Warn(messages...)
}
func (l *Log) Trace(messages ...interface{}) {
	l.log().Trace(messages...)
}
func (l *Log) Debug(messages ...interface{}) {
	l.log().Debug(messages...)
}
func (l *Log) Fatal(messages ...interface{}) {
	l.log().Fatal(messages...)
}
