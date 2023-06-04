package views

import (
	"encoding/json"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary add a new item to the computers
// @ID create-computer
// @Produce json
// @Param data body schemas.ComputerCreateSchema true "todo data"
// @Success 200
// @Failure 400
// @Router /api/v1/computer [post]
func ComputerLab_POST(c *gin.Context) {

	responseCreateComputer := map[string]interface{}{
		"Message": "No Create Computer Lab",
		"Data":    "{}",
		"Success": false,
	}

	// Decodificar el objeto JSON recibido
	var computerLabCreateSchema schemas.ComputerLabCreateSchema
	err := json.NewDecoder(c.Request.Body).Decode(&computerLabCreateSchema)
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

	computerLab := models.ComputerLab{
		IdComputerLab: uuid.New().String(),
		Name:          computerLabCreateSchema.Name,
		Description:   computerLabCreateSchema.Description,
	}

	result, message := services.CreateComputerLab(computerLab)

	if result {

		responseCreateComputer = map[string]interface{}{
			"Message": message,
			"Success": result,
			"Data":    computerLab,
		}
	} else {
		responseCreateComputer = map[string]interface{}{
			"Message": message,
			"Success": responseCreateComputer["Success"],
			"Data":    computerLab,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateComputer)
}

func ComputerLab_GET(c *gin.Context) {

	result, message, computer := services.FindComputerLab(c.Param("id-computer-lab"))

	responseGetComputerLab := map[string]interface{}{
		"Message": message,
		"Success": result,
		"Data":    computer,
	}

	if !result {
		responseGetComputerLab["Data"] = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetComputerLab)
}

func ComputerLab_DELETE(c *gin.Context) {

	result, message, computerLab := services.FindComputerLab(c.Param("id-computer-lab"))

	if result {
		result, message, _ = services.DeleteComputerLab(c.Param("id-computer-lab"))
	}

	responseDeleteComputerLab := map[string]interface{}{
		"Message": message,
		"Success": result,
		"Data":    computerLab,
	}

	if !result {
		responseDeleteComputerLab["Data"] = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseDeleteComputerLab)
}

func ComputerLab_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido
	var computerLabUpdateSchema schemas.ComputerLabUpdateSchema
	err := json.NewDecoder(c.Request.Body).Decode(&computerLabUpdateSchema)
	if err != nil {
		responseUpdateComputerLab := map[string]interface{}{
			"Message": "Error to Get Content JSON",
			"Data":    "{}",
			"Success": false,
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseUpdateComputerLab)
		return
	}

	result, message, computerLab := services.FindComputerLab(c.Param("id-computer"))

	if result {
		result, message, computerLab = services.UpdateComputerLab(c.Param("id-computer-lab"), computerLabUpdateSchema)
	}

	responseRenameComputerLab := map[string]interface{}{
		"Message": message,
		"Success": result,
		"Data":    computerLab,
	}

	if !result {
		responseRenameComputerLab["Data"] = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseRenameComputerLab)
}
