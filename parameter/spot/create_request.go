package spot

import (
	"fmt"

	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/library/net/context/me"
)

// CreateRequest ...
type CreateRequest struct {
	UserID   int64     `form:"user_id" json:"user_id"`
	SpotLogs []SpotLog `form:"spot_logs" json:"spot_logs"`
}

// SpotLog ...
type SpotLog struct {
	Lat       float64 `form:"lat" json:"lat"`
	Lng       float64 `form:"lng" json:"lng"`
	VisitTime int64   `form:"visit_time" json:"visit_time"`
}

// NewCreateRequest ...
func NewCreateRequest(ctx echo.Context) (*CreateRequest, error) {
	req := &CreateRequest{}
	if err := ctx.Bind(req); err != nil {
		return nil, err
	}

	fmt.Println(req)

	if user := me.Get(ctx); user != nil {
		req.UserID = user.ID
	}
	if err := req.Validate(); err != nil {
		return nil, err
	}
	return req, nil
}

// Validate ...
func (r *CreateRequest) Validate() error {
	return nil
}
