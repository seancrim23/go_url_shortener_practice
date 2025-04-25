package models

import "time"

/*
Any other URL metadata we need?
*/
type URL struct {
	Shortened string    `json:"shortened" firestore:"shortened"`
	Original  string    `json:"original" firestore:"original"`
	Expires   time.Time `json:"expires" firestore:"expires"`
	Created   time.Time `json:"created" firestore:"created"`
	CreatedBy string    `json:"createdBy" firestore:"createdBy"`
}
