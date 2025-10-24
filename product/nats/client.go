package nats

import (
	"time"

	nats "github.com/nats-io/nats.go"
)

type NatsClient struct {
	conn *nats.Conn
}
func New() *NatsClient {
	return &NatsClient{}
}

func (c *NatsClient) Connect(url string) (error) {
	conn, err := nats.Connect(url, nats.MaxReconnects(-1), nats.ReconnectWait(2*time.Second))
	if err != nil {
		return err
	}
	c.conn = conn
	return nil
}

func (c *NatsClient) Disconnect() error {
	if c.conn != nil {
		c.conn.Close()
	}
	return nil
}

func (c *NatsClient) RegisterRPCHandler(subject string , queue string , handler func([]byte) ([]byte,error)) error  {
	_, err := c.conn.QueueSubscribe(subject, queue, func(msg *nats.Msg) {
		response, err := handler(msg.Data)
		if err != nil {
			// msg.Nak()
			response := []byte(err.Error())
			msg.Respond(response)
			return
		}
		msg.Respond(response)
	})
	return err	
}	