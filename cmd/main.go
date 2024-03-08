package main

import (
	"app/framework"
	"app/helper"
	"app/redis"
	"fmt"
	"log"
	"time"
)

func main() {

	go func() {
		framework.Route()
		if err := redis.ConnectRedis(); err != nil {

			log.Fatal(err)
		}
	}()

	time.Sleep(10 * time.Minute)

	helper.ListCoin()

	go func() {
		for {
			fmt.Println("Calling ListCoin...")
			helper.ListCoin()
			time.Sleep(10 * time.Minute)
		}
	}()

	// run forever
	select {}
}
