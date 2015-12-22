package models

import (
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
	"github.com/maps90/go-poll/db"
	"github.com/maps90/go-poll/handlers"
	"strconv"
)

const (
	INCR = 1
)

type VoteResult struct {
	Id        int
	Label     string
	VoteCount string
	Pct       string
}

func StoreVote(cid string) (interface{}, error) {
	rds := db.Connect()
	defer rds.Close()

	reply, err := rds.Do("ZINCRBY", "votes", INCR, cid)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func GetVotes() (interface{}, error) {
	rds := db.Connect()
	defer rds.Close()

	reply, err := rds.Do("GET", "results")
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func CalculateResult() {
	rds := db.Connect()
	defer rds.Close()

	candidate, err := GetAllCandidates()
	handlers.Error(err)

	var voteResults []VoteResult

	for i, _ := range candidate {
		var pct float64
		var vr VoteResult
		rank, err := redis.String(rds.Do("ZSCORE", "votes", candidate[i].Id))
		handlers.Error(err)

		total, err := getTotalVoters()
		handlers.Error(err)

		count, err := strconv.ParseFloat(rank, 64)
		handlers.Error(err)

		w := strconv.FormatInt(total.(int64), 10)
		fmt.Println(w)

		t, err := strconv.ParseFloat(w, 64)
		handlers.Error(err)

		pct = (count / t) * 100

		vr.Pct = strconv.FormatFloat(pct, 'f', 6, 64)
		vr.Id = candidate[i].Id
		vr.Label = candidate[i].Name
		vr.VoteCount = rank

		voteResults = append(voteResults, vr)

	}
	m, err := json.Marshal(voteResults)
	handlers.Error(err)

	status, err := rds.Do("SET", "results", string(m))

	if err != nil {
		panic(err)
	}

	fmt.Println(status)

}

func getTotalVoters() (interface{}, error) {
	rds := db.Connect()
	defer rds.Close()

	reply, err := rds.Do("SCARD", "registrar")
	if err != nil {
		return 0, err
	}

	return reply, nil

}
