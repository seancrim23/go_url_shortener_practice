package services

import (
	"fmt"
	"os"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

// need to implement service interface
type FirestoreUrlShortenerService struct {
	firebaseConfig *firebase.Config
	database       *firestore.Client
}

func NewFirestoreUrlShortenerService(projectId string) *FirestoreUrlShortenerService {
	if value := os.Getenv("replace with firestore emulator host"); value != "" {
		fmt.Println("using firestore emulator: ", value)
	}
	//setup any firebase config
	//database connection will be made on app start
	return &FirestoreUrlShortenerService{firebaseConfig: &firebase.Config{ProjectID: projectId}, database: nil}
}

func (f *FirestoreUrlShortenerService) CreateShortUrl(fullUrl string) (string, error) {
	//this function should take in a long url
	//a new id should be generated
	//build the url object (new url, original url, metadata....)
	//store it in the db
}

func (f *FirestoreUrlShortenerService) GetLongUrl(shortUrl string) error {
	//function takes the shortUrl
	//looks up entry in the db
	//return the longurl (in the future do some checks eg user permission, are we limiting user?, etc)
	//redirects (if no previous errors)
}
