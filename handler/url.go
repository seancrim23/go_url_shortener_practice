package handler

import (
	"errors"
	"go_url_shortener/services"
	"go_url_shortener/utils"
	"html/template"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Url struct {
	Service services.UrlShortenerService
}

func (u *Url) GetURLForm(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("../templates/home.html"))
	tmpl.Execute(w, nil)
	return
}

func (u *Url) CreateShortUrl(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var code = 201
	var err error
	var longUrl string

	longUrl = r.FormValue("url")

	if longUrl == "" {
		code = 400
		utils.RespondWithError(w, code, errors.New("no longUrl passed to request").Error())
		return
	}

	shortUrl, err := u.Service.CreateShortUrl(r.Context(), longUrl)
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

	shortUrl := chi.URLParam(r, "id")
	if shortUrl == "" {
		code = 400
		utils.RespondWithError(w, code, errors.New("no shortUrl passed to request").Error())
		return
	}

	//TODO get a check somewhere in here eventually to see if the url is in the cache

	longUrl, err := u.Service.GetLongUrl(r.Context(), shortUrl)
	//determine what type of error and change code and return according error message
	if err != nil {
		code = 500
		utils.RespondWithError(w, code, err.Error())
		return
	}
	if longUrl == "" {
		code = 404
		utils.RespondWithError(w, code, errors.New("cannot find long url").Error())
		return
	}

	//should the backend redirect or should the front end do the redirect?
	utils.RespondWithJSON(w, code, longUrl)
}
