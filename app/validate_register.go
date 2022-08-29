// author: asydevc <asydev@163.com>
// date: 2022-01-05

package app

import (
	"regexp"

	"github.com/asydevc/log/v2"
	i18nTranslator "github.com/go-playground/universal-translator"
	i18nValidator "github.com/go-playground/validator/v10"
)

// 注册校验.
func validateRegister(v *i18nValidator.Validate, t i18nTranslator.Translator) {
	r1 := regexp.MustCompile(`^[_a-zA-Z][_a-zA-Z0-9]*$`)

	for _, x := range []struct {
		override    bool
		tag         string
		translation string
		validate    i18nValidator.Func
	}{
		{
			tag: "var", translation: "{0}不合法", override: true,
			validate: func(f i18nValidator.FieldLevel) bool { return r1.MatchString(f.Field().String()) },
		},
	} {
		if err := v.RegisterValidation(x.tag, x.validate); err != nil {
			log.Errorf("register validate func fail, tag=%s, error=%v.", x.tag, err)
		} else {
			if err = v.RegisterTranslation(
				x.tag, t, func(ut i18nTranslator.Translator) error { return ut.Add(x.tag, x.translation, x.override) }, func(ut i18nTranslator.Translator, fe i18nValidator.FieldError) string {
					us, ue := ut.T(fe.Tag(), fe.Field())
					if ue != nil {
						return fe.(error).Error()
					}
					return us
				},
			); err != nil {
				log.Errorf("register validate translation fail, tag=%s, error=%v.", x.tag, err)
			}
		}
	}
}
