package views

import (
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
)

// @Summary get all item of the UniversityPrograms
// @ID get-university-programs
// @Tags UniversityPrograms
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/university-programs [get]
func UniversityPrograms_GET(c *gin.Context) {

	result, message, states := services.GetUniversityPrograms()

	responseGetRoom := models.Response{
		Message: message,
		Success: result,
		Data:    states,
	}

	if !result {
		responseGetRoom.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetRoom)
}
