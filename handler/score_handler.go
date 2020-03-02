package handler

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	"github.com/gin-gonic/gin"
	"github.com/goinggo/mapstructure"
	"google.golang.org/api/iterator"
)

/*
UpdateScore will update user score,
if there's no such user exits create one
*/
func UpdateScore(c *gin.Context) {
	ctx := context.Background()
	var json scoreJSON
	c.BindJSON(&json)

	log.Printf("%v", json)

	code := updateScoreByID(ctx, json)

	c.Status(code)
}

/*
GetScore will get score by user ID
*/
func GetScore(c *gin.Context) {
	ctx := context.Background()
	ID := c.DefaultQuery("ID", "no ID")

	score, err := getScoreByID(ctx, ID)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}
	user, err := getUserByID(ctx, ID)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	c.JSON(200, gin.H{
		"score": score,
		"name":  user.Name,
		"mail":  user.Mail,
	})
}

/*
GetRank will get top 10 player
*/
func GetRank(c *gin.Context) {
	ctx := context.Background()
	scores, err := getRank(ctx)
	if err != nil {
		log.Panicln(err)
		c.Status(500)
		return
	}

	c.JSON(200, scores)
}

type scoreJSON struct {
	Score  int    `json:"score" form:"score"`
	UserID string `json:"userId" form:"userID"`
}

func getRank(ctx context.Context) ([]map[string]interface{}, error) {
	client := getFireStoreClient(ctx)
	collection := client.Collection("scores")
	docs := collection.OrderBy("Score", firestore.Asc).Limit(10).Documents(ctx)
	var scores []map[string]interface{}

	for {
		var json scoreJSON
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return scores, err
		}
		if err := mapstructure.Decode(doc.Data(), &json); err != nil {
			return scores, err
		}
		user, err := getUserByID(ctx, json.UserID)
		if err != nil {
			return scores, err
		}
		scores = append(scores, map[string]interface{}{
			"score": json.Score,
			"name":  user.Name,
			"mail":  user.Mail,
		})
	}

	return scores, nil
}

func getScoreByID(ctx context.Context, ID string) (score int, err error) {
	client := getFireStoreClient(ctx)
	collection := client.Collection("scores")
	docs := collection.Where("ID", "==", ID).Limit(1).Documents(ctx)
	var json scoreJSON

	defer docs.Stop()
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return score, err
		}

		if err := mapstructure.Decode(doc.Data(), &json); err != nil {
			return score, err
		}
	}
	return json.Score, nil
}

func getScoreIDByUserID(ctx context.Context, ID string) (string, error) {
	client := getFireStoreClient(ctx)
	docs := client.Collection("scores").Where("ID", "==", ID).Limit(1).Documents(ctx)
	defer docs.Stop()
	scoreID := ""
	for {
		doc, err := docs.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			return scoreID, err
		}
		scoreID = doc.Ref.ID
	}
	return scoreID, nil
}

func updateScoreByID(ctx context.Context, json scoreJSON) int {
	client := getFireStoreClient(ctx)
	collection := client.Collection("scores")

	scoreID, err := getScoreIDByUserID(ctx, json.UserID)

	if err != nil {
		log.Panicln("cannot get scoreID", err)
		return 500
	}
	if scoreID == "" {
		if _, err := collection.NewDoc().Set(ctx, json); err != nil {
			log.Panicln("cannot get collection ref at score", err)
			return 500
		}
	} else {
		if _, err := collection.Doc(scoreID).Update(ctx, []firestore.Update{{Path: "Score", Value: json.Score}}); err != nil {
			log.Panicln("cannot get collection ref at user", err)
			return 500
		}
	}
	return 201
}
