package mqtt

import (
	pmqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

type ErrorLogger struct {
}

func (*ErrorLogger) Println(v ...interface{}) {
	log.Error(v)
}
func (*ErrorLogger) Printf(format string, v ...interface{}) {
	log.Errorf(format, v)
}

type WarnLogger struct {
}

func (*WarnLogger) Println(v ...interface{}) {
	log.Warn(v)
}
func (*WarnLogger) Printf(format string, v ...interface{}) {
	log.Warnf(format, v)
}

type Debug struct {
}

func (*Debug) Println(v ...interface{}) {
	log.Info(v)
}
func (*Debug) Printf(format string, v ...interface{}) {
	log.Infof(format, v)
}

func setupLoggers() {
	pmqtt.ERROR = &ErrorLogger{}
	pmqtt.CRITICAL = &WarnLogger{}
	pmqtt.WARN = &WarnLogger{}
	pmqtt.DEBUG = &Debug{}
}

// LogMQTTMessage pretty print the contents of the message from the client
func LogMQTTMessage(m pmqtt.Message) {
	log.Infof("MSG RCV TOPIC: %s", m.Topic())
	log.Infof("-DUPLICATE: %v", m.Duplicate())
	log.Infof("-MESSAGE ID: %d", m.MessageID())
	log.Infof("-PAYLOAD: %s", m.Payload())
}
