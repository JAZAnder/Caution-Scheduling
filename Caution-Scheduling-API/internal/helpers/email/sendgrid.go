// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func main() {

	err := godotenv.Load("sendgrid.env")
	if err != nil {
		log.Println("Error Loading .env file")
	}

	godotenv.Load()
	from := mail.NewEmail("Caution Scheduling", "tutoring@cantusolutions.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Joshua Cantu", "joshua.cantu-2@selu.edu")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>and easy to do anywhere, even with Go</strong>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
