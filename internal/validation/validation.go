package validation

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

// Messages holds the validation messages for different languages.
var Messages map[string]map[string]string

// LoadMessages loads the validation messages for all supported languages.
func LoadMessages() error {
	Messages = make(map[string]map[string]string)

	// Load English messages
	if err := loadMessages("en"); err != nil {
		return err
	}

	// Load Turkish messages
	if err := loadMessages("tr"); err != nil {
		return err
	}

	// Add more languages as needed

	return nil
}

// loadMessages loads the validation messages for a specific language.
func loadMessages(lang string) error {
	filePath := fmt.Sprintf("internal/validation/%s.json", lang)
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	var messages map[string]string
	if err := json.Unmarshal(data, &messages); err != nil {
		return err
	}

	Messages[lang] = messages
	return nil
}
