package views

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Response struct {
	Data    Data   `json:"Data"`
	Message string `json:"Message"`
	Success bool   `json:"Success"`
}

type Data struct {
	Exp          int64  `json:"exp"`
	ServiceName  string `json:"service_name"`
	UserMotherID string `json:"user_mother_id"`
	UserName     string `json:"user_name"`
}

func Rooms_POST(c *gin.Context) {

	responseCreateRoom := map[string]interface{}{
		"Message": "No Create Room",
		"Data":    "{}",
		"Success": false, // o false
	}

	// Decodificar el objeto JSON recibido en la estructura User
	var roomSchema schemas.RoomSchema
	err := json.NewDecoder(c.Request.Body).Decode(&roomSchema)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
		return
	}

	// Crear un cliente HTTP personalizado con los headers deseados
	client := &http.Client{}
	//Todo Remplazar la url de desarrollo o producción
	apiURL := fmt.Sprintf("http://127.0.0.1:%d/api/%s/validate_token", config.UNIACCOUNTS_API_PORT, config.UNIACCOUNTS_API_VERSION)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al serializar los datos JSON"))
		return
	}

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al crear la petición HTTP"))
		return
	}

	// Añadir headers personalizados a la petición
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", config.UNIACCOUNTS_API_KEY)
	req.Header.Set("Service", config.UNIACCOUNTS_API_SERVICE_NAME)
	req.Header.Set("Authorization", fmt.Sprintf("JWT %s", roomSchema.JWT))

	// Realizar la petición HTTP con el cliente personalizado
	resp, err := client.Do(req)

	// Leer el contenido de la respuesta
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al hacer la petición a la API externa"))
		return
	}
	defer resp.Body.Close()

	body, err_json := ioutil.ReadAll(resp.Body)
	if err_json != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al leer el cuerpo de la respuesta"))
		return
	}

	var responseData map[string]interface{}
	err = json.Unmarshal(body, &responseData)
	if err != nil {
		fmt.Println("Error al decodificar el cuerpo JSON:", err)
		return
	}

	success := responseData["Success"].(bool)
	if success == true {
		var response Response
		if err := json.Unmarshal(body, &response); err != nil {
			fmt.Println("Error al analizar el JSON:", err)
			return
		}

		room := models.Room{
			IdUser: response.Data.UserMotherID,
			IdRoom: uuid.New().String(),
			Name:   roomSchema.Name,
		}

		result, message := services.CreateRoom(room)

		responseCreateRoom = map[string]interface{}{
			"Message": message,
			"Success": result,
			"Data":    room,
		}

	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseCreateRoom)
}
