package me

import (
	"errors"

	"github.com/labstack/echo"
)

// LoginRequest ...
type LoginRequest struct {
	Email    string `form:"email" json:"email"`
	Name     string `form:"name" json:"name"`
	PhotoURL string `form:"photo_url" json:"photo_url"`
}

// NewLoginRequest ...
func NewLoginRequest(ctx echo.Context) (*LoginRequest, error) {
	req := &LoginRequest{}
	if err := ctx.Bind(req); err != nil {
		return nil, err
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return req, nil
}

// Validate ...
func (r *LoginRequest) Validate() error {
	if len(r.Email) <= 0 {
		return errors.New("Invalid Email")
	}
	if len(r.Name) <= 0 {
		return errors.New("Invalid Name")
	}
	return nil
}
