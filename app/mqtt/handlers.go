package mqtt

import (
	pmqtt "github.com/eclipse/paho.mqtt.golang"
	log "github.com/sirupsen/logrus"
)

// SetupHandlers adds any sub handlers
func (c *Client) SetupHandlers() {
	c.SubscribeDeviceMessage("1")
}

// SubscribeDeviceMessage establish subscription for topic with serial affix
func (c *Client) SubscribeDeviceMessage(serial string) {
	tok := c.c.Subscribe(DeviceTopic(c.opts.apiVer, serial), c.opts.handlerQos,
		func(client pmqtt.Client, msg pmqtt.Message) {
			message := c.NewDeviceMessage().Decode(msg.Payload())
			log.Infof("JSON: %s", message.Message)
			LogMQTTMessage(msg)
		})

	if tok.WaitTimeout(c.opts.timeout) && tok.Error() != nil {
		log.Errorf("Failed device subscribe: %s action: %v", serial, tok.Error())
	}
}

//PublishDeviceMessage send a json encoded message to a topic with a serial affix
func (c *Client) PublishDeviceMessage(serial string, message string) {
	log.Infof("c.DeviceTopic(serial):%v", DeviceTopic(c.opts.apiVer, serial))
	log.Infof("c.NewDeviceMessage().JSON(): %v", c.NewDeviceMessage().JSON())
	token := c.c.Publish(DeviceTopic(c.opts.apiVer, serial), c.opts.handlerQos,
		c.opts.retained, c.NewDeviceMessage().JSON())

	if token.WaitTimeout(c.opts.timeout) && token.Error() != nil {
		log.Errorf("Failed device publish:%s action: %v", serial, token.Error())
	}
}
