package tour

import (
	"encoding/json"

	tourmodel "github.com/gotokatsuya/break/model/tour"
	usermodel "github.com/gotokatsuya/break/model/user"
)

type responseObject struct {
	ID        int64              `json:"id"`
	Name      string             `json:"name"`
	PhotoURL  string             `json:"photo_url"`
	TakenHour float32            `json:"taken_hour"`
	User      responseUserObject `json:"user"`
}

// Convert ...
func (r *responseObject) Convert(tour *tourmodel.Entity, user *usermodel.Entity) error {
	b, err := json.Marshal(tour)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, r); err != nil {
		return err
	}
	if err := r.User.convertUser(user); err != nil {
		return err
	}
	return nil
}

type responseUserObject struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	PhotoURL string `json:"photo_url"`
}

func (r *responseUserObject) convertUser(user *usermodel.Entity) error {
	b, err := json.Marshal(user)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, r); err != nil {
		return err
	}
	return nil
}
