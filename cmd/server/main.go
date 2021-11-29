package main

import (
	"context"
	"log"
	"net"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/proto/entpb"
	_ "github.com/mattn/go-sqlite3"
	"google.golang.org/grpc"
)

func main() {
	// Initialize an ent client.
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()

	// Run the migration tool (creating tables, etc).
	ctx := context.Background()
	if err = client.Schema.Create(ctx); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	server := grpc.NewServer()

	svc := entpb.NewUserService(client)
	entpb.RegisterUserServiceServer(server, svc)

	svc2 := entpb.NewPetService(client)
	entpb.RegisterPetServiceServer(server, svc2)

	// Open port 5000 for listening to traffic.
	lis, err := net.Listen("tcp", ":5000")
	if err != nil {
		log.Fatalf("failed listening: %s", err)
	}

	// Listen for traffic indefinitely.
	if err := server.Serve(lis); err != nil {
		log.Fatalf("server ended: %s", err)
	}
}
