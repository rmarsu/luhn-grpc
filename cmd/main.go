package main

import (
	"context"
	"flag"
	api "luhn/api/proto"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	flag.Parse()

	if flag.NArg() < 1 {
		log.Fatal("not enough arguments")
	}

	text := flag.Arg(0)

	conn, err := grpc.NewClient(":8080", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}

	c := api.NewValidClient(conn)
	res, err := c.Valid(context.Background(), &api.ValidRequest{Number: text})
	if err != nil {
		log.Fatal(err)
	}
	log.Println(res.GetValid())
}
