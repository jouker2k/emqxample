package mqtt

import (
	"encoding/json"
	"strings"

	log "github.com/sirupsen/logrus"
)

// Message interface
type Message interface {
	JSON() []byte
	Decode([]byte) struct{}
}

// DeviceMessage properties for device messages
type DeviceMessage struct {
	Message   string
	Timestamp int64
}

// NewDeviceMessage returns a new device message
func (c *Client) NewDeviceMessage() *DeviceMessage {
	return &DeviceMessage{"hello", 0}
}

// JSON returns json encoding of message or errors out
func (m *DeviceMessage) JSON() []byte {
	encMsg, err := json.Marshal(m)
	if err != nil {
		log.Errorf("Failed to marshal message:%v err:%v", m, err)
	}
	return encMsg
}

// Decode device message returning the struct or erroring out
func (m *DeviceMessage) Decode(payload []byte) *DeviceMessage {
	decoder := json.NewDecoder(strings.NewReader(string(payload)))
	if err := decoder.Decode(m); err != nil {
		log.Errorf("Failed to decode message payload: %v err: %v", payload, err)
	}
	return m
}
