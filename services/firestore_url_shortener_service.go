package services

import (
	"context"
	"errors"
	"fmt"
	"go_url_shortener/models"
	"math/rand"
	"os"
	"time"

	"cloud.google.com/go/firestore"
	firebase "firebase.google.com/go"
)

// need to implement service interface
type FirestoreUrlShortenerService struct {
	database *firestore.Client
}

func NewFirestoreUrlShortenerService(ctx context.Context, projectId string) (*FirestoreUrlShortenerService, func(), error) {
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

	return &FirestoreUrlShortenerService{database: database}, closeFunc, nil
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (f *FirestoreUrlShortenerService) CreateShortUrl(ctx context.Context, fullUrl string) (string, error) {
	//TODO figure out if theres any reason to change the short url length
	shortUrl := generateShortUrl(8)
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
		return "", errors.New("error generating short url")
	}

	return shortUrl, nil
}

// basically from the internet
// TODO update if theres any custom url or more fancy url generation i find
func generateShortUrl(length int) string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var result []byte

	for i := 0; i < length; i++ {
		index := seededRand.Intn(len(charset))
		result = append(result, charset[index])
	}

	return string(result)
}

func (f *FirestoreUrlShortenerService) GetLongUrl(ctx context.Context, shortUrl string) (string, error) {
	dsnap, err := f.database.Collection("URL").Doc(shortUrl).Get(ctx)
	if err != nil {
		fmt.Println("error getting long url")
		fmt.Println(err)
		return "", errors.New("error getting long url")
	}

	var u models.URL
	err = dsnap.DataTo(&u)
	if err != nil {
		fmt.Println(err)
		return "", errors.New("error getting long url")
	}

	return u.Original, nil
}
