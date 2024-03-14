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

		symbol := strings.ToLower(requestGetHistories.Symbol)

		// Kiểm tra điều kiện về Period
		period := requestGetHistories.Period

		var num int
		var unit string

		fmt.Sscanf(period, "%d%s", &num, &unit)

		if period == "MAX" {
			unit = "MAX"
		}
		fmt.Sscanf(period, "%d%s", &num, &unit)

		period = strings.ToUpper(requestGetHistories.Period)

		//Trường hợp nhỏ hơn 14 day
		day14 := 7 < num && num <= 14 && unit == "D"

		//Trường hợp nhỏ hơn 30 day
		day30 := 14 < num && num <= 30 && unit == "D"

		//Trường hợp nhỏ hơn 90 day
		day90 := 30 < num && num <= 90 && unit == "D"

		//Trường hợp nhỏ hơn 180 day
		day180 := 90 < num && num <= 180 && unit == "D"

		//Trường hợp nhỏ hơn 365 day
		day365 := 180 < num && num <= 365 && unit == "D"

		if num != 0 && num <= 7 && unit == "H" || period == "30M" {
			period = "1D"
		} else if num <= 7 && unit == "D" {
			period = "7D"
		} else if day14 {
			period = "14D"
		} else if day30 {
			period = "30D"
		} else if day90 {
			period = "90D"
		} else if day180 {
			period = "180D"
		} else if day365 {
			period = "365D"
		} else if num == 0 && period == "MAX" {
			period = "MAX"
		} else {
			c.JSON(http.StatusBadRequest, payload.Response{
				Error: errors.New("Invalid period").Error(),
			})
			return
		}
		fmt.Println(period)

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

	}
}
