package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maps90/go-poll/handlers"
	"github.com/maps90/go-poll/models"
	"log"
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
	}
	router.Run(":8080")
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
}

func postVote(c *gin.Context) {
	var vote Vote
	c.Bind(&vote)

	su, err := models.StoreUser(vote.Email)
	if err != nil {
		log.Panic(err)
		return
	}
	gc, err := models.GetCandidateById(vote.Cid)
	handlers.Error(err)

	var vgc int = gc.Id
	if vgc == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"message": "not a valid form"})
		return
	}

	if w := su.(int64); w == 0 {
		c.JSON(http.StatusOK, gin.H{"message": "you have already choosen a path, there is no going back."})
	} else {
		_, err := models.StoreVote(vote.Cid)
		handlers.Error(err)
		c.JSON(http.StatusOK, gin.H{"message": "you have join the " + gc.Name + ", " + gc.Description})
	}
}
