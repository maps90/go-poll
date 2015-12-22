package models

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/maps90/go-poll/db"
	"github.com/maps90/go-poll/handlers"
	"reflect"
	"strconv"
	"time"
)

var currentCandidateId int

type Candidate struct {
	Id          int    `redis:"Id"`
	Name        string `redis:"Name"`
	Description string `redis:"Description`
	Timestamp   string `redis:"Timestamp"`
}

func init() {
	StoreCandidate(Candidate{
		Name:        "Rebel Alliance",
		Description: "May the force be with you",
	})
	StoreCandidate(Candidate{
		Name:        "Galactic Empire",
		Description: "I find your lack of faith disturbing.",
	})
}

func StoreCandidate(ca Candidate) {
	rds := db.Connect()
	defer rds.Close()

	currentCandidateId += 1
	t := time.Now()
	ca.Id = currentCandidateId
	ca.Timestamp = t.Format(time.RFC3339)

	r := reflect.ValueOf(ca)
	key := "baseline:candidate:"

	for i := 0; i < r.NumField(); i++ {
		field := r.Type().Field(i).Name
		value := r.Field(i).Interface()

		_, err := rds.Do("HSET", key+strconv.Itoa(ca.Id), field, value)
		handlers.Error(err)

		_, err = rds.Do("ZINCRBY", "votes", 0, currentCandidateId)
		handlers.Error(err)
	}
}

func GetAllCandidates() ([]Candidate, error) {
	var candidates []Candidate
	rds := db.Connect()
	defer rds.Close()
	keys, err := rds.Do("KEYS", "baseline:candidate:*")
	handlers.Error(err)

	for _, k := range keys.([]interface{}) {
		var candidate Candidate
		v, err := redis.Values(rds.Do("HGETALL", k))
		handlers.Error(err)

		if err = redis.ScanStruct(v, &candidate); err != nil {
			fmt.Println(err)
			return nil, err
		}

		candidates = append(candidates, candidate)
	}

	return candidates, nil
}

func GetCandidateById(cid string) (Candidate, error) {
	var candidate Candidate

	rds := db.Connect()
	defer rds.Close()

	reply, err := redis.Values(rds.Do("HGETALL", "baseline:candidate:"+cid))
	handlers.Error(err)

	if err := redis.ScanStruct(reply, &candidate); err != nil {
		fmt.Println("err: ", err)
		return candidate, err
	}
	return candidate, nil
}
