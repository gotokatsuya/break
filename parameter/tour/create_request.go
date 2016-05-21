package tour

import (
	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/library/net/context/me"
)

// CreateRequest ...
type CreateRequest struct {
	Name     string  `form:"name" json:"name"`
	PhotoURL string  `form:"photo_url" json:"photo_url"`
	UserID   int64   `form:"user_id" json:"user_id"`
	SpotIDs  []int64 `form:"spot_ids" json:"spot_ids"`
}

// NewCreateRequest ...
func NewCreateRequest(ctx echo.Context) (*CreateRequest, error) {
	req := &CreateRequest{}
	if err := ctx.Bind(req); err != nil {
		return nil, err
	}
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
