package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"go_url_shortener/services"
	"go_url_shortener/utils"
	"io"
	"net/http"
)

type Url struct {
	service services.UrlShortenerService
}

func (u *Url) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var code = 201
	var err error
	var longUrl string

	//body of request only needs to be the longurl
	reqBody, err := io.ReadAll(r.Body)
	if err != nil {
		code = 400
		fmt.Println(err)
		utils.RespondWithError(w, code, err.Error())
		return
	}

	err = json.Unmarshal(reqBody, &longUrl)
	if err != nil {
		code = 400
		fmt.Println(err)
		utils.RespondWithError(w, code, err.Error())
		return
	}

	shortUrl, err := u.service.CreateShortUrl(r.Context(), longUrl)
	//determine what type of error and change code and return according error message
	if err != nil {
		code = 500
		utils.RespondWithError(w, code, err.Error())
		return
	}

	utils.RespondWithJSON(w, code, shortUrl)
}

func (u *Url) GetLongUrl(w http.ResponseWriter, r *http.Request) {
	var code = 200
	var err error

	shortUrl := chi.UrlParam(r, "id")
	if shortUrl == "" {
		code = 400
		utils.RespondWithError(w, code, errors.New("no shortUrl passed to request").Error())
		return
	}

	//TODO get a check somewhere in here eventually to see if the url is in the cache

	longUrl, err := u.service.GetLongUrl(shortUrl)
	//determine what type of error and change code and return according error message
	if err != nil {
		code = 500
		utils.RespondWithError(w, code, err.Error())
		return
	}
	if longUrl == nil {
		code = 404
		utils.RespondWithError(w, code, errors.New("cannot find long url").Error())
		return
	}

	//should the backend redirect or should the front end do the redirect?
	utils.RespondWithJSON(w, code, longUrl)
}
