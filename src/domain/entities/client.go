package entities

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"regexp"
)



type CreateClientRequest struct {
	Name string `validate:"notempty"`	
	Birthdate string `validate:"notempty"`
	Email string `validate:"email"`
	Mobilephone string `validate:"notempty"`
	Address string `validate:"notempty"`
	Diseases string
	OtherInfo string
}

type UpdateClientRequest struct {
	Name string
	Birthdate string
	Email string
	Mobilephone string
	Address string
	Diseases string
	OtherInfo string
}

type PaginatedClients struct {
	Clients []Client
	Page int
	Limit int
}
type Client struct {
	ID string `bson:"_id,omitempty"`
	Name string
	Birthdate string
	Email string
	Mobilephone string
	Address string
	Diseases string
	OtherInfo string
}


func isEmailValid(e string) bool {
    emailRegex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
    return emailRegex.MatchString(e)
}

func (ccr *CreateClientRequest) Validate() error {
	validate := validator.New()
	validate.RegisterValidation("notempty", func(fl validator.FieldLevel) bool {
		return fl.Field().String() != ""
	})

	validate.RegisterValidation("email", func(fl validator.FieldLevel) bool {
		return isEmailValid(fl.Field().String())
	})

	

	if err := validate.Struct(ccr); err != nil {
		return errors.New(err.Error())
	}
	return nil
}
