package main

import (
	"app/framework"
	"app/redis"
)

func main() {

	redis.ConnectRedis()
	framework.Route()

}
