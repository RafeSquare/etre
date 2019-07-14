// Copyright 2017, Square, Inc.

package mock

import (
	"errors"

	"github.com/globalsign/mgo"
)

var (
	ErrConnector = errors.New("error in connector")
)

type Connector struct {
	ConnectFunc func() (*mgo.Session, error)
	CloseFunc   func()
}

func (c *Connector) Init() error {
	return nil
}

func (c *Connector) Connect() (*mgo.Session, error) {
	if c.ConnectFunc != nil {
		return c.ConnectFunc()
	}
	return nil, nil
}

func (c *Connector) Close() {
	if c.CloseFunc != nil {
		c.CloseFunc()
	}
}
