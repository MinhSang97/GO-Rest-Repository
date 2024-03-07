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

//// Hàm để tính toán khoảng thời gian dựa trên giá trị của Period
//func addTimeForPeriod(endDate time.Time, period string) time.Time {
//	// Tách số và đơn vị từ chuỗi Period
//	var num int
//	var unit string
//	fmt.Sscanf(period, "%d%s", &num, &unit)

//	// Kiểm tra đơn vị thời gian và thêm tương ứng vào endDate
//	switch unit {
//	case "M":
//		return endDate.Add(time.Minute * time.Duration(num))
//	case "H":
//		return endDate.Add(time.Hour * time.Duration(num))
//	case "D":
//		return endDate.AddDate(0, 0, num)
//	default:
//		return endDate
//	}
//}

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
			c.JSON(http.StatusBadRequest, payload.Response{
				Error: fmt.Errorf("Error parsing start date: %w", err).Error(),
			})
			return
		}

		endDate, err := time.ParseInLocation("02-01-2006 15:04:05", requestGetHistories.EndDate, loc)
		if err != nil {
			c.JSON(http.StatusBadRequest, payload.Response{
				Error: fmt.Errorf("Error parsing end date: %w", err).Error(),
			})
			return
		}
		fmt.Println("end before: ", endDate)

		symbol := strings.ToLower(requestGetHistories.Symbol)

		//// Kiểm tra điều kiện về Period
		//switch requestGetHistories.Period {
		//case "30M", "1H", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "11H", "12H", "13H", "14H", "15H", "16H", "17H", "18H", "19H", "20H", "21H", "22H", "23H", "24H":
		//	endDate = addTimeForPeriod(endDate, requestGetHistories.Period)
		//	fmt.Println("end after: ", endDate)
		//case "2D", "3D", "4D", "5D", "6D", "7D":
		//	endDate = startDate.AddDate(0, 0, 7)
		//default:
		//	c.JSON(http.StatusBadRequest, payload.Response{
		//		Error: errors.New("Invalid period").Error(),
		//	})
		//	return
		//}

		period := strings.ToUpper(requestGetHistories.Period)

		uc := usecases.NewHistoriesUseCase()

		data, err := uc.GetHistories(c.Request.Context(), startDate, endDate, period, symbol)
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
