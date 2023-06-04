package views

import (
	"encoding/json"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary add a new item of the computer-lab
// @ID create-computer-lab
// @Tags Computer Lab
// @Produce json
// @Param data body schemas.ComputerLabCreateSchema true "Schema by Create New Computer Lab"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-lab [post]
// @Example responseSuccess
// {
//   "Message": "Create ComputerLab Successful",
//   "Success": true,
//   "Data": {
//	   "Message": "Create ComputerLab Successful",
//	   "Success": true,
//	   "Data": {
//	     "ID": 0,
//	     "CreatedAt": "0001-01-01T00:00:00Z",
//	     "UpdatedAt": "0001-01-01T00:00:00Z",
//	     "DeletedAt": null,
//	     "IdComputerLab": "3ed96d42-63a2-44d2-82a8-baf6f12fe80a",
//	     "Name": "Sala 2",
//	     "Description": ""
//	   }
//	 }
// }
// @Example responseError1
// {
//   "Message": "Error to Get Content JSON",
//   "Success": false,
//   "Data": "{}"
// }
// @Example responseError2
// {
//   "Message": "No Create Computer Lab",
//   "Success": false,
//   "Data": "{}"
// }
func ComputerLab_POST(c *gin.Context) {

	responseCreateComputer := models.Response{
		Message: "No Create Computer Lab",
		Success: false,
		Data:    "{}",
	}

	// Decodificar el objeto JSON recibido
	var computerLabCreateSchema schemas.ComputerLabCreateSchema
	err := json.NewDecoder(c.Request.Body).Decode(&computerLabCreateSchema)
	if err != nil {
		responseCreateComputer := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    "{}",
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

		responseCreateComputer = models.Response{
			Message: message,
			Success: result,
			Data:    computerLab,
		}
	} else {

		responseCreateComputer = models.Response{
			Message: message,
			Success: responseCreateComputer.Success,
			Data:    computerLab,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateComputer)
}

// @Summary get a item of the computer-lab
// @ID get-computer-lab
// @Tags Computer Lab
// @Produce json
// @Param id-computer-lab path string true "ID of ComputerLab"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-lab [get]
func ComputerLab_GET(c *gin.Context) {

	result, message, computerLab := services.FindComputerLab(c.Param("id-computer-lab"))

	responseGetComputerLab := models.Response{
		Message: message,
		Success: result,
		Data:    computerLab,
	}

	if !result {
		responseGetComputerLab.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetComputerLab)
}

// @Summary delete a item of the computer-lab
// @ID delete-computer-lab
// @Tags Computer Lab
// @Produce json
// @Param id-computer-lab path string true "ID del item"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-lab [delete]
func ComputerLab_DELETE(c *gin.Context) {

	result, message, computerLab := services.FindComputerLab(c.Param("id-computer-lab"))

	if result {
		result, message, _ = services.DeleteComputerLab(c.Param("id-computer-lab"))
	}

	responseDeleteComputerLab := models.Response{
		Message: message,
		Success: result,
		Data:    computerLab,
	}

	if !result {
		responseDeleteComputerLab.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseDeleteComputerLab)
}

// @Summary update a item of the computer-lab
// @ID put-computer-lab
// @Tags Computer Lab
// @Produce json
// @Param id-computer-lab path string true "ID del item"
// @Param data body schemas.ComputerLabUpdateSchema true "Schema by Update New Computer Lab"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-lab [put]
func ComputerLab_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido
	var computerLabUpdateSchema schemas.ComputerLabUpdateSchema
	err := json.NewDecoder(c.Request.Body).Decode(&computerLabUpdateSchema)
	if err != nil {
		responseUpdateComputerLab := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    "{}",
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseUpdateComputerLab)
		return
	}

	result, message, computerLab := services.FindComputerLab(c.Param("id-computer-lab"))

	if result {
		result, message, computerLab = services.UpdateComputerLab(c.Param("id-computer-lab"), computerLabUpdateSchema)
	}

	responseUpdateComputerLab := models.Response{
		Message: message,
		Success: result,
		Data:    computerLab,
	}

	if !result {
		responseUpdateComputerLab.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseUpdateComputerLab)
}
