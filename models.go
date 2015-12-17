package main

var currentCandidateId int
var currentVoteId int

type User struct {
	Id    int    `redis:"Id"`
	Email string `redis:"Email"`
}

type Candidate struct {
	Id          int    `redis:"Id"`
	Name        string `redis:"Name"`
	Description string `redis:"Description`
	Timestamp   string `redis:"Timestamp"`
}

type Vote struct {
	Id        int       `redis:"Id"`
	User      User      `redis:"User"`
	Candidate Candidate `redis:"Candidate"`
	Timestamp string    `redis:"Timestamp"`
}

type Votes []Vote
type Candidates []Candidate
type Users []User

func init() {
	createCandidate(Candidate{
		Name:        "dark-side",
		Description: "Join Dark Side",
	})
	createCandidate(Candidate{
		Name:        "light-side",
		Description: "Join Rebel Alliance",
	})
}
