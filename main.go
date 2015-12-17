package main

import (
	"github.com/garyburd/redigo/redis"
	"github.com/gin-gonic/gin"
	"gopkg.in/gcfg.v1"
)

type Config struct {
	Redis struct {
		Host string
		Port string
		Pass string
		Db   string
	}
}

var cfg Config

func main() {
	r := gin.Default()
	v1 := r.Group("api/v1")
	{
		v1.GET("candidates", GetCandidates)
		v1.GET("candidate/:id", GetCandidate)
		v1.GET("votes", GetVotes)
		v1.POST("votes", PostVote)
	}
	r.Run(":8080")
}

func RedisConnect() redis.Conn {
	err := gcfg.ReadFileInto(&cfg, "config.gcfg")
	HandleError(err)

	c, err := redis.Dial("tcp", cfg.Redis.Host+":"+cfg.Redis.Port)
	if cfg.Redis.Pass != "" {
		c.Do("AUTH", cfg.Redis.Pass)
	}
	if cfg.Redis.Db != "" {
		c.Do("SELECT", cfg.Redis.Db)
	}
	HandleError(err)
	return c
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}
