package converter

import (
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model"
	"github.com/Emanuelvss13/meu-primeiro-crud-go/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {

	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}

}
