package views

import (
	"strings"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
)

// @Summary add a new item of the computers
// @ID create-computer
// @Tags Computers
// @Produce json
// @Param data body schemas.ComputerCreateSchema true "Schema by Create New Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [post]
func Computer_POST(c *gin.Context) {

	responseCreateComputer := models.Response{
		Message: "No Create Computer",
		Success: false,
		Data:    "{}",
	}

	// Decodificar el objeto JSON recibido
	var computerCreateSchema schemas.ComputerCreateSchema
	if err := c.ShouldBindWith(&computerCreateSchema, binding.JSON); err != nil {
		responseCreateComputer := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateComputer)
		return
	}

	computer := models.Computer{
		IdRoom:     computerCreateSchema.IdRoom,
		IdComputer: uuid.New().String(),
		Pos_x:      0.0,
		Pos_y:      0.0,
		Name:       computerCreateSchema.Name,
		State:      "Disponible",
		Message:    "",
		Type:       computerCreateSchema.Type,
	}

	result, message, new_compuer := services.CreateComputer(computer)

	if result {

		responseCreateComputer = models.Response{
			Message: message,
			Success: result,
			Data:    new_compuer,
		}
	} else {
		responseCreateComputer = models.Response{
			Message: message,
			Success: responseCreateComputer.Success,
			Data:    nil,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateComputer)
}

// @Summary get a item of the computers
// @ID get-computer
// @Tags Computers
// @Produce json
// @Param id-computer path string true "ID of Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [get]
func Computer_GET(c *gin.Context) {

	result, message, computer := services.FindComputer(c.Param("id-computer"))

	responseGetComputer := models.Response{
		Message: message,
		Success: result,
		Data:    computer,
	}

	if !result {
		responseGetComputer.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetComputer)
}

// @Summary delete a item of the computers
// @ID delete-computer
// @Tags Computers
// @Produce json
// @Param id-computer path string true "ID of Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [delete]
func Computer_DELETE(c *gin.Context) {

	result, message, computer := services.FindComputer(c.Param("id-computer"))

	if result {
		result, message, _ = services.DeleteComputer(c.Param("id-computer"))
	}

	responseDeleteComputer := models.Response{
		Message: message,
		Success: result,
		Data:    computer,
	}

	if !result {
		responseDeleteComputer.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseDeleteComputer)
}

// @Summary update a item of the computers
// @ID put-computer
// @Tags Computers
// @Produce json
// @Param id-computer path string true "ID of Computer"
// @Param data body schemas.ComputerUpdateSchema true "Schema by Update New Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [put]
func Computer_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido
	var computerUpdateSchema schemas.ComputerUpdateSchema
	if err := c.ShouldBindWith(&computerUpdateSchema, binding.JSON); err != nil {
		responseUpdateComputer := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseUpdateComputer)
		return
	}

	result, message, computer := services.FindComputer(c.Param("id-computer"))

	if result {
		result, message, computer = services.UpdateComputer(c.Param("id-computer"), computerUpdateSchema)
	}

	responseUpdateComputer := models.Response{
		Message: message,
		Success: result,
		Data:    computer,
	}

	if !result {
		responseUpdateComputer.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseUpdateComputer)
}
