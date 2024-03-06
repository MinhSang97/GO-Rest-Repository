//package handler
//
//import (
//	"app/payload"
//	"app/usecases"
//	"errors"
//	"fmt"
//	"github.com/gin-gonic/gin"
//	"net/http"
//)
//
//func Get_Histories() func(*gin.Context) {
//	return func(c *gin.Context) {
//
//		requestGetHistories := payload.RequestGetHistories{}
//		if err := c.ShouldBindJSON(&requestGetHistories); err != nil {
//			c.JSON(http.StatusBadRequest, payload.Response{
//				Error: fmt.Errorf("Login error: %w", err).Error(),
//			})
//		}
//		// In thông tin về request vào console
//		fmt.Println("Received request:")
//		fmt.Printf("StartTime: %s\n", requestGetHistories.StartTime)
//		fmt.Printf("EndTime: %s\n", requestGetHistories.EndTime)
//		fmt.Printf("Period: %s\n", requestGetHistories.Period)
//		fmt.Printf("Symbol: %s\n", requestGetHistories.Symbol)
//
//		// Kiểm tra xem có ít nhất một tham số được truyền vào không
//		if requestGetHistories.StartTime == "" || requestGetHistories.EndTime == "" ||
//			requestGetHistories.Period == "" || requestGetHistories.Symbol == "" {
//			c.JSON(http.StatusBadRequest, payload.Response{
//				Error: errors.New("At least one search parameter is required").Error(),
//			})
//			return
//		}
//
//		uc := usecases.NewRequestGetHistoriesUseCase()
//
//		student, err := uc.GetHistories(c.Request.Context(), requestGetHistories.StartTime, requestGetHistories.EndTime, requestGetHistories.Period, requestGetHistories.Symbol)
//		if err != nil {
//			c.JSON(http.StatusInternalServerError, gin.H{
//				"error": err.Error(),
//			})
//			return
//		}
//
//		c.JSON(http.StatusOK, gin.H{
//			"student": student,
//		})
//	}
//}

package handler

import (
	"app/payload"
	"app/usecases"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

func Get_Histories() func(*gin.Context) {
	return func(c *gin.Context) {
		requestGetHistories := payload.RequestGetHistories{}
		if err := c.ShouldBindJSON(&requestGetHistories); err != nil {
			c.JSON(http.StatusBadRequest, payload.Response{
				Error: fmt.Errorf("Request error: %w", err).Error(),
			})
			return
		}

		// Kiểm tra xem có ít nhất một tham số được truyền vào không
		if requestGetHistories.StartDate == "" || requestGetHistories.EndDate == "" ||
			requestGetHistories.Period == "" || requestGetHistories.Symbol == "" {
			c.JSON(http.StatusBadRequest, payload.Response{
				Error: errors.New("At least one search parameter is required").Error(),
			})
			return
		}

		loc, _ := time.LoadLocation("Asia/Ho_Chi_Minh") // Lấy múi giờ của Việt Nam
		startDate, err := time.ParseInLocation("02-01-2006 15:04:05", requestGetHistories.StartDate, loc)
		if err != nil {
			fmt.Println("Error parsing start date:", err)
		}
		fmt.Println(startDate)

		endDate, err := time.ParseInLocation("02-01-2006 15:04:05", requestGetHistories.EndDate, loc)
		if err != nil {
			fmt.Println("Error parsing end date:", err)
		}

		fmt.Println(endDate)

		symbol := strings.ToLower(requestGetHistories.Symbol)

		uc := usecases.NewHistoriesUseCase()

		data, err := uc.GetHistories(c.Request.Context(), startDate, endDate, requestGetHistories.Period, symbol)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}
