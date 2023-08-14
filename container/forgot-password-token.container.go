package container

import (
	"ecommerce/internal/repository"
)

func (c *Container) BindForgotPasswordToken() {
	c.ForgotPasswordRepository = repository.NewForgotPasswordTokenRepository(c.DB)
}
