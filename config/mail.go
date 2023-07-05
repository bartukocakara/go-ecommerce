package config

import "os"

// MailConfig holds the configuration for mail settings
type MailConfig struct {
	Driver          string
	MailtrapAPIKey  string
	GmailUsername   string
	GmailPassword   string
	MailchimpAPIKey string
}

// GetMailConfig retrieves the mail configuration from environment variables
func GetMailConfig() *MailConfig {
	return &MailConfig{
		Driver:          os.Getenv("MAIL_DRIVER"),
		MailtrapAPIKey:  os.Getenv("MAILTRAP_API_KEY"),
		GmailUsername:   os.Getenv("GMAIL_USERNAME"),
		GmailPassword:   os.Getenv("GMAIL_PASSWORD"),
		MailchimpAPIKey: os.Getenv("MAILCHIMP_API_KEY"),
	}
}
