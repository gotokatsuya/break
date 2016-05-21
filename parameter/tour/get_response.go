package tour

import (
	tourmodel "github.com/gotokatsuya/break/model/tour"
)

// GetResponse ...
type GetResponse struct {
	Instances []responseObject `json:"instances"`
}

// NewGetResponse ...
func NewGetResponse() *GetResponse {
	return &GetResponse{}
}

// Convert ...
func (r *GetResponse) Convert(tours []tourmodel.Entity) error {
	reses := make([]responseObject, len(tours))
	for i, tour := range tours {
		res := responseObject{}
		if err := res.Convert(&tour, &tour.User); err != nil {
			return err
		}
		reses[i] = res
	}
	r.Instances = reses
	return nil
}
