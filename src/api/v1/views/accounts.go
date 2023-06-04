package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
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

func LinkAccount_POST(c *gin.Context) {

	// Decodificar el objeto JSON recibido en la estructura User
	var user schemas.LinkAccountRequisitionSchema
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
		return
	}
	//crear un numero random de 6 digitos
	// Generar una semilla única utilizando la hora actual
	rand.Seed(time.Now().UnixNano())

	// Generar un número aleatorio de 6 dígitos
	code := rand.Intn(900000) + 100000

	//Obtener el nombre de sala
	//Verificar que exista el User Name

	//generar Email
	html_body := tools.GenerateLinkAccountHTML(strconv.Itoa(code))

	// Renderiza el contenido HTML
	tmpl, err := template.New("").Parse(html_body)
	if err != nil {
		c.String(500, "Error al renderizar el contenido HTML")
		return
	}

	// Crea un buffer para almacenar la salida renderizada
	var buf bytes.Buffer
	err = tmpl.Execute(&buf, nil)
	if err != nil {
		c.String(500, "Error al renderizar el contenido HTML")
		return
	}

	// Establece el tipo de contenido en la respuesta
	c.Header("Content-Type", "text/html; charset=utf-8")

	// Envía el contenido HTML renderizado como respuesta
	c.Status(http.StatusOK)
	c.Writer.Write(buf.Bytes())

}

func LinkAccount_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido en la estructura User
	var user schemas.LinkAccountConfirmationSchema
	err := json.NewDecoder(c.Request.Body).Decode(&user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error al obtener el contenido del archivo"))
		return
	}

	//añadir a la tabla LinkAccounts los campos [idComputerLab, username]

}
