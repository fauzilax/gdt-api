package user

import (
	"github.com/labstack/echo/v4"
)

type Core struct {
	ID       uint
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserHandler interface {
	Register() echo.HandlerFunc
	Login() echo.HandlerFunc
}

type UserService interface {
	Register(registerData Core) (Core, error)
	Login(username string, password string) (string, Core, error)
}

type UserData interface {
	Register(registerData Core) (Core, error)
	Login(username string) (Core, error)
}
