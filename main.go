// shippy-service-vessel/main.go
package main

import (
	"context"
	"log"
	"os"

	pb "github.com/chanho3114/shippy-service-vessel/proto/vessel"
	"github.com/micro/go-micro/v2"
)

const (
	defaultHost = "datastore:27017"
)

func createDummyData(repo repository) {
	vessels := []*Vessel{
		{ID: "vessel001", Name: "Kane's Salty Secret", MaxWeight: 200000, Capacity: 500},
	}
	for _, v := range vessels {
		log.Printf("createDummyData : %v", *v)
		repo.Create(context.Background(), v)
	}
}

func main() {
	service := micro.NewService(
		micro.Name("shippy.service.vessel"),
	)

	service.Init()

	uri := os.Getenv("DB_HOST")
	if uri == "" {
		uri = defaultHost
	}
	log.Printf("uri : %s", uri)

	//ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	client, err := CreateClient(context.Background(), uri, 0)
	if err != nil {
		log.Panic(err)
	}
	defer client.Disconnect(context.Background())

	vesselCollection := client.Database("shippy").Collection("vessels")
	repository := &MongoRepository{
		vesselCollection,
	}

	createDummyData(repository)

	// Register our implementation with
	if err := pb.RegisterVesselServiceHandler(service.Server(), &handler{repository}); err != nil {
		log.Panic(err)
	}

	if err := service.Run(); err != nil {
		log.Panic(err)
	}
}
