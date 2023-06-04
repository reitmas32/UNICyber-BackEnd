package views

import (
	"encoding/json"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary add a new item of the computers
// @ID create-computer
// @Tags Computer
// @Produce json
// @Param data body schemas.ComputerCreateSchema true "Schema by Create New Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [post]
func Computer_POST(c *gin.Context) {

	responseCreateComputer := map[string]interface{}{
		"Message": "No Create Computer",
		"Data":    "{}",
		"Success": false,
	}

	// Decodificar el objeto JSON recibido
	var computerCreateSchema schemas.ComputerCreateSchema
	err := json.NewDecoder(c.Request.Body).Decode(&computerCreateSchema)
	if err != nil {
		responseCreateComputer := map[string]interface{}{
			"Message": "Error to Get Content JSON",
			"Data":    "{}",
			"Success": false,
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

	result, message := services.CreateComputer(computer)

	if result {

		responseCreateComputer = map[string]interface{}{
			"Message": message,
			"Success": result,
			"Data":    computer,
		}
	} else {
		responseCreateComputer = map[string]interface{}{
			"Message": message,
			"Success": responseCreateComputer["Success"],
			"Data":    computer,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateComputer)
}

// @Summary get a item of the computers
// @ID get-computer
// @Tags Computer
// @Produce json
// @Param id-computer path string true "ID of Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [get]
func Computer_GET(c *gin.Context) {

	result, message, computer := services.FindComputer(c.Param("id-computer"))

	responseGetComputer := map[string]interface{}{
		"Message": message,
		"Success": result,
		"Data":    computer,
	}

	if !result {
		responseGetComputer["Data"] = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetComputer)
}

// @Summary delete a item of the computers
// @ID delete-computer
// @Tags Computer
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

	responseDeleteComputer := map[string]interface{}{
		"Message": message,
		"Success": result,
		"Data":    computer,
	}

	if !result {
		responseDeleteComputer["Data"] = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseDeleteComputer)
}

// @Summary update a item of the computers
// @ID put-computer
// @Tags Computer
// @Produce json
// @Param id-computer path string true "ID of Computer"
// @Param data body schemas.ComputerUpdateSchema true "Schema by Update New Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [put]
func Computer_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido
	var computerUpdateSchema schemas.ComputerUpdateSchema
	err := json.NewDecoder(c.Request.Body).Decode(&computerUpdateSchema)
	if err != nil {
		responseUpdateComputer := map[string]interface{}{
			"Message": "Error to Get Content JSON",
			"Data":    "{}",
			"Success": false,
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseUpdateComputer)
		return
	}

	result, message, computer := services.FindComputer(c.Param("id-computer"))

	if result {
		result, message, computer = services.UpdateComputer(c.Param("id-computer"), computerUpdateSchema)
	}

	responseRenameComputer := map[string]interface{}{
		"Message": message,
		"Success": result,
		"Data":    computer,
	}

	if !result {
		responseRenameComputer["Data"] = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseRenameComputer)
}
