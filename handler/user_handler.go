package handler

import (
	"context"
	"dino-infra/model"
	"dino-infra/utility"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
)

const (
	projectURL = "https://dino-infra-f2f43.firebaseio.com"
)

func getFireStoreClient(ctx context.Context) *firestore.Client {
	app := utility.NewFirebase(ctx, projectURL)
	return utility.NewDatabase(ctx, app)
}

func pushUserToFirestore(ctx context.Context, user model.User) (*firestore.WriteResult, error) {
	client := getFireStoreClient(ctx)
	collection := client.Collection("users")
	return collection.NewDoc().Set(ctx, user)
}

/*
PushUser push user data to firebase
*/
func PushUser(c *gin.Context) {
	var user model.User
	ctx := context.Background()
	c.BindJSON(&user)

	if _, err := pushUserToFirestore(ctx, user); err != nil {
		log.Println("failed on push user to firestore")
		c.Status(500)
	}
	c.Status(201)
}
