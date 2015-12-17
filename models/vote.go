package models

type Vote struct {
	Id        int       `redis:"Id"`
	User      User      `redis:"User"`
	Candidate Candidate `redis:"Candidate"`
	Timestamp string    `redis:"Timestamp"`
}
