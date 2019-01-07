package mongodb

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"time"
)

// Collection ...
type Collection struct {
	ctx     context.Context
	Timeout time.Duration
	*mongo.Collection
}

// DeleteMany ...
func (c *Collection) DeleteMany(filter interface{},
	opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	timeout, _ := context.WithTimeout(c.ctx, c.Timeout)
	return c.Collection.DeleteMany(timeout, filter, opts...)
}

// InsertOne ...
func (c *Collection) InsertOne(document interface{},
	opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	//timeout, _ := context.WithTimeout(c.ctx, c.Timeout)
	return c.Collection.InsertOne(c.ctx, document, opts...)
}
