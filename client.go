package mongodb

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"time"
)

type Client struct {
	Timeout time.Duration
	ctx     context.Context
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

func (c *Client) Connect() {

}
