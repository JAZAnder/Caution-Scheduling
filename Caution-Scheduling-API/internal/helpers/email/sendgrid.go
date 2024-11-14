// using SendGrid's Go Library
// https://github.com/sendgrid/sendgrid-go
package email

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"

	"github.com/JAZAnder/Caution-Scheduling/internal/helpers/logger"
)

func Demo() {

	from := mail.NewEmail("Caution Scheduling", "scheduling@tutoring.cantusolutions.com")
	subject := "Sending with SendGrid is Fun"
	to := mail.NewEmail("Joshua Cantu", "joshua.cantu-2@selu.edu")
	plainTextContent := "and easy to do anywhere, even with Go"
	htmlContent := "<strong>Good Demo</strong>"
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

func sendEmail(toEmail string, toName string, subject string, htmlContent string, plainTextContent string) {
	from := mail.NewEmail("Caution Scheduling", "scheduling@tutoring.cantusolutions.com")
	to := mail.NewEmail(toName, toEmail)
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)

	if err != nil {
		logger.Log(4, "Email", "Send Email", "emailManager", err.Error())
	} else {
		logger.Log(2, "Email", "Send Email", "emailManager", strconv.Itoa(response.StatusCode)+"  "+response.Body)
	}

}
