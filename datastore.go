// shippy-service-consignment/datastore.go
package main

import (
	"context"
	"log"
	"time"

	//	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// CreateClient -
func CreateClient(ctx context.Context, uri string, retry int32) (*mongo.Client, error) {
	//clientOptions := options.Client().ApplyURI(uri)
	//clientOptions.SetMaxPoolSize(100)
	//clientOptions.SetMinPoolSize(10)
	//clientOptions.SetMaxConnIdleTime(10 * time.Second)
	//conn, err := mongo.Connect(ctx, clientOptions)
	conn, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://datastore:27017"))

	if err := conn.Ping(context.Background(), nil); err != nil {
		if retry >= 3 {
			return nil, err
		}
		retry = retry + 1
		time.Sleep(time.Second * 2)
		return CreateClient(ctx, uri, retry)
	}

	/*
		if err := conn.Ping(ctx, nil); err != nil {
			if retry >= 3 {
				return nil, err
			}
			retry = retry + 1
			time.Sleep(time.Second * 2)
			return CreateClient(ctx, uri, retry)
		}
	*/
	log.Printf("db connetion : %v", *conn)
	return conn, err
}
