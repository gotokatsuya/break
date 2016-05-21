package me

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	"github.com/gotokatsuya/break/library/net/context/me"
	usermodel "github.com/gotokatsuya/break/model/user"
	parameter "github.com/gotokatsuya/break/parameter/me"
)

func getNewUser(ctx echo.Context, name, email, photoURL string) (*usermodel.Entity, error) {
	var (
		userRepository = usermodel.NewRepository(ctx)
		user           *usermodel.Entity
		err            error
	)
	user = usermodel.New(name, email, photoURL)
	if user, err = userRepository.Create(user); err != nil {
		return nil, err
	}
	if err = user.RenewToken(); err != nil {
		return nil, err
	}
	if user, err = userRepository.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

// LoginMe login or register
func LoginMe(ctx echo.Context, req *parameter.LoginRequest) (*parameter.LoginResponse, error) {
	var (
		userRepository = usermodel.NewRepository(ctx)
		user, err      = userRepository.GetByEmail(req.Email)
	)
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	switch {
	case user == nil:
		// Register
		user, err = getNewUser(ctx, req.Name, req.Email, req.PhotoURL)
	}
	if err != nil {
		return nil, err
	}
	res := parameter.NewLoginResponse()
	if err := res.Convert(user); err != nil {
		return nil, err
	}
	return res, nil
}

func GetMe(ctx echo.Context, req *parameter.GetRequest) (*parameter.GetResponse, error) {
	user := me.Get(ctx)
	res := parameter.NewGetResponse()
	if err := res.Convert(user); err != nil {
		return nil, err
	}
	return res, nil
}
