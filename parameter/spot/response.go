package spot

import (
	"encoding/json"

	spotmodel "github.com/gotokatsuya/break/model/spot"
)

type responseObject struct {
	ID           int64   `json:"id"`
	VisitTime    int64   `json:"visit_time"`
	Name         string  `json:"name"`
	PhotoURL     string  `json:"photo_url"`
	CategoryName string  `json:"category_name"`
	State        string  `json:"state"`
	City         string  `json:"city"`
	Address      string  `json:"address"`
	Lat          float64 `json:"lat"`
	Lng          float64 `json:"lng"`
}

// Convert ...
func (r *responseObject) Convert(spot *spotmodel.Entity) error {
	b, err := json.Marshal(spot)
	if err != nil {
		return err
	}
	if err := json.Unmarshal(b, r); err != nil {
		return err
	}
	return nil
}
