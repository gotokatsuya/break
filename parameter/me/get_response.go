package me

import (
	usermodel "github.com/gotokatsuya/break/model/user"
)

// GetResponse ...
type GetResponse struct {
	Instance responseObject `json:"instance"`
}

// NewGetResponse ...
func NewGetResponse() *GetResponse {
	return &GetResponse{}
}

// Convert ...
func (r *GetResponse) Convert(user *usermodel.Entity) error {
	res := responseObject{}
	if err := res.Convert(user); err != nil {
		return err
	}
	r.Instance = res
	return nil
}
