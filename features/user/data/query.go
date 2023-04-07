package data

import (
	"errors"
	"gdt-api/features/user"
	"log"

	"gorm.io/gorm"
)

type userQuery struct {
	db *gorm.DB
}

func New(db *gorm.DB) user.UserData {
	return &userQuery{
		db: db,
	}
}

// Register implements user.UserData
func (uq *userQuery) Register(registerData user.Core) (user.Core, error) {
	cnv := CoreToData(registerData)
	err := uq.db.Create(&cnv).Error
	if err != nil {
		log.Println("query error", err.Error())
		return user.Core{}, errors.New("server error")
	}

	registerData.ID = cnv.ID
	return registerData, nil
}

// Login implements user.UserData
func (uq *userQuery) Login(username string) (user.Core, error) {
	if username == "" {
		log.Println("data empty, query error")
		return user.Core{}, errors.New("username not allowed empty")
	}
	res := User{}
	if err := uq.db.Where("username = ?", username).First(&res).Error; err != nil {
		log.Println("login query error", err.Error())
		return user.Core{}, errors.New("data not found")
	}

	return DataToCore(res), nil
}
