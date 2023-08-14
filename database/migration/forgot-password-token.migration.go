package migration

import (
	"ecommerce/internal/entity"
	"fmt"

	"gorm.io/gorm"
)

type ForgotPasswordTokenMigration struct {
	DB *gorm.DB
}

func NewForgotPasswordTokenMigration(db *gorm.DB) *ForgotPasswordTokenMigration {
	return &ForgotPasswordTokenMigration{DB: db}
}

func (m *ForgotPasswordTokenMigration) Migrate() {
	err := m.DB.AutoMigrate(&entity.ForgotPasswordToken{})
	if err != nil {
		fmt.Println("Failed to migrate ForgotPasswordToken table:", err)
		return
	}

	// Add a trigger to set the expires_at field to the current timestamp on insert
	// Create or replace the function to update expires_at
	m.DB.Exec(`
		CREATE OR REPLACE FUNCTION update_forgot_password_token_expires_at()
		RETURNS TRIGGER AS $$
		BEGIN
			NEW.expires_at = NOW() + INTERVAL '24 hours'; -- Set expiration to 24 hours from now
			RETURN NEW;
		END;
		$$ LANGUAGE plpgsql;
	`)

	// Create or replace the trigger
	m.DB.Exec(`
		CREATE TRIGGER set_forgot_password_token_expires_at
		BEFORE INSERT ON "forgot_password_tokens"
		FOR EACH ROW
		EXECUTE FUNCTION update_forgot_password_token_expires_at();
	`)
}
