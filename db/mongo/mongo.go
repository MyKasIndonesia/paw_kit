// Package mongo consists of mongodb client initialization function
package mongo

import (
	"time"

	pawctx "github.com/PAWSOME-INDONESIA/paw_kit/ctx"
	"github.com/pkg/errors"
	mgo "go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const (
	// DefaultTimeout default timeout for database action.
	DefaultTimeout = 10 * time.Second
)

// Client ...
type Client struct {
	*mgo.Client
}

// New connect to given uri and return its client.
func New(uri string) (*Client, error) {
	ctx, _ := pawctx.BgWithTimeout(DefaultTimeout)

	c, err := mgo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		return nil, errors.Wrap(err, "failed to connect mongodb")
	}

	ctx, _ = pawctx.BgWithTimeout(DefaultTimeout)
	if err := c.Ping(ctx, readpref.Primary()); err != nil {
		return nil, errors.Wrap(err, "failed to ping mongodb")
	}

	return &Client{c}, nil
}

// Check connection check.
func (mc *Client) Check() error {
	ctx, _ := pawctx.BgWithTimeout(DefaultTimeout)
	if err := mc.Ping(ctx, readpref.Primary()); err != nil {
		return errors.Wrap(err, "failed to ping mongodb")
	}
	return nil
}

// IsDupKeyErr ...
func IsDupKeyErr(err error) bool {
	var e mgo.WriteException
	if errors.As(err, &e) {
		for _, we := range e.WriteErrors {
			if we.Code == 11000 {
				return true
			}
		}
	}
	return false
}
