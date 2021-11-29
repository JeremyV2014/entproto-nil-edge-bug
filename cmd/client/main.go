package main

import (
	"context"
	"entgo.io/bug/ent/proto/entpb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"log"
)

func main() {
	conn, err := grpc.Dial(":5000", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("failed connecting to server: %s", err)
	}
	defer conn.Close()

	userClient := entpb.NewUserServiceClient(conn)
	petClient := entpb.NewPetServiceClient(conn)

	ctx := context.Background()

	// // Attempt to create pet first
	//pet := &entpb.Pet{}
	//pet, err = petClient.Create(ctx, &entpb.CreatePetRequest{Pet: pet})
	//if err != nil {
	//	se, _ := status.FromError(err)
	//	log.Fatalf("failed creating pet: status=%s message=%s", se.Code(), se.Message())
	//}
	//log.Printf("%v", pet)
	//
	//user := &entpb.User{
	//	Age:  30,
	//	Name: "Bob",
	//	Pet: pet,
	//}
	//user, err = userClient.Create(ctx, &entpb.CreateUserRequest{User: user})
	//if err != nil {
	//	se, _ := status.FromError(err)
	//	log.Fatalf("failed creating user: status=%s message=%s", se.Code(), se.Message())
	//}
	//log.Printf("%v", user.Pet)
	//
	//user, err = userClient.Get(ctx, &entpb.GetUserRequest{
	//	Id:   user.Id,
	//	View: entpb.GetUserRequest_WITH_EDGE_IDS,
	//})
	//if err != nil {
	//	se, _ := status.FromError(err)
	//	log.Fatalf("failed getting user: status=%s message=%s", se.Code(), se.Message())
	//}
	//log.Printf("%v", user.GetPet())

	// Attempt to create user first
	user2 := &entpb.User{
		Age:  32,
		Name: "Alice",
	}
	user2, err = userClient.Create(ctx, &entpb.CreateUserRequest{User: user2})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating user: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("%v", user2.Pet)

	pet2 := &entpb.Pet{Owner: user2}
	pet2, err = petClient.Create(ctx, &entpb.CreatePetRequest{Pet: pet2})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed creating pet: status=%s message=%s", se.Code(), se.Message())
	}

	pet2, err = petClient.Get(ctx, &entpb.GetPetRequest{
		Id:   pet2.Id,
		View: entpb.GetPetRequest_WITH_EDGE_IDS,
	})
	if err != nil {
		se, _ := status.FromError(err)
		log.Fatalf("failed getting pet: status=%s message=%s", se.Code(), se.Message())
	}
	log.Printf("%v", pet2.GetOwner())
}
