package mail

import (
	"bytes"
	"ecommerce/config"
	"ecommerce/internal/lib"
	"text/template"
)

type ForgotPasswordMailer struct {
	MailLib *lib.MailLib
}

// NewForgotPasswordMailer creates a new instance of the ForgotPasswordMailer
func NewForgotPasswordMailer(mailConfig *config.MailConfig) *ForgotPasswordMailer {
	mailLib := lib.NewMailLib(mailConfig)
	return &ForgotPasswordMailer{
		MailLib: mailLib,
	}
}

// SendForgotPasswordEmail sends a forgot password email to the specified recipient
func (m *ForgotPasswordMailer) SendForgotPasswordEmail(recipientEmail string, resetLink string) error {
	// Load the email template
	t, err := template.ParseFiles("public/mails/forgot-password-mail.html")
	if err != nil {
		return err
	}

	// Prepare the data for the template
	data := struct {
		ResetLink string
	}{
		ResetLink: resetLink,
	}

	// Create a buffer to store the rendered template
	var body bytes.Buffer
	if err := t.Execute(&body, data); err != nil {
		return err
	}

	// Customize the email content, subject, and any other required details
	emailSubject := "Reset Your Password"
	emailBody := body.String()

	// Use the Mail library to send the email
	m.MailLib.SendMail(recipientEmail, emailSubject, emailBody)
	// Replace SendMail with the appropriate method based on the mail provider
	return nil
}
