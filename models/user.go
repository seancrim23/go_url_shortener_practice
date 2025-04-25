package models

import "time"

/*
What metadata do Users need?
How often should users be able to shorten urls?
*/
type User struct {
	Username     string    `json:"username" firestore:"shortened"`
	Email        string    `json:"email" firestore:"email"`
	Password     string    `json:"password" firestore:"password"`
	LastCreation time.Time `json:"lastCreated" firestore:"lastCreated"` //ideally to limit shortening
}
