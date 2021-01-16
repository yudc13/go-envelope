package main

import (
	"fmt"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	log "github.com/sirupsen/logrus"
)

type Person struct {
	FirstName string `validate:"required"`
	LastName  string `validate:"required"`
	Age       int    `validate:"gte=0,lte=100"`
	Email     string `validate:"required,email"`
}

func main() {
	person := &Person{
		FirstName: "firstName",
		LastName:  "lastName",
		Age:       130,
		Email:     "jackqq.com",
	}
	ZH := zh.New()
	uni := ut.New(ZH, ZH)
	trans, _ := uni.GetTranslator("zh")
	validate := validator.New()
	err := zh_translations.RegisterDefaultTranslations(validate, trans)
	if err != nil {
		log.Fatal("RegisterDefaultTranslations Error: ", err)
	}
	err = validate.Struct(person)
	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			log.Errorf("validate person error: %v \n", err)
			return
		}
		for _, err := range err.(validator.ValidationErrors) {
			log.Errorf("validate error: %v \n", err.Translate(trans))
			fmt.Println(err.Namespace())
			fmt.Println(err.Field())
			fmt.Println(err.StructNamespace())
			fmt.Println(err.StructField())
			fmt.Println(err.Tag())
			fmt.Println(err.ActualTag())
			fmt.Println(err.Type())
			fmt.Println(err.Param())
		}
		return
	}
	log.Info("validate success")
}
