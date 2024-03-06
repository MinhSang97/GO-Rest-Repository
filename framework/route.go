package framework

import (
	"app/dbutil"
	"app/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Route() {
	db := dbutil.ConnectDB()
	fmt.Println("Connected: ", db)

	r := gin.Default()
	//r.Use(middleware.ErrorHandler())
	//r.Use(middleware.BasicAuthMiddleware())

	//v1 := r.Group("/v1")
	//{
	//
	//	v1.GET("/:id", handler.Get_Histories(db))
	//
	//}
	r.GET("/get_histories", handler.Get_Histories())
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")

}
