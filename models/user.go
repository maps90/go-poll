package models

import "github.com/maps90/go-poll/db"

import (
	"encoding/json"
	"github.com/garyburd/redigo/redis"
	"github.com/maps90/go-poll/handlers"
	"time"
)

type User struct {
	Name      string    `redis:"Name" json:"Name"`
	Email     string    `redis:"Email" json:"Email"`
	Timestamp string    `redis:"Timestamp" json:"Timestamp"`
	Candidate Candidate `redis:"Candidate" json:"Candidate"`
}

type UserOption struct {
	Cid   string
	Name  string
	Email string
}

func StoreUser(opt UserOption) (interface{}, error) {

	rds := db.Connect()
	defer rds.Close()

	ca, err := GetCandidateById(opt.Cid)
	handlers.Error(err)

	if ca.Id == 0 {
		return nil, nil
	}

	t := time.Now()

	d := &User{
		opt.Name,
		opt.Email,
		t.Format(time.RFC3339),
		ca,
	}

	m, err := json.Marshal(d)

	reply, err := rds.Do("SADD", "registrar", opt.Email)
	if err != nil {
		return nil, err
	}

	if w := reply.(int64); w != 0 {
		_, err := rds.Do("SADD", "users", string(m))
		if err != nil {
			return nil, err
		}
	}

	return reply, nil
}

func GetAllUsers() ([]interface{}, error) {
	rds := db.Connect()
	defer rds.Close()
	keys, err := redis.Values(rds.Do("SMEMBERS", "users"))
	if err != nil {
		return nil, err
	}

	return keys, nil

}

func GetSith() (interface{}, error) {
	rds := db.Connect()
	defer rds.Close()
	keys, err := rds.Do("SSCAN", "users", 0, "MATCH", "*Galactic Empire*")
	if err != nil {
		return nil, err
	}

	return keys, nil
}

func GetJedi() (interface{}, error) {
	rds := db.Connect()
	defer rds.Close()
	keys, err := rds.Do("SSCAN", "users", 0, "MATCH", "*Rebel Alliance*")
	if err != nil {
		return nil, err
	}

	return keys, nil
}
