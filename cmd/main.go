package main

import (
	"app/framework"
	"app/redis"
)

func main() {

	framework.Route()
	redis.ConnectRedis()

}
