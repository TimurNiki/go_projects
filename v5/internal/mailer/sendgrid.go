package mailer

import (
	"bytes"
	"fmt"
	"text/template"
	"time"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

// SendGrindMailer struct holds the configuration for sending emails
type SendGrindMailer struct {
	fromEmail string        // Email address from which the emails will be sent
	apiKey    string        // SendGrid API key for authentication
	client    *sendgrid.Client // SendGrid client for sending emails
}

// NewSendgrid initializes a new SendGrindMailer with the provided API key and from email
func NewSendgrid(apiKey, fromEmail string) *SendGrindMailer {
	client := sendgrid.NewSendClient(apiKey) // Create a new SendGrid client using the API key
	return &SendGrindMailer{
		fromEmail: fromEmail, // Set the from email address
		apiKey:    apiKey,    // Set the API key
		client:    client,    // Set the SendGrid client
	}
}

// Send method constructs and sends an email using a template
func (m *SendGrindMailer) Send(templateFile, username, email string, data any, isSandbox bool) (int, error) {
	// Create a new email object for the sender
	from := mail.NewEmail(FromName, m.fromEmail)
	// Create a new email object for the recipient
	to := mail.NewEmail(username, email)

	// Parse the email template from the file system
	tmpl, err := template.ParseFS(FS, "templates/"+templateFile)
	if err != nil {
		return -1, err // Return error if template parsing fails
	}

	// Create a buffer to hold the subject of the email
	subject := new(bytes.Buffer)
	// Execute the template for the subject using the provided data
	err = tmpl.ExecuteTemplate(subject, "subject", data)
	if err != nil {
		return -1, err // Return error if subject execution fails
	}

	// Create a buffer to hold the body of the email
	body := new(bytes.Buffer)
	// Execute the template for the body using the provided data
	err = tmpl.ExecuteTemplate(body, "body", data)
	if err != nil {
		return -1, err // Return error if body execution fails
	}

	// Create a new email message with the constructed subject and body
	message := mail.NewSingleEmail(from, subject.String(), to, "", body.String())

	// Set mail settings, specifically the sandbox mode based on the isSandbox flag
	message.SetMailSettings(&mail.MailSettings{
		SandboxMode: &mail.Setting{
			Enable: &isSandbox, // Enable or disable sandbox mode
		},
	})

	var retryErr error
	// Attempt to send the email up to maxRetries times
	for i := 0; i < maxRetires; i++ {
		response, retryErr := m.client.Send(message) // Send the email
		if retryErr != nil {
			// If there is an error, wait for an exponential backoff before retrying
			time.Sleep(time.Second * time.Duration(i+1))
			continue // Try sending the email again
		}

		return response.StatusCode, nil // Return the response status code if successful
	}

	// Return an error if the email failed to send after maxRetries attempts
	return -1, fmt.Errorf("failed to send email after %d attempts, error: %v", maxRetires, retryErr)
}
