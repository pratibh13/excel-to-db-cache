package config

import (
	"assignment/models"
	"context"
	"database/sql"
	"log"

	"github.com/go-redis/redis/v8"
	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB
var RDB *redis.Client
var Ctx = context.Background()

func InitDB() {
	var err error
	dsn := "root:root@tcp(127.0.0.1:3306)/assignment"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to MySQL: %v", err)
	}
	models.Migrate(DB)
}

func InitRedis() {
	RDB = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	_, err := RDB.Ping(Ctx).Result()
	if err != nil {
		log.Fatalf("Could not connect to Redis: %v", err)
	}
}
