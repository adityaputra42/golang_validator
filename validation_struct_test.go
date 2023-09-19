package golangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

type LoginRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required,min=5"`
}

type RegisterRequest struct {
	Username string `validate:"required,email"`
	Password string `validate:"required,min=5"`
	Confirm  string `validate:"required,min=5,eqfield=Password"`
}

func TestValidationStruct(t *testing.T) {
	valdate := validator.New()
	loginRequest := LoginRequest{
		Username: "adit27@gmail.com",
		Password: "jagajaga",
	}
	err := valdate.Struct(loginRequest)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidationError(t *testing.T) {
	valdate := validator.New()
	loginRequest := LoginRequest{
		Username: "adit27",
		Password: "jaga",
	}
	err := valdate.Struct(loginRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}
}

func TestValidationCrossField(t *testing.T) {
	valdate := validator.New()
	registerRequest := RegisterRequest{
		Username: "adit27@gmail.com",
		Password: "jagalagi",
		Confirm:  "jagalag",
	}
	err := valdate.Struct(registerRequest)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}
}

type School struct {
	Name string `validate:"required"`
}
type Address struct {
	City   string `validate:"required"`
	Street string `validate:"required"`
}

type User struct {
	Id        string            `validate:"required"`
	Name      string            `validate:"required"`
	Addresses []Address         `validate:"required,dive"`
	Hobbies   []string          `validate:"dive,required,min=3"`
	Schools   map[string]School `validate:"dive,keys,required,min=2,endkeys"`
	Wallet    map[string]int    `validate:"dive,keys,required,endkeys,required,gt=0"`
}

func TestValidationNestedStruct(t *testing.T) {
	valdate := validator.New()
	user := User{
		Id:   "1",
		Name: "Aditya",
		Addresses: []Address{
			{
				City:   "City",
				Street: "Sempor baru",
			}, {
				City:   "Semarang",
				Street: "Gendong raya",
			},
		},
	}
	err := valdate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}
}

func TestBaseCollection(t *testing.T) {
	valdate := validator.New()
	user := User{
		Id:   "1",
		Name: "Aditya",
		Addresses: []Address{
			{
				City:   "City",
				Street: "Sempor baru",
			}, {
				City:   "Semarang",
				Street: "Gendong raya",
			},
		},
		Hobbies: []string{
			"football",
			"coding",
		},
	}
	err := valdate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}
}

func TestBaseMap(t *testing.T) {
	valdate := validator.New()
	user := User{
		Id:   "1",
		Name: "Aditya",
		Addresses: []Address{
			{
				City:   "City",
				Street: "Sempor baru",
			}, {
				City:   "Semarang",
				Street: "Gendong raya",
			},
		},
		Hobbies: []string{
			"football",
			"coding",
		},
		Schools: map[string]School{
			"SD":  {Name: "SD N 1 Semarang"},
			"SMP": {Name: "SMP N 2 Semarang"},
			"SMK": {Name: "SMK Ma'arif 2 Semarang"},
		},
	}
	err := valdate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}
}

func TestBasicMap(t *testing.T) {
	valdate := validator.New()
	user := User{
		Id:   "1",
		Name: "Aditya",
		Addresses: []Address{
			{
				City:   "City",
				Street: "Sempor baru",
			}, {
				City:   "Semarang",
				Street: "Gendong raya",
			},
		},
		Hobbies: []string{
			"football",
			"coding",
		},
		Schools: map[string]School{
			"SD":  {Name: "SD N 1 Semarang"},
			"SMP": {Name: "SMP N 2 Semarang"},
			"SMK": {Name: "SMK Ma'arif 2 Semarang"},
		},
		Wallet: map[string]int{
			"balance": 0,
		},
	}
	err := valdate.Struct(user)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}
}

func TestTagAlias(t *testing.T) {
	validate := validator.New()
	validate.RegisterAlias("varchar", "required,max=255")

	type Seller struct {
		Id     string `validate:"varchar"`
		Name   string `validate:"varchar"`
		Owner  string `validate:"varchar"`
		Slogan string `validate:"varchar"`
	}

	seller := Seller{
		Id:     "1",
		Name:   "Paijo",
		Owner:  "Mail",
		Slogan: "Prediksi Jaya Jaya Jaya",
	}
	err := validate.Struct(seller)
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}
}
