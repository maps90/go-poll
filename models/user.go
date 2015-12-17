package models

type User struct {
	Id    int    `redis:"Id"`
	Email string `redis:"Email"`
}
