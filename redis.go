package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

type RedisData struct {
	Key string `json:"key"`
	Val string `json:"val"`
}

func redisClient() *redis.Client {
	const ADDR string = "localhost:6379"
	const PASSWORD string = ""
	const DB int = 0

	c := redis.NewClient(&redis.Options{
		Addr:     ADDR,
		Password: PASSWORD,
		DB:       DB,
	})

	return c
}

func GetData(w http.ResponseWriter, r *http.Request) {
	c := redisClient()
	ctx := context.Background()

	params := mux.Vars(r)
	key := params["key"]

	val, err := c.Get(ctx, key).Result()

	switch {
	case err == redis.Nil:
		panic("key does not exist")
	case err != nil:
		panic(err)
	case val == "":
		panic("value is empty")
	}

	data := RedisData{
		Key: key,
		Val: val,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func RegisterData(w http.ResponseWriter, r *http.Request) {
	c := redisClient()
	ctx := context.Background()

	reqBody, _ := ioutil.ReadAll(r.Body)

	var data []RedisData

	if err := json.Unmarshal(reqBody, &data); err != nil {
		fmt.Println(err)
	}

	for _, d := range data {
		if err := c.Set(ctx, d.Key, d.Val, 0).Err(); err != nil {
			panic(err)
		}
	}
}
