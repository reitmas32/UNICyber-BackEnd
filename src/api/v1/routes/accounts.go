package routes

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
)

type UserSignUp struct {
	Name        string `json:"name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	DateOfBirth string `json:"date_of_birth"`
	Password    string `json:"password"`
	Role        string `json:"role"`
}

type UserSignIn struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func SignUp_POST(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		// Cadena de texto JSON estática
		jsonData := `{
		"Data": {},
		"Message": "No Method Valid",
		"Success": false
	  }`

		// Establecer el encabezado Content-Type a "application/json"
		w.Header().Set("Content-Type", "application/json")

		// Enviar la cadena de texto JSON como respuesta HTTP
		fmt.Fprint(w, jsonData)

		return
	}

	// Decodificar el objeto JSON recibido en la estructura User
	var user UserSignUp
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar los datos JSON", http.StatusBadRequest)
		return
	}

	// Crear un cliente HTTP personalizado con los headers deseados
	client := &http.Client{}
	//Todo Remplazar la url de desarrollo o producción
	apiURL := fmt.Sprintf("http://127.0.0.1:%d/api/%s/signup", config.UNIACCOUNTS_API_PORT, config.UNIACCOUNTS_API_VERSION)
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error al serializar los datos JSON", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPost, apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error al crear la petición HTTP", http.StatusInternalServerError)
		return
	}

	// Añadir headers personalizados a la petición
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", "1234")
	req.Header.Set("Service", "uniaccounts")

	// Realizar la petición HTTP con el cliente personalizado
	resp, err := client.Do(req)

	// Leer el contenido de la respuesta
	if err != nil {
		http.Error(w, "Error al hacer la petición a la API externa", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copiar el contenido de la respuesta del otro servicio a la respuesta HTTP en tu endpoint
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}

func SignIn_PUT(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPut {
		// Cadena de texto JSON estática
		jsonData := `{
		"Data": {},
		"Message": "No Method Valid",
		"Success": false
	  }`

		// Establecer el encabezado Content-Type a "application/json"
		w.Header().Set("Content-Type", "application/json")

		// Enviar la cadena de texto JSON como respuesta HTTP
		fmt.Fprint(w, jsonData)

		return
	}

	// Decodificar el objeto JSON recibido en la estructura User
	var user UserSignIn
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Error al decodificar los datos JSON", http.StatusBadRequest)
		return
	}

	// Crear un cliente HTTP personalizado con los headers deseados
	client := &http.Client{}
	//Todo Remplazar la url de desarrollo o producción
	apiURL := "http://127.0.0.1:5000/api/v1/signin"
	jsonData, err := json.Marshal(user)
	if err != nil {
		http.Error(w, "Error al serializar los datos JSON", http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPut, apiURL, bytes.NewBuffer(jsonData))
	if err != nil {
		http.Error(w, "Error al crear la petición HTTP", http.StatusInternalServerError)
		return
	}

	// Añadir headers personalizados a la petición
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Api-Key", "1234")
	req.Header.Set("Service", "uniaccounts")

	// Realizar la petición HTTP con el cliente personalizado
	resp, err := client.Do(req)

	// Leer el contenido de la respuesta
	if err != nil {
		http.Error(w, "Error al hacer la petición a la API externa", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	// Copiar el contenido de la respuesta del otro servicio a la respuesta HTTP en tu endpoint
	w.Header().Set("Content-Type", "application/json")
	io.Copy(w, resp.Body)
}
