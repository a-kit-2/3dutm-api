package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/mux"
)

func redisClient() *redis.Client {
	addr := os.Getenv("ADDR")
	password := os.Getenv("PASSWORD")
	db, _ := strconv.Atoi(os.Getenv("DB"))

	c := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
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
		panic(err)
	}

	data := RedisData{
		Key: key,
		Val: val,
	}

	fmt.Println("Get data ", time.Now())

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func RegisterData(w http.ResponseWriter, r *http.Request) {
	c := redisClient()
	ctx := context.Background()

	reqBody, _ := ioutil.ReadAll(r.Body)

	var data RedisData

	if err := json.Unmarshal(reqBody, &data); err != nil {
		panic(err)
	}

	err := c.RPush(ctx, data.Key, data.Val).Err()
	if err != nil {
		panic(err)
	}

	fmt.Println("Register data ", time.Now())
}
