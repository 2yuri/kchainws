package kchainws

import (
	"context"
	"errors"
	"net/url"
)

//Option is responsible for the client setup
type Option func(*Client) error

//Client ws client
type Client struct {
	ctx           context.Context
	url           url.URL
	customization *Customization

	*conn
}

//NewClient creates a new ws client
func NewClient(opts ...Option) (*Client, error) {
	c := &Client{conn: &conn{}}

	for _, opt := range opts {
		err := opt(c)
		if err != nil {
			return nil, errors.New("error applying option: " + err.Error())
		}
	}

	return c, nil
}

//StarConnection create a conn and return messages to readers
func (c *Client) StarConnection() error {
	if _, err := url.Parse(c.url.String()); err != nil {
		return ErrInvalidURI
	}
	if len(c.conn.deliveries) == 0 {
		return ErrEmptyDeliveries
	}
	if err := c.validateCustomization(); err != nil {
		return err
	}

	if c.ctx == nil {
		c.ctx = context.Background()
	}

	err := c.conn.connect(c.url, c.customization)
	if err != nil {
		return err
	}

	ctx, cancel := context.WithCancel(c.ctx)
	c.ctx = ctx

	go c.conn.read(cancel)

	return nil
}

func (c *Client) validateCustomization() error {
	if c.customization == nil {
		c.customization = &Customization{}
	}

	return c.customization.validate()
}

//Close conn and remove readers
func (c *Client) Close() error {
	return c.conn.close()
}

//Context return the context
func (c *Client) Context() context.Context {
	return c.ctx
}
