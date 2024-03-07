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
	// Khởi động một goroutine để gọi framework.Route() và redis.ConnectRedis()
	go func() {
		framework.Route()
		if err := redis.ConnectRedis(); err != nil {
			// Xử lý lỗi khi kết nối Redis
			log.Fatal(err)
		}
	}()

	// Chờ cho framework.Route() và redis.ConnectRedis() hoàn thành
	time.Sleep(10 * time.Minute) // Chờ 2 giây, có thể điều chỉnh tùy theo thời gian thực hiện của các hàm

	// Gọi hàm ListCoin() một lần khi ứng dụng khởi động
	helper.ListCoin()

	// Khai báo một goroutine để gọi ListCoin() mỗi 10 giây
	go func() {
		for {
			fmt.Println("Calling ListCoin...")
			helper.ListCoin()
			time.Sleep(10 * time.Second)
		}
	}()

	// Để cho chương trình chạy mãi mãi
	select {}
}
