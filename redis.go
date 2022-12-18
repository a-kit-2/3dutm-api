package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/go-redis/redis/v8"
)

type Data struct {
	key   string
	value string
}

func GetData(w http.ResponseWriter, r *http.Request) {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	d := Data{
		key:   "key1",
		value: "value1",
	}
	val, err := c.Get(ctx, d.key).Result()
	switch {
	case err == redis.Nil:
		panic("key does not exist")
	case err != nil:
		panic(err)
	case val == "":
		panic("value is empty")
	}
	fmt.Println(d.key, val)
}

func RegisterData(w http.ResponseWriter, r *http.Request) {
	c := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
	ctx := context.Background()
	d := Data{
		key:   "key1",
		value: "value1",
	}
	if err := c.Set(ctx, d.key, d.value, 0).Err(); err != nil {
		panic(err)
	}
}
