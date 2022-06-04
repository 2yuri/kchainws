package kchainws

import (
	"context"
	"encoding/json"
	"github.com/gorilla/websocket"
	"net/url"
	"sync"
)

//Delivery is the response from your ws
type Delivery chan []byte

//conn manage the connection and the readers
type conn struct {
	mu sync.Mutex

	deliveries []Delivery
	conn       *websocket.Conn
}

//add append new readers
func (c *conn) add(message Delivery) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.deliveries = append(c.deliveries, message)
}

//connect to ws
func (c *conn) connect(u url.URL, customs *Customization) error {
	conn, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		return err
	}

	b, err := json.Marshal(customs)
	if err != nil {
		return err
	}

	if err := conn.WriteMessage(websocket.TextMessage, b); err != nil {
		return err
	}

	c.conn = conn

	return nil
}

//close the conn and remove readers
func (c *conn) close() error {
	for _, d := range c.deliveries {
		for len(d) > 0 {
			<-d
		}
		close(d)
	}
	if c.conn == nil {
		return ErrNilClient
	}
	return c.conn.Close()
}

//GetReader return a reader to watch messages
func (c *conn) GetReader() Delivery {
	delivery := make(Delivery)
	c.add(delivery)

	return delivery
}

//read ws messages
func (c *conn) read(cancelFunc context.CancelFunc) {
	for {
		_, m, err := c.conn.ReadMessage()
		if err != nil {
			cancelFunc()
			return
		}

		c.mu.Lock()
		for _, d := range c.deliveries {
			d <- m
		}
		c.mu.Unlock()
	}
}
