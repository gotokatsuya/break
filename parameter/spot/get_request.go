package spot

import (
	"strconv"

	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/library/net/context/me"
)

// GetRequest ...
type GetRequest struct {
	TourID         int64 `json:"tour_id"`
	UserID         int64 `json:"user_id"`
	StartVisitTime int64 `json:"start_visit_time"`
	EndVisitTime   int64 `json:"end_visit_time"`
}

// NewGetRequest ...
func NewGetRequest(ctx echo.Context) (*GetRequest, error) {
	req := &GetRequest{}
	tourIDStr := ctx.QueryParam("tour_id")
	switch {
	case len(tourIDStr) != 0:
		tourID, err := strconv.Atoi(tourIDStr)
		if err != nil {
			return nil, err
		}
		req.TourID = int64(tourID)
	default:
		startVisitTime, err := strconv.Atoi(ctx.QueryParam("start_visit_time"))
		if err != nil {
			return nil, err
		}
		endVisitTime, err := strconv.Atoi(ctx.QueryParam("end_visit_time"))
		if err != nil {
			return nil, err
		}
		req.StartVisitTime = int64(startVisitTime)
		req.EndVisitTime = int64(endVisitTime)
	}
	if user := me.Get(ctx); user != nil {
		req.UserID = user.ID
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return req, nil
}

// HasTourID ...
func (r *GetRequest) HasTourID() bool {
	return r.TourID != 0
}

// Validate ...
func (r *GetRequest) Validate() error {
	return nil
}
