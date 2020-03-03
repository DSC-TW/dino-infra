package handler

import (
	"context"
	"dino-infra/model"
	"dino-infra/utility"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
)

const (
	projectURL = "https://dino-infra-f2f43.firebaseio.com"
)

func getFireStoreClient(ctx context.Context) *firestore.Client {
	app := utility.NewFirebase(ctx, projectURL)
	return utility.NewDatabase(ctx, app)
}

func pushUserToFirestore(ctx context.Context, user model.User) (string, error) {
	client := getFireStoreClient(ctx)
	collection := client.Collection("users")
	doc := collection.NewDoc()
	if _, err := collection.NewDoc().Set(ctx, user); err != nil {
		return "", err
	}
	return doc.ID, nil
}

func getUserByID(ctx context.Context, ID string) (returnUser model.User, err error) {
	var user model.User
	client := getFireStoreClient(ctx)
	collection := client.Collection("users")
	doc, err := collection.Doc(ID).Get(ctx)

	if err != nil {
		return user, err
	}

	if err := mapstructure.Decode(doc.Data(), &user); err != nil {
		return user, err
	}
	return user, nil
}

/*
PushUser push user data to firebase
*/
func PushUser(c *gin.Context) {
	var user model.User
	ctx := context.Background()
	c.BindJSON(&user)

	ID, err := pushUserToFirestore(ctx, user)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}
	c.JSON(201, gin.H{
		"ID": ID,
	})
}

/*
GetUser returns user by user's email
*/
func GetUser(c *gin.Context) {
	ctx := context.Background()
	ID := c.Param("ID")

	user, err := getUserByID(ctx, ID)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	if ID == "" {
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
