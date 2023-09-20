package golangvalidation

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"testing"

	"github.com/go-playground/validator/v10"
)

func MustValidUsername(field validator.FieldLevel) bool {
	value, ok := field.Field().Interface().(string)

	if ok {
		if value != strings.ToUpper(value) {
			return false
		}
		if len(value) < 5 {
			return false
		}
	}
	return true
}
func TestCustomValidation(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("username", MustValidUsername)

	user := "AKUAJA"
	err := validate.Var(user, "required,username")
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}

}

var regexNumber = regexp.MustCompile("^[0-9]+$")

func MustValidPin(field validator.FieldLevel) bool {
	length, err := strconv.Atoi(field.Param())
	if err != nil {
		panic(err)
	}
	value := field.Field().String()

	if !regexNumber.MatchString(value) {
		return false
	}

	return len(value) == length
}

func TestCustomPin(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("pin", MustValidPin)

	pin := "2178126782"
	err := validate.Var(pin, "required,pin=10")
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}

}

func TestOrRule(t *testing.T) {
	validate := validator.New()

	username := "aditya@gmail.com"
	err := validate.Var(username, "required,email|numeric")
	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}

}

func MustEqualIgnoreCase(field validator.FieldLevel) bool {
	value, _, _, ok := field.GetStructFieldOK2()
	if !ok {
		panic("field not ok")
	}
	firstValue := strings.ToUpper(field.Field().String())
	secondValue := strings.ToUpper(value.String())

	return firstValue == secondValue
}

func TestCustomStructValidate(t *testing.T) {
	validate := validator.New()
	validate.RegisterValidation("field_equal_ignore_case", MustEqualIgnoreCase)

	type User struct {
		Username string `validate:"required,field_equal_ignore_case=Email|field_equal_ignore_case=Phone"`
		Email    string `validate:"required,email"`
		Phone    string `validate:"required,numeric"`
		Name     string `validate:"required"`
	}

	user := User{
		Username: "09163871692876",
		Email:    "aditya27@gmail.com",
		Phone:    "09163871692876",
		Name:     "Aditya",
	}
	err := validate.Struct(user)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}

}

type Register struct {
	Username string `validate:"required"`
	Email    string `validate:"required,email"`
	Phone    string `validate:"required,numeric"`
	Password string `validate:"required"`
}

func MustValidRegisterSucces(level validator.StructLevel) {
	registerRequest := level.Current().Interface().(Register)
	if registerRequest.Username == registerRequest.Email || registerRequest.Username == registerRequest.Phone {
		// Succes
	} else {
		// Gagal
		level.ReportError(registerRequest.Username, "Username", "Username", "username", "")
	}

}

func TestCustomStructLevel(t *testing.T) {
	validate := validator.New()
	validate.RegisterStructValidation(MustValidRegisterSucces, Register{})

	register := Register{
		Username: "09163871692876",
		Email:    "aditya27@gmail.com",
		Phone:    "09163871692876",
		Password: "Rahasia",
	}
	err := validate.Struct(register)

	if err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, fieldError := range validationErrors {
			fmt.Println("Error", fieldError.Field(), "on Tag", fieldError.Tag(), "with error", fieldError.Error())

		}
	}

}
