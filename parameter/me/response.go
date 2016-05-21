package me

import (
	"encoding/json"

	usermodel "github.com/gotokatsuya/break/model/user"
)

type responseObject struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	PhotoURL string `json:"photo_url"`
	Token    string `json:"token"`
}

// Convert ...
func (r *responseObject) Convert(user *usermodel.Entity) error {
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, r); err != nil {
		return err
	}
	return nil
}
