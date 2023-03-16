package model

type UserDomainInterface interface {
	GetID() string
	GetName() string
	GetAge() int8
	GetEmail() string
	GetPassword() string

	ChangeID(string)

	EncryptPassword()
}

func NewUserDomain(
	email, password, name string,
	age int8,
) UserDomainInterface {
	return &userDomain{
		id:       "",
		name:     name,
		age:      age,
		email:    email,
		password: password,
	}
}
