package main

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type RedisData struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func GetData(w http.ResponseWriter, r *http.Request) {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	ctx := context.Background()

	val, err := c.Get(ctx, "key1").Result()

	switch {
	case err == redis.Nil:
		panic("key does not exist")
	case err != nil:
		panic(err)
	case val == "":
		panic("value is empty")
	}

	data := RedisData{
		Key: "key1",
		Val: val,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func RegisterData(w http.ResponseWriter, r *http.Request) {
	data := RedisData{
		Key: "time",
		Val: "000100100101010100",
	}

	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()

	if err := c.Set(ctx, data.Key, data.Val, 0).Err(); err != nil {
		panic(err)
	}
}
