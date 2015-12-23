package routes

import "github.com/gin-gonic/gin"

func PourGin() {
	router := gin.Default()
	router.Static("/img", "./public/img")
	router.Static("/js", "./public/js")
	router.Static("/css", "./public/css")
	router.StaticFile("/favicon.ico", "./public/favicon.ico")
	router.LoadHTMLGlob("./views/*.templ.html")
	router.GET("/", index)
	v1 := router.Group("api/v1")
	{
		v1.GET("candidates", getCandidates)
		v1.GET("candidate/:id", getCandidate)
		v1.GET("votes", getVote)
		v1.POST("votes", postVote)
		v1.GET("apprentices", getUser)
		v1.GET("jedi", getJedi)
		v1.GET("sith", getSith)
	}
	router.Run(":8080")
}

func index(c *gin.Context) {
	c.HTML(200, "index.templ.html", gin.H{"title": "Star Wars Poll"})
}
