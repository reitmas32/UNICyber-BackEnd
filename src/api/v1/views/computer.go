package views

import (
	"encoding/json"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func Computer_POST(c *gin.Context) {

	responseCreateRoom := map[string]interface{}{
		"Message": "No Create Computer",
		"Data":    "{}",
		"Success": false,
	}

	// Decodificar el objeto JSON recibido
	var computerCreateSchema schemas.ComputerCreateSchema
	err := json.NewDecoder(c.Request.Body).Decode(&computerCreateSchema)
	if err != nil {
		responseCreateRoom := map[string]interface{}{
			"Message": "Error to Get Content JSON",
			"Data":    "{}",
			"Success": false,
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateRoom)
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

		responseCreateRoom = map[string]interface{}{
			"Message": message,
			"Success": result,
			"Data":    computer,
		}
	} else {
		responseCreateRoom = map[string]interface{}{
			"Message": message,
			"Success": responseCreateRoom["Success"],
			"Data":    computer,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateRoom)
}

func Computer_GET(c *gin.Context) {

	result, message, computer := services.FindComputer(c.Param("id-computer"))

	responseCreateRoom := map[string]interface{}{
		"Message": message,
		"Success": result,
		"Data":    computer,
	}

	if !result {
		responseCreateRoom["Data"] = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateRoom)
}

func Computer_DELETE(c *gin.Context) {

	result, message, computer := services.FindComputer(c.Param("id-computer"))

	if result {
		result, message, _ = services.DeleteComputer(c.Param("id-computer"))
	}

	responseDeleteRoom := map[string]interface{}{
		"Message": message,
		"Success": result,
		"Data":    computer,
	}

	if !result {
		responseDeleteRoom["Data"] = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseDeleteRoom)
}
