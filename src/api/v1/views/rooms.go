package views

import (
	"encoding/json"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary add a new item of the rooms
// @ID create-room
// @Tags Rooms
// @Produce json
// @Param data body schemas.RoomCreateSchema true "Schema by Create New Room"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/computer [post]
func Room_POST(c *gin.Context) {

	responseCreateRoom := models.Response{
		Message: "No Create Room",
		Success: false,
		Data:    "{}",
	}

	// Decodificar el objeto JSON recibido
	var roomSchema schemas.RoomCreateSchema
	err := json.NewDecoder(c.Request.Body).Decode(&roomSchema)
	if err != nil {
		responseCreateRoom := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    "{}",
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateRoom)
		return
	}

	room := models.Room{
		IdComputerLab: roomSchema.IdComputerLab,
		IdRoom:        uuid.New().String(),
		Name:          roomSchema.Name,
	}

	result, message := services.CreateRoom(room)

	if result {

		responseCreateRoom = models.Response{
			Message: message,
			Success: result,
			Data:    room,
		}
	} else {
		responseCreateRoom = models.Response{
			Message: message,
			Success: responseCreateRoom.Success,
			Data:    room,
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateRoom)
}
