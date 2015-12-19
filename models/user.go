package models

import "github.com/maps90/go-poll/db"

type User struct {
	Id    int    `redis:"Id"`
	Email string `redis:"Email"`
}

func StoreUser(e string) (interface{}, error) {
	rds := db.Connect()
	defer rds.Close()

	//create a unique store to redis
	reply, err := rds.Do("ZADD", "registrar", 1, e)
	if err != nil {
		return nil, err
	}

	return reply, nil
}
