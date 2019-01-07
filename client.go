package mongodb

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"time"
)

type Client struct {
	ctx     context.Context
	Timeout time.Duration
	*mongo.Client
}

func NewClient(ctx context.Context, uri string) (*Client, error) {
	client, err := mongo.NewClient(uri)
	if err != nil {
		return nil, err
	}
	if ctx == nil {
		ctx = context.Background()
	}

	cli := Client{
		Timeout: 5 * time.Second,
		ctx:     ctx,
		Client:  client,
	}
	c, _ := context.WithTimeout(cli.ctx, cli.Timeout)
	err = client.Connect(c)
	if err != nil {
		return nil, err
	}
	return &cli, nil
}

func (c *Client) Reconnect() error {
	ctx, _ := context.WithTimeout(c.ctx, c.Timeout)
	return c.Client.Connect(ctx)
}

func (c *Client) Database(name string, opts ...*options.DatabaseOptions) *Database {
	database := c.Client.Database(name, opts...)
	return &Database{
		ctx:      c.ctx,
		Timeout:  c.Timeout,
		Database: database,
	}
}

func (c *Client) Context() context.Context {
	return c.ctx
}
