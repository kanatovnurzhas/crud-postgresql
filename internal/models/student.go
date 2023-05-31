package models

type Student struct {
	Id         int    `json:"id"`
	FirstName  string `json:"firstName"`
	SecondName string `json:"secondName"`
	Email      string `json:"email"`
	Age        int    `json:"age"`
}
