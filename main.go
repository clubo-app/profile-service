package main

import (
	"log"

	"github.com/clubo-app/packages/stream"
	"github.com/clubo-app/profile-service/config"
	"github.com/clubo-app/profile-service/repository"
	"github.com/clubo-app/profile-service/rpc"
	"github.com/clubo-app/profile-service/service"
	"github.com/nats-io/nats.go"
)

func main() {
	c, err := config.LoadConfig()
	if err != nil {
		log.Fatalln(err)
	}

	opts := []nats.Option{nats.Name("User Service")}
	nc, err := stream.Connect(c.NATS_CLUSTER, opts)
	if err != nil {
		log.Fatalln(err)
	}
	defer nc.Close()
	stream := stream.New(nc)

	pool, err := repository.NewPGXPool(c.DB_USER, c.DB_PW, c.DB_NAME, c.DB_HOST, c.DB_PORT)
	if err != nil {
		log.Fatal(err)
	}
	defer pool.Close()

	q := repository.New(pool)

	up := service.NewUploadService(c.SPACES_ENDPOINT, c.SPACES_TOKEN)
	ps := service.NewProfileService(q)

	p := rpc.NewProfileServer(ps, up, stream)

	rpc.Start(p, c.PORT)
}
