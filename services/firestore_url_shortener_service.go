package services

import (
	"context"
	"errors"
	"fmt"
	"go_url_shortener/models"
	"os"
	"time"

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

func (f *FirestoreUrlShortenerService) CreateShortUrl(ctx context.Context, fullUrl string) (string, error) {
	//this function should take in a long url
	//a new id should be generated
	shortUrl, err := generateShortUrl(fullUrl)
	if err != nil {
		return "", err
	}

	currentTime := time.Now()
	builtUrl := &models.URL{
		Shortened: shortUrl,
		Original:  fullUrl,
		//TODO probably a config thing in the future to determine when a url expires
		//TODO maybe expired just gets pulled from db
		Expires:   currentTime.AddDate(0, 1, 0),
		Created:   currentTime,
		CreatedBy: "this will be whatever user is currently logged in but needs to be built",
	}

	//I think having the doc id be the shortened url is fine
	//one to one lookup
	//just have to make sure the db doesnt auto update when id exists
	_, err := f.database.Collection("URL").Doc(shortUrl).Set(ctx, builtUrl)
	if err != nil {
		fmt.Println("error generating short url")
		fmt.Println(err)
		return nil, errors.New("error generating short url")
	}

	return shortUrl
}

func generateShortUrl(longUrl string) (string, error) {
	//actually do some sort of business logic here
	return "abc123" + longUrl, nil
}

func (f *FirestoreUrlShortenerService) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	dsnap, err := f.database.Collection("URL").Doc(shortUrl).Get(ctx)
	if err != nil {
		fmt.Println("error getting long url")
		fmt.Println(err)
		return nil, errors.New("error getting long url")
	}

	var u models.URL
	err = dsnap.DataTo(&u)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("error getting long url")
	}

	return u.Original, nil
}
