package handler

import (
	"app/payload"
	"app/usecases"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
)

func Get_Histories(db *gorm.DB) func(*gin.Context) {
	return func(c *gin.Context) {

		requestGetHistories := payload.RequestGetHistories{}
		if err := c.ShouldBindJSON(&requestGetHistories); err != nil {
			c.JSON(http.StatusBadRequest, payload.Response{
				Error: fmt.Errorf("Login error: %w", err).Error(),
			})
		}

		// Kiểm tra xem có ít nhất một tham số được truyền vào không
		if requestGetHistories.StartTime == "" || requestGetHistories.EndTime == "" {
			c.JSON(http.StatusBadRequest, payload.Response{
				Error: errors.New("At least one search parameter is required").Error(),
			})
			return
		}

		uc := usecases.NewStudentUseCase()

		student, err := uc.GetOneByID(c.Request.Context(), requestGetHistories.StartTime, requestGetHistories.EndTime, requestGetHistories.Period, requestGetHistories.Symbol)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"student": student,
		})
	}
}
