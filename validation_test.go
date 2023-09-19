package golangvalidation

import (
	"fmt"
	"testing"

	"github.com/go-playground/validator/v10"
)

func TestValidation(t *testing.T) {
	var validate = validator.New()
	if validate == nil {
		t.Error("validate is nill")
	}
}

func TestValidationVariable(t *testing.T) {
	valdate := validator.New()
	user := "aku"
	err := valdate.Var(user, "required")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestValidation2Variable(t *testing.T) {
	valdate := validator.New()
	password := "main"
	confirm := "main"
	err := valdate.VarWithValue(password, confirm, "eqfield")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestMultipleTagValidation(t *testing.T) {
	valdate := validator.New()
	userdata := "adit2369"
	err := valdate.Var(userdata, "required,number")
	if err != nil {
		fmt.Println(err.Error())
	}
}

func TestTagParameter(t *testing.T) {
	valdate := validator.New()
	userdata := "993322"
	err := valdate.Var(userdata, "required,numeric,min=5,max=10")
	if err != nil {
		fmt.Println(err.Error())
	}
}
