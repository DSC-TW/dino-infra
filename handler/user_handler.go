package handler

import (
	"context"
	"dino-infra/model"
	"dino-infra/utility"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"google.golang.org/api/iterator"
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

func getUserByMail(ctx context.Context, mail string) (model.User, error) {
	client := getFireStoreClient(ctx)
	collection := client.Collection("users")
	docs := collection.Where("Mail", "==", mail).Limit(1).Documents(ctx)
	var user model.User
	iter := docs

	defer iter.Stop()
	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return user, err
		}

		if err := mapstructure.Decode(doc.Data(), &user); err != nil {
			return user, err
		}
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

	if _, err := pushUserToFirestore(ctx, user); err != nil {
		c.Status(500)
	}
	c.Status(201)
}

/*
GetUser returns user by user's email
*/
func GetUser(c *gin.Context) {
	ctx := context.Background()
	mail := c.DefaultQuery("mail", "no mail")

	user, err := getUserByMail(ctx, mail)
	if err != nil {
		c.Status(500)
		return
	}

	if mail == "" || user.Mail == "" {
		log.Println("mail ", mail)
		log.Println("user mail ", user.Mail)
		c.Status(404)
		return
	}

	c.JSON(200, gin.H{
		"user": user,
	})
}
