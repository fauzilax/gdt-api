package services

import (
	"errors"
	"gdt-api/features/user"
	"gdt-api/helper"
	"log"
	"strings"
)

type userUseCase struct {
	qry user.UserData
}

func New(ud user.UserData) user.UserService {
	return &userUseCase{
		qry: ud,
	}
}

// Register implements user.UserService
func (uuc *userUseCase) Register(registerData user.Core) (user.Core, error) {

	//validation
	err := helper.RegistrationValidate(registerData)
	if err != nil {
		return user.Core{}, errors.New("validate: " + err.Error())
	}

	hashed := helper.GeneratePassword(registerData.Password)
	registerData.Password = string(hashed)

	res, err := uuc.qry.Register(registerData)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "duplicated") {
			msg = "username already registered"
		} else {
			msg = "server error"
		}
		return user.Core{}, errors.New(msg)
	}

	return res, nil
}

// Login implements user.UserService
func (uuc *userUseCase) Login(username string, password string) (string, user.Core, error) {
	res, err := uuc.qry.Login(username)
	if err != nil {
		msg := ""
		if strings.Contains(err.Error(), "empty") {
			msg = "username or password not allowed empty"
		} else {
			msg = "account not registered or server error"
		}
		return "", user.Core{}, errors.New(msg)
	}
	if err := helper.ComparePassword(res.Password, password); err != nil {
		log.Println("login compare", err.Error())
		return "", user.Core{}, errors.New("password not matched")
	}
	useToken := helper.GenerateToken(int(res.ID))

	return useToken, res, nil
}
