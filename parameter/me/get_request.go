package me

import (
	"github.com/labstack/echo"
)

// GetRequest ...
type GetRequest struct {
}

// NewGetRequest ...
func NewGetRequest(ctx echo.Context) (*GetRequest, error) {
	req := &GetRequest{}
	return req, nil
}
