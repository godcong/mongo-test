package mongodb

import (
	"context"
	"github.com/mongodb/mongo-go-driver/mongo/readpref"
	"testing"
)

var client, err = NewClient(context.Background(), "mongodb://root:v2RgzSuIaBlx@localhost:27017")

// TestNewClient ...
func TestNewClient(t *testing.T) {
	err = client.Ping(context.Background(), readpref.Primary())
	t.Log(err)
}

// TestDatabase_Collection ...
func TestDatabase_Collection(t *testing.T) {

}
