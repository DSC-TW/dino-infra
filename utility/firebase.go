package utility

import (
	"context"
	"log"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

/*
NewFirebase returns a firebase app
*/
func NewFirebase(ctx context.Context, databaseaURL string) *firebase.App {
	config := &firebase.Config{
		DatabaseURL: databaseaURL,
	}
	app, err := firebase.NewApp(ctx, config)
	if err != nil {
		log.Fatalf("error initializing app: %v\n", err)
	}

	return app
}

/*
NewDatabase returns a firebase firestore client
*/
func NewDatabase(ctx context.Context, app *firebase.App) *firestore.Client {
	client, err := app.Firestore(ctx)
	if err != nil {
		log.Fatalln("Error initializing database client:", err)
	}
	return client
}
