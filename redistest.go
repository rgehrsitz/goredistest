package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

func sendRandomValues() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	fmt.Println("Sending data")
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 100; i++ {
		err := rdb.Set(ctx, "key"+strconv.Itoa(i), rand.Int(), 0).Err()
		if err != nil {
			panic(err)
		}
	}

}

func receiveValues() {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	ctx := context.Background()

	fmt.Println("Making query")

	for i := 0; i < 100; i++ {
		val, err := rdb.Get(ctx, "key"+strconv.Itoa(i)).Result()
		if err != nil {
			if err == redis.Nil {
				fmt.Println("key does not exist")
				return
			}
			panic(err)
		}
		fmt.Println(val)
	}

}

func main() {

	if len(os.Args) < 2 {
		os.Exit(3)
	}
	arg := os.Args[1]

	switch arg {
	case "send":
		sendRandomValues()
	case "receive":
		receiveValues()
	default:
		os.Exit(3)
	}

}
