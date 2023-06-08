package views

import (
	"strings"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary add a new item of the rooms
// @ID create-room
// @Tags Rooms
// @Produce json
// @Param data body schemas.RoomCreateSchema true "Schema by Create New Room"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/room [post]
func Room_POST(c *gin.Context) {

	responseCreateRoom := models.Response{
		Message: "No Create Room",
		Success: false,
		Data:    "{}",
	}

	// Decodificar el objeto JSON recibido
	var roomSchema schemas.RoomCreateSchema
	if err := c.ShouldBindWith(&roomSchema, binding.JSON); err != nil {
		responseCreateRoom := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateRoom)
		return
	}

	room := models.Room{
		IdComputerLab: roomSchema.IdComputerLab,
		Name:          roomSchema.Name,
		//Create Index whit one more current max index of ComputerLab
	}

	result, message, _ := services.FindComputerLab(room.IdComputerLab)

	if !result {
		responseCreateRoom := models.Response{
			Message: message,
			Success: result,
			Data:    nil,
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateRoom)
		return
	}

	result, message, new_room := services.CreateRoom(room)

	if result {

		responseCreateRoom = models.Response{
			Message: message,
			Success: result,
			Data:    new_room,
		}
	} else {
		responseCreateRoom = models.Response{
			Message: message,
			Success: responseCreateRoom.Success,
			Data:    nil,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateRoom)
}

// @Summary get a item of the rooms
// @ID get-room
// @Tags Rooms
// @Produce json
// @Param id path string true "ID of Room"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/room [get]
func Room_GET(c *gin.Context) {

	result, message, room := services.FindRoom(c.Param("id"))

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

// @Summary delete a item of the rooms
// @ID delete-room
// @Tags Rooms
// @Produce json
// @Param id path string true "ID of Room"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/room [delete]
func Room_DELETE(c *gin.Context) {

	result, message, computer := services.FindRoom(c.Param("id"))

	if result {
		result, message, _ = services.DeleteRoom(c.Param("id"))
	}

	responseDeleteRoom := models.Response{
		Message: message,
		Success: result,
		Data:    computer,
	}

	if !result {
		responseDeleteRoom.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseDeleteRoom)
}

// @Summary update a item of the rooms
// @ID put-room
// @Tags Rooms
// @Produce json
// @Param id path string true "ID of Room"
// @Param data body schemas.RoomUpdateSchema true "Schema by Update New Room"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/room [put]
func Room_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido
	var roomUpdateSchema schemas.RoomUpdateSchema
	if err := c.ShouldBindWith(&roomUpdateSchema, binding.JSON); err != nil {
		responseUpdateRoom := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseUpdateRoom)
		return
	}

	result, message, room := services.FindRoom(c.Param("id"))

	if result {
		result, message, room = services.UpdateRoom(c.Param("id"), roomUpdateSchema)
	}

	responseUpdateRoom := models.Response{
		Message: message,
		Success: result,
		Data:    room,
	}

	if !result {
		responseUpdateRoom.Data = "{}"
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseUpdateRoom)
}
