package models

import (
	"github.com/maps90/go-poll/db"
)

const (
	INCR = 1
)

func StoreVote(cid string) (interface{}, error) {
	rds := db.Connect()
	defer rds.Close()

	reply, err := rds.Do("ZINCRBY", "votes", INCR, cid)
	if err != nil {
		return nil, err
	}

	return reply, nil
}

func GetAllVote() {
}

func GetVoteBy() {
}
