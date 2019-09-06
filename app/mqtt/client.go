package mqtt

import (
	pmqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

type Client struct {
	c    pmqtt.Client
	opts ClientOpts
}

type JSONMessage struct {
	Message string `json: "message"`
}

// NewClient returns a new client connection
func NewClient(o *ClientOpts) *Client {
	setupLoggers()

	opts := pmqtt.NewClientOptions().AddBroker(
		o.brokerURL).SetPassword(o.password).SetUsername(
		o.username).SetClientID(o.clientID).SetKeepAlive(o.keepAlive)

	opts.OnConnectionLost = DefConHandler

	c := pmqtt.NewClient(opts)

	if token := c.Connect(); token.Wait() && token.Error() != nil {
		log.Fatalf("MQTT Connect failed: %v \n", token.Error())
	}

	return &Client{c, *o}
}

// DefConHandler for lost connections
func DefConHandler(client pmqtt.Client, reason error) {
	log.Fatalf("MQTT lost connection: %v \n", reason.Error())
}
