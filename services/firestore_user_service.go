package services

import (
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

type FirestoreUserService struct {
	firebaseConfig *firebase.Config
	database       *firestore.Client
}

func NewFirestoreUserService(projectId string) *FirestoreUserService {
	if value := os.Getenv("replace with firestore emulator host"); value != "" {
		fmt.Println("using firestore emulator: ", value)
	}
	//setup any firebase config
	//database connection will be made on app start
	return &FirestoreUserService{firebaseConfig: &firebase.Config{ProjectID: projectId}, database: nil}
}
