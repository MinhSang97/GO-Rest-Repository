package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Tiếp tục chuỗi middleware
		c.Next()

		// Kiểm tra nếu có lỗi trong request
		if len(c.Errors) > 0 {
			// Lấy ra lỗi đầu tiên trong danh sách lỗi
			err := c.Errors[0].Err

			// In thông báo lỗi ra console (có thể thay đổi cách xử lý ở đây)
			fmt.Println("Error:", err)

			// Trả về một phản hồi lỗi cho client
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Internal Server Error",
			})
		}
	}
}
