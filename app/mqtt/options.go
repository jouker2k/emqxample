package mqtt

import "time"

const timeout = time.Second * 30
const qos = 2
const retained = false
const keepAlive = 3 * time.Second

// ClientOpts
type ClientOpts struct {
	brokerURL  string
	password   string
	username   string
	clientID   string
	apiVer     string
	handlerQos byte
	retained   bool
	timeout    time.Duration
	keepAlive  time.Duration
}

// NewClientOpts returns a new ClientOpts struct with default values
func NewClientOpts() *ClientOpts {
	return &ClientOpts{
		brokerURL:  "",
		password:   "",
		username:   "",
		apiVer:     "",
		handlerQos: qos,
		retained:   retained,
		timeout:    timeout,
		keepAlive:  keepAlive,
	}
}

// SetAPIVer for mqtt comms
func (o *ClientOpts) SetAPIVer(ver string) *ClientOpts {
	o.apiVer = ver
	return o
}

// SetBrokerURL for the service
func (o *ClientOpts) SetBrokerURL(url string) *ClientOpts {
	o.brokerURL = url
	return o
}

// SetPassword for client used to connect to broker
func (o *ClientOpts) SetPassword(pass string) *ClientOpts {
	o.password = pass
	return o
}

// SetClientID used to connect with the broker
func (o *ClientOpts) SetClientID(clientID string) *ClientOpts {
	o.clientID = clientID
	return o
}

// SetHandlerQos for the service handlers this value is only used for service handlers
func (o *ClientOpts) SetHandlerQos(qos byte) *ClientOpts {
	o.handlerQos = qos
	return o
}

//SetUsername for the service connection to the broker
func (o *ClientOpts) SetUsername(u string) *ClientOpts {
	o.username = u
	return o
}
