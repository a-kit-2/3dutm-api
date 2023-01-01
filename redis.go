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

	val, err := c.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		fmt.Println("redis LRange Error:", err)
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
		panic(err)
	}

	for _, d := range data {
		err := c.RPush(ctx, d.Key, d.Val).Err()
		if err != nil {
			fmt.Println("redis RPush Error:", err)
		}
	}
}
