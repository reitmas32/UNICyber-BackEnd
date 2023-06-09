package views

import (
	"strconv"
	"strings"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
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

	responseCreateComputerLab := models.Response{
		Message: "No Create Computer Lab",
		Success: false,
		Data:    "{}",
	}

	// Decodificar el objeto JSON recibido
	var computerLabCreateSchema schemas.ComputerLabCreateSchema
	if err := c.ShouldBindWith(&computerLabCreateSchema, binding.JSON); err != nil {
		responseCreateComputerLab := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateComputerLab)
		return
	}

	computerLab := models.ComputerLab{
		Name:        computerLabCreateSchema.Name,
		Description: computerLabCreateSchema.Description,
	}

	result, message, newComputerLab := services.CreateComputerLab(computerLab)

	if result {

		responseCreateComputerLab = models.Response{
			Message: message,
			Success: result,
			Data:    newComputerLab,
		}
	} else {

		responseCreateComputerLab = models.Response{
			Message: message,
			Success: responseCreateComputerLab.Success,
			Data:    computerLab,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateComputerLab)
}

// @Summary get a item of the computer-lab
// @ID get-computer-lab
// @Tags Computer Lab
// @Produce json
// @Param id path string true "ID of ComputerLab"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-lab [get]
func ComputerLab_GET(c *gin.Context) {

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, computerLab := services.FindComputerLab(uint(id))

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
// @Param id path string true "ID del item"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-lab [delete]
func ComputerLab_DELETE(c *gin.Context) {

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, computerLab := services.FindComputerLab(uint(id))

	if result {
		result, message, _ = services.DeleteComputerLab(uint(id))
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
// @Param id path string true "ID del item"
// @Param data body schemas.ComputerLabUpdateSchema true "Schema by Update New Computer Lab"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-lab [put]
func ComputerLab_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido
	var computerLabUpdateSchema schemas.ComputerLabUpdateSchema
	if err := c.ShouldBindWith(&computerLabUpdateSchema, binding.JSON); err != nil {
		responseUpdateComputerLab := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseUpdateComputerLab)
		return
	}

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, computerLab := services.FindComputerLab(uint(id))

	if result {
		result, message, computerLab = services.UpdateComputerLab(uint(id), computerLabUpdateSchema)
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

// @Summary get all items of the computer-lab
// @ID get-computer-labs
// @Tags Computer Lab
// @Produce json
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-labs [get]
func ComputerLabs_GET(c *gin.Context) {

	result, message, computerLabs := services.GetComputerLabs()

	responseGetComputerLab := models.Response{
		Message: message,
		Success: result,
		Data:    computerLabs,
	}

	if !result {
		responseGetComputerLab.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetComputerLab)
}

// @Summary get N items of the computer-lab
// @ID get-computer-labs_limit
// @Tags Computer Lab
// @Produce json
// @Param length path string true "Length of result"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-labs-limit/{length} [get]
func ComputerLabs_Limit_GET(c *gin.Context) {

	length, _ := strconv.Atoi(c.Param("length"))

	result, message, computerLabs := services.GetComputerLabs_Limit(length)

	responseGetComputerLab := models.Response{
		Message: message,
		Success: result,
		Data:    computerLabs,
	}

	if !result {
		responseGetComputerLab.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetComputerLab)
}

// @Summary get N items of the computer-lab
// @ID get-computer-labs-by-user
// @Tags Computer Lab
// @Produce json
// @Param user_name path string true "User Name"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-labs-user/{user_name} [get]
func ComputerLabs_User_GET(c *gin.Context) {

	user_name := c.Param("user_name")

	result, message, computerLabs := services.GetComputerLabs_User(user_name)

	responseGetComputerLab := models.Response{
		Message: message,
		Success: result,
		Data:    computerLabs,
	}
	if result {
		data_list := responseGetComputerLab.Data.([]models.LinkAccount)
		var computerLabs []models.ComputerLab
		for i := 0; i < len(data_list); i++ {
			_, _, computerLab := services.FindComputerLab(data_list[i].IdComputerLab)
			computerLabs = append(computerLabs, computerLab)
		}

		responseGetComputerLab.Data = computerLabs
	} else {
		responseGetComputerLab.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetComputerLab)
}
