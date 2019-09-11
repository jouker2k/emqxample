package main

import (
	"net/http"
	"os"
	"strconv"

	"github.com/jwtea/emqxample/app/mqtt"

	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

// Specification struct to hold app env config
type Specification struct {
	APPURL       string `env:"APPURL"`
	MQTTAPIVer   string `env:"MQTT_API_VER"`
	MQTTDBHost   string `env:"MQTTDB_HOST"`
	MQTTDBPort   string `env:"MQTTDB_PORT"`
	MQTTDBUser   string `env:"MQTTDB_USER"`
	MQTTDBName   string `env:"MQTTDB_NAME"`
	MQTTDBPass   string `env:"MQTTDB_PASS"`
	MQTTURL      string `env:"MQTT_URL"`
	MQTTClientID string `env:"MQTT_CLIENT_ID"`
	MQTTPass     string `env:"MQTT_PASS"`
	MQTTUser     string `env:"MQTT_USER"`
	MigrateDB    bool   `env:"APP_MIGRATE_DB"`
}

func NewSpec() *Specification {
	migrate, err := strconv.ParseBool(getenv("APP_DB_MIGRATE", "true"))
	if err != nil {
		log.Fatalf("Cannot parse APP_DB_MIGRATE env var %e", err)
	}

	return &Specification{
		APPURL:       getenv("APP_URL", "0.0.0.0:8111"),
		MQTTDBHost:   getenv("MQTTDB_HOST", "127.0.0.1"),
		MQTTDBPort:   getenv("MQTTDB_PORT", "5432"),
		MQTTDBUser:   getenv("MQTTDB_USER", "docker"),
		MQTTDBName:   getenv("MQTTDB_NAME", "mqtt_db"),
		MQTTDBPass:   getenv("MQTTDB_PASS", "docker"),
		MQTTURL:      getenv("MQTT_URL", "ssl://192.168.99.100:32335"),
		MQTTClientID: getenv("MQTT_CLIENT_ID", "adminclient"),
		MQTTPass:     getenv("MQTT_PASS", "pass"),
		MQTTUser:     getenv("MQTT_USER", "admin"),
		MQTTAPIVer:   "/v1.0",
		MigrateDB:    migrate,
	}
}

func getenv(key, fallback string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return fallback
	}
	return value
}

func main() {
	log.SetLevel(log.DebugLevel)

	s := NewSpec()

	mqtt.NewStore(s.MQTTDBHost, s.MQTTDBPort, s.MQTTDBUser, s.MQTTDBName, s.MQTTDBPass, s.MigrateDB)

	mOpts := mqtt.NewClientOpts().SetAPIVer(s.MQTTAPIVer).SetBrokerURL(s.MQTTURL).SetUsername(s.MQTTUser).SetPassword(s.MQTTPass).SetClientID(s.MQTTClientID)
	mqtt.NewClient(mOpts).SetupHandlers()

	log.Info("MQTT client setup")

	log.Fatal(http.ListenAndServe(s.APPURL, nil))
}
