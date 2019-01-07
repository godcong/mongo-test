package mongodb

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/mongo/options"
	"time"
)

type Database struct {
	ctx     context.Context
	Timeout time.Duration
	*mongo.Database
}

func (d *Database) Collection(name string, opts ...*options.CollectionOptions) *Collection {
	collection := d.Database.Collection(name, opts...)
	return &Collection{
		ctx:        d.ctx,
		Timeout:    d.Timeout,
		Collection: collection,
	}
}

