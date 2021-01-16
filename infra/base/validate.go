package base

import (
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
	log "github.com/sirupsen/logrus"
	"goEnvelope/infra"
)

var validate *validator.Validate
var translator ut.Translator

func Validate() *validator.Validate {
	return validate
}

func Translator() ut.Translator {
	return translator
}

type ValidateStarter struct {
	infra.BaseStarter
}

func (v *ValidateStarter) Init(ctx infra.StarterContext) {
	ZH := zh.New()
	uni := ut.New(ZH, ZH)
	translator, found := uni.GetTranslator("zh")
	validate := validator.New()
	if found {
		err := zh_translations.RegisterDefaultTranslations(validate, translator)
		if err != nil {
			log.Error("RegisterDefaultTranslations Error: ", err)
			return
		}
	} else {
		log.Error("translator zh not found")
		return
	}
	log.Info("Validator init success")
}
