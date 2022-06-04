package kchainws

import (
	"context"
	"net/url"
)

//WithURL setup the node host to Client
func WithURL(nodeHost string) Option {
	return func(c *Client) error {
		u := url.URL{Scheme: "ws", Host: nodeHost, Path: "/subscribe"}
		c.url = u
		return nil
	}
}

//WithContext setup a custom context to the client
func WithContext(ctx context.Context) Option {
	return func(c *Client) error {
		c.ctx = ctx
		return nil
	}
}

//WithCustom setup your customization pattern
func WithCustom(custom *Customization) Option {
	return func(c *Client) error {
		c.customization = custom
		return nil
	}
}
