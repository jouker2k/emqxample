package mqtt

import "time"

const timeout = time.Second * 30
const qos = 2
const retained = false
const keepAlive = 3 * time.Second

type LastWillOpts struct {
	topic    string
	message  string
	qos      byte
	retained bool
}

// ClientOpts
type ClientOpts struct {
	enableLog    bool
	enableTLS    bool
	caPath       string // file containing concatenated CA certificates if there is more than 1 in the chain.
	brokerURL    string
	password     string
	username     string
	clientID     string
	apiVer       string
	lastWillOpts *LastWillOpts
	handlerQos   byte
	retained     bool
	timeout      time.Duration
	keepAlive    time.Duration
}

// NewClientOpts returns a new ClientOpts struct with default values
func NewClientOpts() *ClientOpts {
	return &ClientOpts{
		enableLog:    false,
		enableTLS:    true,
		caPath:       "../certs/emqxca.pem",
		brokerURL:    "",
		password:     "",
		username:     "",
		apiVer:       "",
		lastWillOpts: nil,
		handlerQos:   qos,
		retained:     retained,
		timeout:      timeout,
		keepAlive:    keepAlive,
	}
}

// SetAPIVer for mqtt comms
func (o *ClientOpts) SetAPIVer(ver string) *ClientOpts {
	o.apiVer = ver
	return o
}

// EnableLogging configures mqtt client to output mqtt logs
func (o *ClientOpts) EnableLogging() *ClientOpts {
	o.enableLog = true
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

//SetLastWillOpts for the client connection to broker
func (o *ClientOpts) SetLastWillOpts(lw *LastWillOpts) *ClientOpts {
	o.lastWillOpts = lw
	return o
}

//NewLastWillOpts returns last will options with defaults
func NewLastWillOpts() *LastWillOpts {
	return &LastWillOpts{
		message:  "",
		topic:    "",
		qos:      qos,
		retained: retained,
	}
}

//SetTopic for the last will
func (lw *LastWillOpts) SetTopic(t string) *LastWillOpts {
	lw.topic = t
	return lw
}

//SetMessage for last will
func (lw *LastWillOpts) SetMessage(m string) *LastWillOpts {
	lw.message = m
	return lw
}

//SetQos for last will
func (lw *LastWillOpts) SetQos(q byte) *LastWillOpts {
	lw.qos = q
	return lw
}

//SetRetained for last will
func (lw *LastWillOpts) SetRetained(r bool) *LastWillOpts {
	lw.retained = r
	return lw
}
