package handler

import (
	"app/model"
	"app/payload"
	"app/usecases"
	"encoding/json"
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

		var num int
		var unit string
		fmt.Sscanf(requestGetHistories.Period, "%d%s", &num, &unit)

		if num < 7 {
			days := 1
			fmt.Println("days", days)
		} else if num < 14 {
			days := 7
			fmt.Println("days", days)
		}

		symbol := strings.ToLower(requestGetHistories.Symbol)
		period := strings.ToUpper(requestGetHistories.Period)

		// Kiểm tra điều kiện về Period
		switch requestGetHistories.Period {
		case "30M", "1H", "2H", "3H", "4H", "5H", "6H", "7H", "8H", "9H", "10H", "11H", "12H", "13H", "14H", "15H", "16H", "17H", "18H", "19H", "20H", "21H", "22H", "23H", "24H":

			uc := usecases.NewHistoriesUseCase()
			// Trong hàm GetHistories
			dataJSON, err := uc.GetHistories(c.Request.Context(), startDate, endDate, period, symbol)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			// Chuyển đổi dữ liệu JSON thành slice byte
			jsonData, err := json.Marshal(dataJSON)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			// Chuyển đổi dữ liệu JSON thành mảng cấu trúc model.OHLCData
			var ohldData []model.OHLCData
			err = json.Unmarshal(jsonData, &ohldData)
			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error(),
				})
				return
			}

			// Trả về dữ liệu đã chuyển đổi
			c.JSON(http.StatusOK, gin.H{
				"data": ohldData,
			})

		case "2D", "3D", "4D", "5D", "6D", "7D":

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
		default:
			c.JSON(http.StatusBadRequest, payload.Response{
				Error: errors.New("Invalid period").Error(),
			})
			return
		}

	}
}
