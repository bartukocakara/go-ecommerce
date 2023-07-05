package lib

import (
	"fmt"
	"log"
	"net/smtp"

	"ecommerce/config"
)

type MailLib struct {
	Config *config.MailConfig
	// Add any other required fields for the mail library
}

// NewMailLib creates a new instance of the Mail library
func NewMailLib(mailConfig *config.MailConfig) *MailLib {
	return &MailLib{
		Config: mailConfig,
		// Initialize any other required fields for the mail library
	}
}

// Add methods for sending emails using different providers
// For example, you can have methods like SendGmailEmail, SendMailtrapEmail, SendMailchimpEmail

func (m *MailLib) SendMail(recipientEmail, subject, body string) error {
	// Compose the email message
	message := []byte("To: " + recipientEmail + "\r\n" +
		"Subject: " + subject + "\r\n" +
		"Content-Type: text/html; charset=UTF-8\r\n" +
		"\r\n" +
		body)

	// Send the email using SMTP
	var auth smtp.Auth
	var smtpServer string

	switch m.Config.Driver {
	case "mailtrap":
		auth = smtp.PlainAuth("", "", "", "smtp.mailtrap.io")
		smtpServer = "smtp.mailtrap.io:587"
	case "gmail":
		auth = smtp.PlainAuth("", m.Config.GmailUsername, m.Config.GmailPassword, "smtp.gmail.com")
		smtpServer = "smtp.gmail.com:587"
	case "mailchimp":
		// Handle Mailchimp email sending logic
		// Replace the following with the appropriate implementation
		return fmt.Errorf("Mailchimp email sending is not implemented yet")
	default:
		return fmt.Errorf("Invalid mail driver specified")
	}

	err := smtp.SendMail(smtpServer, auth, m.Config.GmailUsername, []string{recipientEmail}, message)
	if err != nil {
		log.Println("Failed to send email:", err)
		return err
	}

	return nil
}
