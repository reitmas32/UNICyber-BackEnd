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
		IdRoom:  computerCreateSchema.IdRoom,
		Pos_x:   0.0,
		Pos_y:   0.0,
		Name:    computerCreateSchema.Name,
		State:   "Disponible",
		Message: "",
		Type:    computerCreateSchema.Type,
		IdState: computerCreateSchema.IdState,
	}

	result, message, _ := services.FindRoom(computer.IdRoom)

	if !result {
		responseCreateComputer := models.Response{
			Message: message,
			Success: result,
			Data:    nil,
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateComputer)
		return
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
// @Param id path string true "ID of Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [get]
func Computer_GET(c *gin.Context) {

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, computer := services.FindComputer(uint(id))

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
// @Param id path string true "ID of Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [delete]
func Computer_DELETE(c *gin.Context) {

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, computer := services.FindComputer(uint(id))

	if result {
		result, message, _ = services.DeleteComputer(uint(id))
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
// @Param id path string true "ID of Computer"
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

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, computer := services.FindComputer(uint(id))

	if result {
		result, message, computer = services.UpdateComputer(uint(id), computerUpdateSchema)
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

// @Summary get a item of the computers of Room
// @ID get-computers-of-room
// @Tags Computers
// @Produce json
// @Param id-room path string true "ID of Room"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computers/{id-room} [get]
func ComputersOfRoom_GET(c *gin.Context) {

	id, _ := (strconv.ParseUint(c.Param("id-room"), 10, 32))

	result, message, room := services.FindComputerOfRoom(uint(id))

	responseGetRoom := models.Response{
		Message: message,
		Success: result,
		Data:    room,
	}

	if !result {
		responseGetRoom.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseGetRoom)
}

// @Summary Set State of Compute
// @ID put-computer-set-state
// @Tags Computers
// @Produce json
// @Param id path string true "ID of Computer"
// @Param data body schemas.SetStateSchema true "Schema by Update New Computer"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer-set-state/{id} [put]
func ComputerSetState_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido
	var setStateSchema schemas.SetStateSchema
	if err := c.ShouldBindWith(&setStateSchema, binding.JSON); err != nil {
		setStateComputerResult := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, setStateComputerResult)
		return
	}

	id, _ := (strconv.ParseUint(c.Param("id"), 10, 32))

	result, message, computer := services.FindComputer(uint(id))

	if result {
		result, message, _ = services.FindComputer(uint(id))

		if result {

			result, message, computer = services.SetStateComputer(uint(id), setStateSchema.IdState)
		}
	}

	setStateComputerResult := models.Response{
		Message: message,
		Success: result,
		Data:    computer,
	}

	if !result {
		setStateComputerResult.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, setStateComputerResult)
}
