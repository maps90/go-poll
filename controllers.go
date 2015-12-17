package main

import (
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"reflect"
	"strconv"
	"time"
)

type candidates struct {
	candidate []Candidate
}

func GetCandidates(c *gin.Context) {
	var candidates []Candidate
	rds := RedisConnect()
	defer rds.Close()

	keys, err := rds.Do("KEYS", "baselines:candidate:*")
	HandleError(err)

	for _, k := range keys.([]interface{}) {
		var candidate Candidate
		v, err := redis.Values(rds.Do("HGETALL", k))
		HandleError(err)

		if err = redis.ScanStruct(v, &candidate); err != nil {
			fmt.Println(err)
			return
		}

		candidates = append(candidates, candidate)
	}
	c.JSON(200, candidates)

}

func createCandidate(ca Candidate) {
	currentCandidateId += 1
	t := time.Now()
	ca.Id = currentCandidateId
	ca.Timestamp = t.String()

	rds := RedisConnect()
	defer rds.Close()

	r := reflect.ValueOf(ca)
	for i := 0; i < r.NumField(); i++ {
		reply, err := rds.Do("HSET", "baselines:candidate:"+strconv.Itoa(ca.Id), r.Type().Field(i).Name, r.Field(i).Interface())
		HandleError(err)
		fmt.Println("store: ", reply)
	}
}

func GetCandidate(c *gin.Context) {
	id := c.Params.ByName("id")
	var candidate Candidate

	rds := RedisConnect()
	defer rds.Close()

	reply, err := redis.Values(rds.Do("HGETALL", "baselines:candidate:"+id))
	HandleError(err)

	if err := redis.ScanStruct(reply, &candidate); err != nil {
		fmt.Println("err: ", err)
		return
	}
	c.JSON(200, candidate)
}

func GetVotes(c *gin.Context) {
}

func PostVote(c *gin.Context) {
}
