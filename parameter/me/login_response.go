package me

import (
	usermodel "github.com/gotokatsuya/break/model/user"
)

// LoginResponse ...
type LoginResponse struct {
	Instance responseObject `json:"instance"`
}

// NewLoginResponse ...
func NewLoginResponse() *LoginResponse {
	return &LoginResponse{}
}

// Convert ...
func (r *LoginResponse) Convert(user *usermodel.Entity) error {
	res := responseObject{}
	if err := res.Convert(user); err != nil {
		return err
	}
	r.Instance = res
	return nil
}
