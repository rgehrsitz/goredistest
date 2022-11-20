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

	for {
		err := rdb.Set(ctx, "key"+strconv.Itoa(rand.Intn(1000)), rand.Intn(1000000), 0).Err()
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

	for i := 0; i < 1000; i++ {
		val, err := rdb.Get(ctx, "key"+strconv.Itoa(i)).Result()
		if err != nil {
			if err == redis.Nil {
				fmt.Println("key does not exist")
				//return
			}
			//panic(err)
		}
		fmt.Println(val)
		if i == 999 {
			i = 0
		}
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
