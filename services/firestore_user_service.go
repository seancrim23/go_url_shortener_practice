package services

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type FirestoreUserService struct {
	database *firestore.Client
}

func NewFirestoreUserService(ctx context.Context, projectId string) (*FirestoreUserService, func(), error) {
	if value := os.Getenv("replace with firestore emulator host"); value != "" {
		fmt.Println("using firestore emulator: ", value)
	}

	conf := &firebase.Config{ProjectID: projectId}
	app, err := firebase.NewApp(ctx, conf)
	if err != nil {
		fmt.Println("error making new firebase app: ", err)
		return nil, nil, err
	}

	database, err := app.Firestore(ctx)
	if err != nil {
		fmt.Println("error making firestore connection: ", err)
		return nil, nil, err
	}

	closeFunc := func() {
		database.Close()
	}

	return &FirestoreUserService{database: database}, closeFunc, nil
}
