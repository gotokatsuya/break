package tour

import (
	"strconv"

	"github.com/labstack/echo"

	libmath "github.com/gotokatsuya/break/library/math"
)

// GetRequest ...
type GetRequest struct {
	UserID int64   `json:"user_id"`
	Lat    float64 `json:"lat"`
	Lng    float64 `json:"lng"`
}

// NewGetRequest ...
func NewGetRequest(ctx echo.Context) (*GetRequest, error) {
	req := &GetRequest{}
	userIDStr := ctx.QueryParam("user_id")
	switch {
	case len(userIDStr) != 0:
		userID, err := strconv.Atoi(userIDStr)
		if err != nil {
			return nil, err
		}
		req.UserID = int64(userID)
	default:
		lat, lng, err := libmath.ParseLatLng(ctx.QueryParam("lat"), ctx.QueryParam("lng"))
		if err != nil {
			return nil, err
		}
		lat, lng, err = libmath.RoundLatLng(lat, lng)
		if err != nil {
			return nil, err
		}
		req.Lat = lat
		req.Lng = lng
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return req, nil
}

func (r *GetRequest) HasUserID() bool {
	return r.UserID != 0
}

// Validate ...
func (r *GetRequest) Validate() error {
	return nil
}
