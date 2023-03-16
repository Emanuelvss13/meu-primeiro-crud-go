package model

type UserDomainInterface interface {
	GetName() string
	GetAge() int8
	GetEmail() string
	GetPassword() string

	ChangeID(string)

	GetJSONValue() (string, error)
	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		ID:       "",
		Name:     name,
		Age:      age,
		Email:    email,
		Password: password,
	}
}
