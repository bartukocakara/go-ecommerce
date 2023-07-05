package config

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type MailchimpService struct {
	APIKey    string
	Server    string
	ListID    string
	FromName  string
	FromEmail string
}

func NewMailchimpService(apiKey, server, listID, fromName, fromEmail string) *MailchimpService {
	return &MailchimpService{
		APIKey:    apiKey,
		Server:    server,
		ListID:    listID,
		FromName:  fromName,
		FromEmail: fromEmail,
	}
}

type MailchimpMergeFields struct {
	FNAME string `json:"FNAME"`
	LNAME string `json:"LNAME"`
}

type MailchimpRecipient struct {
	Email       string               `json:"email_address"`
	Status      string               `json:"status"`
	MergeFields MailchimpMergeFields `json:"merge_fields"`
}

type MailchimpMessage struct {
	Subject   string `json:"subject"`
	FromName  string `json:"from_name"`
	FromEmail string `json:"from_email"`
	ToName    string `json:"to_name"`
	ToEmail   string `json:"to_email"`
	HTML      string `json:"html"`
}

func (m *MailchimpService) SendForgetPasswordEmail(email, firstName, lastName, resetLink string) error {
	// Construct the recipient
	recipient := MailchimpRecipient{
		Email:  email,
		Status: "subscribed",
		MergeFields: MailchimpMergeFields{
			FNAME: firstName,
			LNAME: lastName,
		},
	}

	// Construct the message
	message := MailchimpMessage{
		Subject:   "Password Reset",
		FromName:  m.FromName,
		FromEmail: m.FromEmail,
		ToName:    firstName + " " + lastName,
		ToEmail:   email,
		HTML:      fmt.Sprintf("Click the link to reset your password: <a href=\"%s\">%s</a>", resetLink, resetLink),
	}

	// Construct the request body
	data := map[string]interface{}{
		"message":    message,
		"recipient":  recipient,
		"send_email": true,
	}

	body, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// Send the request
	url := fmt.Sprintf("https://%s.api.mailchimp.com/3.0/lists/%s/messages", m.Server, m.ListID)
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(body))
	if err != nil {
		return err
	}

	req.SetBasicAuth("apikey", m.APIKey)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send forget password email, status code: %d", resp.StatusCode)
	}

	return nil
}
