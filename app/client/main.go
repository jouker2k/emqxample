package main

import (
	"os"
	"time"

	"github.com/jwtea/emqxample/app/mqtt"
	log "github.com/sirupsen/logrus"
)

// Specification struct to hold app env config
type Specification struct {
	MQTTAPIVer   string `env:"MQTT_API_VER"`
	MQTTURL      string `env:"MQTT_URL"`
	MQTTClientID string `env:"MQTT_CLIENT_ID"`
	MQTTPass     string `env:"MQTT_PASS"`
	MQTTUser     string `env:"MQTT_USER"`
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func NewSpec() *Specification {
	return &Specification{
		MQTTURL:      getenv("MQTT_URL", "ssl://192.168.99.100:32335"),
		MQTTClientID: getenv("MQTT_CLIENT_ID", "clientclient"),
		MQTTPass:     getenv("MQTT_PASS", "pass2"),
		MQTTUser:     getenv("MQTT_USER", "client2"),
		MQTTAPIVer:   "/v1.0",
	}
}

func main() {

	serial := "2"

	log.SetLevel(log.DebugLevel)
	s := NewSpec()
	mOpts := mqtt.NewClientOpts().SetAPIVer(s.MQTTAPIVer).SetBrokerURL(s.MQTTURL).SetUsername(s.MQTTUser).SetPassword(s.MQTTPass).SetClientID(s.MQTTClientID)

	lwMessage := mqtt.DeviceMessage{Message: "Last will", Timestamp: 0}

	lwOpts := mqtt.NewLastWillOpts().SetTopic(mqtt.DeviceTopic(s.MQTTAPIVer, serial)).SetMessage(string(lwMessage.JSON()))

	mOpts.SetLastWillOpts(lwOpts)

	mc := mqtt.NewClient(mOpts)

	for i := 0; i < 2; i++ {
		mc.PublishDeviceMessage(serial, "yoi")
		time.Sleep(1000)
	}
}
