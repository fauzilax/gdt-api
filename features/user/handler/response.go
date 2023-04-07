package handler

import (
	"gdt-api/features/user"
)

type RegResp struct {
}

func ToRegResp(data user.Core) RegResp {
	return RegResp{}
}

type UserReponse struct {
	ID       uint   `json:"id"`
	Username string `json:"username"`
}

func ToResponse(data user.Core) UserReponse {
	return UserReponse{
		ID:       data.ID,
		Username: data.Username,
	}
}
