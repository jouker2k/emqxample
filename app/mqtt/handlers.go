package mqtt

import (
	pmqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

// SetupHandlers adds any sub handlers
func (c *Client) SetupHandlers() {
	c.SubscribeAllDeviceMessages()
}

// SubscribeAllDeviceMessages establish wildcard subscription on all device topics
func (c *Client) SubscribeAllDeviceMessages() {
	tok := c.c.Subscribe(AllDeviceTopic(c.opts.apiVer), c.opts.handlerQos,
		func(client pmqtt.Client, msg pmqtt.Message) {
			message := NewDeviceMessage().Decode(msg.Payload())
			log.Infof("Topic: %s JSON: %s", msg.Topic(), message.Message)
			LogMQTTMessage(msg)
		})

	if tok.WaitTimeout(c.opts.timeout) && tok.Error() != nil {
		log.Errorf("Failed subscribe action: %v", tok.Error())
	}
}

// SubscribeDeviceMessage establish subscription for topic with serial affix
func (c *Client) SubscribeDeviceMessage(serial string) {
	tok := c.c.Subscribe(DeviceTopic(c.opts.apiVer, serial), c.opts.handlerQos,
		func(client pmqtt.Client, msg pmqtt.Message) {
			message := NewDeviceMessage().Decode(msg.Payload())
			log.Infof("JSON: %s", message.Message)
			LogMQTTMessage(msg)
		})

	if tok.WaitTimeout(c.opts.timeout) && tok.Error() != nil {
		log.Errorf("Failed device subscribe: %s action: %v", serial, tok.Error())
	}
}

//PublishDeviceMessage send a json encoded message to a topic with a serial affix
func (c *Client) PublishDeviceMessage(serial string, message string) {
	token := c.c.Publish(DeviceTopic(c.opts.apiVer, serial), c.opts.handlerQos,
		c.opts.retained, NewDeviceMessage().JSON())

	if token.WaitTimeout(c.opts.timeout) && token.Error() != nil {
		log.Errorf("Failed device publish:%s action: %v", serial, token.Error())
	}
}
