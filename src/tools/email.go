package tools

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"github.com/UNIHacks/UNIAccounts-BackEnd/src/config"
	gomail "gopkg.in/mail.v2"
)

func GenerateLinkAccountHTML(code string) string {
	filePath := "templates/link-account.html"

	// Leer el contenido del archivo HTML
	htmlContent, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}

	// Convertir el contenido HTML en una cadena
	htmlStr := string(htmlContent)

	// Reemplazar "{{ codigo }}" con el valor del código
	htmlContent = []byte(strings.Replace(htmlStr, "{{ codigo }}", code, 1))

	return string(htmlContent)
}

func SedMail(destination string, subject string, html_content string) bool {
	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", config.SMTP_USER)

	// Set E-Mail receivers
	m.SetHeader("To", destination)

	// Set E-Mail subject
	m.SetHeader("Subject", "Vinculacion Solicitada")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/html", html_content)

	// Settings for SMTP server
	d := gomail.NewDialer("smtp.office365.com", 587, config.SMTP_USER, config.SMTP_PASSWORD)

	// This is only needed when SSL/TLS certificate is not valid on server.
	// In production this should be set to false.
	d.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	// Now send E-Mail
	if err := d.DialAndSend(m); err != nil {
		fmt.Println(err)
		panic(err)
	}

	return true
}
