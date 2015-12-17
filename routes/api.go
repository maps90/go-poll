package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/maps90/go-poll/handlers"
	"github.com/maps90/go-poll/models"
)

func PourGin() {
	router := gin.New()
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
	c.JSON(200, candidates)
}

func getCandidate(c *gin.Context) {
	id := c.Params.ByName("id")
	candidate, err := models.GetCandidateById(id)
	handlers.Error(err)
	c.JSON(200, candidate)
}

func getVote(c *gin.Context) {

}

func postVote(c *gin.Context) {

}
