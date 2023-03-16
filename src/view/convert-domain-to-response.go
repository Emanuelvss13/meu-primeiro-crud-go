package view

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/controller/model/response"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
)

func ConvetDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
