package routes

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/maps90/go-poll/handlers"
	"github.com/maps90/go-poll/models"
	"net/http"
)

type Vote struct {
	Cid   string `form:"cid" json:"cid" binding:"required"`
	Email string `form:"email" json:"email" binding:"required"`
	Name  string `form:"name" json:"name" binding:"required"`
}

func PourGin() {
	router := gin.Default()
	v1 := router.Group("api/v1")
	{
		v1.GET("candidates", getCandidates)
		v1.GET("candidate/:id", getCandidate)
		v1.GET("votes", getVote)
		v1.POST("votes", postVote)
		v1.GET("apprentices", getUser)
	}
	router.Run(":8080")
}

func getUser(c *gin.Context) {
	var users []models.User
	reply, err := models.GetAllUsers()
	handlers.Error(err)
	for i := range reply {
		var usr models.User
		a := reply[i].([]uint8)
		if err := json.Unmarshal([]byte(a), &usr); err != nil {
			handlers.Error(err)
		}
		users = append(users, usr)
	}

	c.JSON(http.StatusOK, users)

}

func getCandidates(c *gin.Context) {
	candidates, err := models.GetAllCandidates()
	handlers.Error(err)
	c.JSON(http.StatusOK, candidates)
}

func getCandidate(c *gin.Context) {
	id := c.Params.ByName("id")
	candidate, err := models.GetCandidateById(id)
	handlers.Error(err)
	c.JSON(http.StatusOK, candidate)
}

func getVote(c *gin.Context) {
	votes, err := models.GetVotes()
	if err != nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "I've got a bad feeling about this"})
		return
	}

	switch w := votes.(type) {
	case []uint8:
		var vr []models.VoteResult
		err := json.Unmarshal([]byte(w), &vr)
		handlers.Error(err)
		c.JSON(http.StatusOK, vr)
	default:
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "I've got a bad feeling about this"})
		return
	}

}

func postVote(c *gin.Context) {
	var vote Vote

	if c.Bind(&vote) != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "not a valid form"})
		return
	}

	gc, err := models.GetCandidateById(vote.Cid)
	handlers.Error(err)
	var vgc int = gc.Id
	if vgc == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "not a valid form"})
		return
	}

	opt := models.UserOption{
		Cid:   vote.Cid,
		Name:  vote.Name,
		Email: vote.Email,
	}

	su, err := models.StoreUser(opt)

	if su == nil {
		c.JSON(http.StatusServiceUnavailable, gin.H{"message": "Opps, Something went wrong."})
		return
	}

	if w := su.(int64); w == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "you have already choosen a path, there is no going back."})
	} else {
		_, err := models.StoreVote(vote.Cid)
		handlers.Error(err)
		c.JSON(http.StatusOK, gin.H{"message": "you have join the " + gc.Name + ", " + gc.Description})
		models.CalculateResult()
	}
}
