package views

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/schemas"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/api/v1/services"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/models"
	"github.com/UNIHacks/UNIAccounts-BackEnd/src/tools"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

// @Summary SignIn User
// @ID signin-user
// @Tags Accounts
// @Produce json
// @Param data body schemas.UserSignInSchema true "Schema by SignIn User"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/signin [put]
func SignIn_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido en la estructura User
	var user schemas.UserSignInSchema
	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error to Decode JSON Body"))
		return
	}

	// Crear un cliente HTTP personalizado con los headers deseados
	client := &http.Client{}
	//Todo Remplazar la url de desarrollo o producción
	apiURL := "http://unicappaccountsdev.pythonanywhere.com/api/v1/signin"
	jsonData, err := json.Marshal(user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error to serializer data on JSON"))
		return
	}

	req, err := http.NewRequest(http.MethodPut, apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error to Create HTTP Request"))
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
		c.Writer.Write([]byte("Error when making the request to the external API"))
		return
	}
	defer resp.Body.Close()

	// Copiar el contenido de la respuesta del otro servicio a la respuesta HTTP en tu endpoint
	c.Header("Content-Type", "application/json")
	io.Copy(c.Writer, resp.Body)
}

// @Summary SignUp User
// @ID signup-user
// @Tags Accounts
// @Produce json
// @Param data body schemas.UserSignUpSchema true "Schema by SignUp User"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/signup [post]
func SignUp_POST(c *gin.Context) {

	// Decodificar el objeto JSON recibido en la estructura User
	var user schemas.UserSignUpSchema
	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error to Decode JSON Body"))
		return
	}

	// Crear un cliente HTTP personalizado con los headers deseados
	client := &http.Client{}
	//Todo Remplazar la url de desarrollo o producción
	apiURL := fmt.Sprintf("http://unicappaccountsdev.pythonanywhere.com/api/%s/signup", config.UNIACCOUNTS_API_VERSION)
	jsonData, err := json.Marshal(user)
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error to serializer data on JSON"))
		return
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error to Create HTTP Request"))
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
		c.Writer.Write([]byte("Error when making the request to the external API"))
		return
	}
	defer resp.Body.Close()

	// Copiar el contenido de la respuesta del otro servicio a la respuesta HTTP en tu endpoint
	c.Header("Content-Type", "application/json")
	io.Copy(c.Writer, resp.Body)
}

// @Summary LinkAccount User whit ComputerLab
// @ID post-link-account
// @Tags Accounts
// @Produce json
// @Param data body schemas.LinkAccountRequisitionSchema true "Schema by LinkAccount User whit ComputerLab"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/link-account [post]
func LinkAccount_POST(c *gin.Context) {

	// Decodificar el objeto JSON recibido en la estructura User
	var user schemas.LinkAccountRequisitionSchema
	if err := c.ShouldBindWith(&user, binding.JSON); err != nil {
		responseLinkAccount := models.Response{
			Message: "Error to Get Content JSON",
			Success: false,
			Data:    strings.Split(err.Error(), "\n"),
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseLinkAccount)
		return
	}
	//crear un numero random de 6 digitos
	// Generar una semilla única utilizando la hora actual
	rand.Seed(time.Now().UnixNano())

	// Generar un número aleatorio de 6 dígitos
	code := rand.Intn(900000) + 100000

	//Obtener el nombre de sala
	//Verificar que exista el User Name

	responseLinkAccount := models.Response{
		Message: "Failed to request the linking.",
		Success: false,
		Data:    "{}",
	}

	linkAccountCode := models.LinkAccountCode{
		UserName:      user.UserName,
		IdComputerLab: user.IdComputerLab,

		Code: strconv.Itoa(code),
	}

	result, message, _ := services.FindComputerLab(linkAccountCode.IdComputerLab)

	if !result {
		responseCreateLinkAccountCode := models.Response{
			Message: message,
			Success: result,
			Data:    nil,
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseCreateLinkAccountCode)
		return
	}

	result, message, new_linkAccountCode := services.CreateLinkAccountCode(linkAccountCode)

	responseLinkAccount.Message = message
	responseLinkAccount.Success = result
	responseLinkAccount.Data = new_linkAccountCode

	if result {
		//generar Email
		linkAccountCode.Code = "******"
		responseLinkAccount.Data = new_linkAccountCode
		html_body := tools.GenerateLinkAccountHTML(strconv.Itoa(code))
		response := tools.SedMail(config.SMTP_USER, "Vinculacion de Sala", html_body)
		if !response {
			responseLinkAccount = models.Response{
				Message: "Failed to request the linking.",
				Success: response,
				Data:    "{}",
			}
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseLinkAccount)

}

// @Summary LinkAccount User whit ComputerLab
// @ID put-link-account
// @Tags Accounts
// @Produce json
// @Param data body schemas.LinkAccountConfirmationSchema true "Schema by LinkAccount User whit ComputerLab"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/link-account [put]
func LinkAccount_PUT(c *gin.Context) {

	// Decodificar el objeto JSON recibido en la estructura User
	var linkAccountConfirmationSchema schemas.LinkAccountConfirmationSchema
	if err := c.ShouldBindWith(&linkAccountConfirmationSchema, binding.JSON); err != nil {
		c.Writer.WriteHeader(http.StatusInternalServerError)
		c.Writer.Write([]byte("Error to Decode JSON Body"))
		return
	}

	responseLinkAccount := models.Response{
		Message: "Failed to confirm the linking.",
		Success: false,
		Data:    "{}",
	}

	result, _, linkAccountCode := services.FindLinkAccountCode(linkAccountConfirmationSchema.Code, linkAccountConfirmationSchema.UserName)

	if result {
		linkAccount := models.LinkAccount{
			UserName:      linkAccountCode.UserName,
			IdComputerLab: linkAccountCode.IdComputerLab,
		}

		result, message, new_linkAccount := services.CreateLinkAccount(linkAccount)

		if result {

			responseLinkAccount = models.Response{
				Message: message,
				Success: result,
				Data:    new_linkAccount,
			}
		} else {
			responseLinkAccount = models.Response{
				Message: message,
				Success: result,
				Data:    "{}",
			}
		}

	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseLinkAccount)
}

// @Summary LinkAccount User whit ComputerLab
// @ID post-link-account
// @Tags Accounts
// @Produce json
// @Param data body schemas.LinkAccountRequisitionSchema true "Schema by LinkAccount User whit ComputerLab"
// @Success 200 {object} models.Response
// @Failure 400 {object} models.Response
// @Router /api/v1/link-account [post]
func LinkAccountJanky_POST(c *gin.Context) {

	// Decodificar el objeto JSON recibido en la estructura User
	var linkAccountSchema schemas.LinkAccountRequisitionSchema
	if err := c.ShouldBindWith(&linkAccountSchema, binding.JSON); err != nil {
		responseLinkAccount := models.Response{
			Message: "Failed to confirm the linking.",
			Success: false,
			Data:    "{}",
		}
		c.Header("Content-Type", "application/json")
		c.JSON(200, responseLinkAccount)
		return
	}

	linkAccount := models.LinkAccount{
		UserName:      linkAccountSchema.UserName,
		IdComputerLab: linkAccountSchema.IdComputerLab,
	}
	responseLinkAccount := models.Response{
		Message: "Failed to confirm the linking.",
		Success: false,
		Data:    "{}",
	}

	result, message, new_linkAccount := services.CreateLinkAccount(linkAccount)

	if result {

		responseLinkAccount = models.Response{
			Message: message,
			Success: result,
			Data:    new_linkAccount,
		}
	} else {
		responseLinkAccount = models.Response{
			Message: message,
			Success: result,
			Data:    "{}",
		}
	}

	c.Header("Content-Type", "application/json")
	c.JSON(200, responseLinkAccount)

}
