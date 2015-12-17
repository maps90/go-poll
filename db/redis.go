package db

import (
	"github.com/garyburd/redigo/redis"
	"github.com/maps90/go-poll/handlers"
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

func Connect() redis.Conn {
	var cfg Config
	err := gcfg.ReadFileInto(&cfg, "conf.gcfg")
	handlers.Error(err)

	c, err := redis.Dial("tcp", cfg.Redis.Host+":"+cfg.Redis.Port)
	if cfg.Redis.Pass != "" {
		c.Do("AUTH", cfg.Redis.Pass)
	}
	if cfg.Redis.Db != "" {
		c.Do("SELECT", cfg.Redis.Db)
	}
	handlers.Error(err)
	return c
}
