package spot

import (
	spotmodel "github.com/gotokatsuya/break/model/spot"
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
func (r *GetResponse) Convert(spots []spotmodel.Entity) error {
	reses := make([]responseObject, len(spots))
	for i, spot := range spots {
		res := responseObject{}
		if err := res.Convert(&spot); err != nil {
			return err
		}
		reses[i] = res
	}
	r.Instances = reses
	return nil
}
