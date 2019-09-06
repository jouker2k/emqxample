package mqtt

import (
	"strings"
)

const deviceTopic = "device"

// Wildcard value represneting mqtt wildcard
const Wildcard = "#"

// DeviceTopic returns the device topic for a given serial
func DeviceTopic(ver string, serial string) string {
	s := []string{ver, deviceTopic, serial}
	return strings.Join(s, "/")
}

// AllDeviceTopic returns the topic required for access to all devices
func AllDeviceTopic(ver string) string {
	s := []string{ver, deviceTopic, Wildcard}
	return strings.Join(s, "/")
}
