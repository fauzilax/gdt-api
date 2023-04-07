package data

import (
	"gdt-api/features/user"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
}

func DataToCore(data User) user.Core {
	return user.Core{
		ID:       data.ID,
		Username: data.Username,
		Password: data.Password,
	}
}
func CoreToData(core user.Core) User {
	return User{
		Model:    gorm.Model{ID: core.ID},
		Username: core.Username,
		Password: core.Password,
	}
}
