package mail

import (
	"bytes"
	"ecommerce/config"
	"ecommerce/internal/lib"
	"text/template"
)

type WelcomeMailer struct {
	MailLib *lib.MailLib
}

// NewWelcomeMailer creates a new instance of the WelcomeMailer
func NewWelcomeMailer(mailConfig *config.MailConfig) *WelcomeMailer {
	mailLib := lib.NewMailLib(mailConfig)
	return &WelcomeMailer{
		MailLib: mailLib,
	}
}

// SendWelcomeEmail sends a welcome email to the specified recipient
func (m *WelcomeMailer) SendWelcomeEmail(recipientEmail string, firstName string, lastName string) error {
	// Customize the email content, subject, and any other required details
	// Load the email template
	t, err := template.ParseFiles("public/mails/forgot-password-mail.html")
	if err != nil {
		return err
	}

	// Prepare the data for the template
	data := struct {
		WelcomeMessage string
	}{
		WelcomeMessage: "Welcome Message",
	}

	// Create a buffer to store the rendered template
	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return err
	}

	// Customize the email content, subject, and any other required details
	emailSubject := "Welcome"
	emailBody := body.String()

	// Use the Mail library to send the email
	m.MailLib.SendMail(recipientEmail, emailSubject, emailBody)
	// Replace SendMail with the appropriate method based on the mail provider
	return nil
}
