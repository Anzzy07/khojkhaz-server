package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendMail sends an email using SendGrid
func SendMail(userEmail string, subject string, html string) (bool, error) {
	apiKey := os.Getenv("SENDGRID_API_KEY")

	// Create a new SendGrid client
	client := sendgrid.NewSendClient(apiKey)

	// Define sender and recipient
	from := mail.NewEmail("Khojkhaz", "giftg4754@gmail.com")
	to := mail.NewEmail("User", userEmail)

	// Create the email content
	message := mail.NewSingleEmail(from, subject, to, "", html)

	// Send the email
	response, err := client.Send(message)
	if err != nil {
		log.Println("Failed to send email:", err)
		return false, err
	}

	fmt.Println("Email sent successfully:", response.StatusCode)
	fmt.Println(response.Body)
	fmt.Println(response.Headers)

	return true, nil
}
