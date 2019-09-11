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

	if o.enableLog {
		setupLoggers()
	}

	opts := pmqtt.NewClientOptions().AddBroker(
		o.brokerURL).SetPassword(o.password).SetUsername(
		o.username).SetClientID(o.clientID).SetKeepAlive(o.keepAlive)

	if o.enableTLS {
		opts.SetTLSConfig(NewTLSConfig(o.caPath))
	}

	opts.OnConnectionLost = DefConHandler

	if o.lastWillOpts != nil {
		opts.SetWill(o.lastWillOpts.topic,
			o.lastWillOpts.message, o.lastWillOpts.qos, o.lastWillOpts.retained)
	}

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
