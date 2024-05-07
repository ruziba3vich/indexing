package models

type Person struct {
	Id        int    `json:"id"`
	Firstname string `json:"first_name"`
	Lastname  string `json:"last_name"`
}
