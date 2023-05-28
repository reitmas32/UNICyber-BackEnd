package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/gin-gonic/gin"
)

func SignIn_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido en la estructura User
	var user schemas.UserSignInSchema
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
		return
	}

	// Crear un cliente HTTP personalizado con los headers deseados
	client := &http.Client{}
	//Todo Remplazar la url de desarrollo o producción
	apiURL := "http://127.0.0.1:5000/api/v1/signin"
	jsonData, err := json.Marshal(user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al serializar los datos JSON"))
		return
	}

	req, err := http.NewRequest(http.MethodPut, apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al crear la petición HTTP"))
		return
	}

	// Añadir headers personalizados a la petición
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", config.UNIACCOUNTS_API_KEY)
	req.Header.Set("Service", config.UNIACCOUNTS_API_SERVICE_NAME)

	// Realizar la petición HTTP con el cliente personalizado
	resp, err := client.Do(req)

	// Leer el contenido de la respuesta
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al hacer la petición a la API externa"))
		return
	}
	defer resp.Body.Close()

	// Copiar el contenido de la respuesta del otro servicio a la respuesta HTTP en tu endpoint
	c.Header("Content-Type", "application/json")
	io.Copy(c.Writer, resp.Body)
}

func SignUp_POST(c *gin.Context) {

	// Decodificar el objeto JSON recibido en la estructura User
	var user schemas.UserSignUpSchema
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
		return
	}

	// Crear un cliente HTTP personalizado con los headers deseados
	client := &http.Client{}
	//Todo Remplazar la url de desarrollo o producción
	apiURL := fmt.Sprintf("http://127.0.0.1:%d/api/%s/signup", config.UNIACCOUNTS_API_PORT, config.UNIACCOUNTS_API_VERSION)
	jsonData, err := json.Marshal(user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al serializar los datos JSON"))
		return
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al crear la petición HTTP"))
		return
	}

	// Añadir headers personalizados a la petición
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", config.UNIACCOUNTS_API_KEY)
	req.Header.Set("Service", config.UNIACCOUNTS_API_SERVICE_NAME)

	// Realizar la petición HTTP con el cliente personalizado
	resp, err := client.Do(req)

	// Leer el contenido de la respuesta
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al hacer la petición a la API externa"))
		return
	}
	defer resp.Body.Close()

	// Copiar el contenido de la respuesta del otro servicio a la respuesta HTTP en tu endpoint
	c.Header("Content-Type", "application/json")
	io.Copy(c.Writer, resp.Body)
}