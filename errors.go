package kchainws

import "errors"

var (
	ErrNilClient       = errors.New("client is nil")
	ErrInvalidURI      = errors.New("URI provided is invalid")
	ErrEmptyDeliveries = errors.New("empty deliveries")
	ErrInvalidType     = errors.New("type provided is invalid")
)
