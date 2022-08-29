// author: asydevc <asydev@163.com>
// date: 2021-03-26

package app

import (
	"encoding/json"
	"errors"
	"reflect"

	i18n "github.com/go-playground/locales/zh"
	i18nTranslator "github.com/go-playground/universal-translator"
	i18nValidator "github.com/go-playground/validator/v10"
	i18nTranslations "github.com/go-playground/validator/v10/translations/zh"
	"github.com/kataras/iris/v12"
)

// Validate
// 验证器实例.
var Validate Validation

// Validation
// 验证器接口.
type Validation interface {
	// Body
	// 校验.
	//
	// 将body数据赋值给ptr, 然后校验ptr数据格式.
	Body(body []byte, ptr interface{}) error

	// Json
	// 校验.
	//
	// 将body数据赋值给ptr, 然后校验ptr数据格式.
	Json(text string, ptr interface{}) error

	// Struct
	// 结构体校验.
	Struct(ptr interface{}) error

	// Iris
	// IRIS RawBody.
	//
	// 将RawBody数据(JSON)赋值给ptr, 然后校验ptr的数据
	// 格式.
	// type Example struct{
	//     Id int `json:"id" validate:"required,min=1" label:"任务ID"`
	// }
	Iris(ctx iris.Context, ptr interface{}) error

	// IrisForm
	// IRIS Form.
	//
	// 将Post表单数据赋值给ptr, 然后校验ptr的数据格式.
	// type Example struct{
	//     Id int `form:"id" validate:"required,min=1" label:"任务ID"`
	// }
	IrisForm(ctx iris.Context, ptr interface{}) error

	// IrisQuery
	// IRIS Query.
	//
	// 将URL中的参数赋值给ptr, 然后校验ptr的数据格式.
	// type Example struct{
	//     Id int `url:"id" validate:"required,min=1" label:"任务ID"`
	// }
	IrisQuery(ctx iris.Context, ptr interface{}) error
}

type validation struct {
	Translator i18nTranslator.Translator
	Validator  *i18nValidator.Validate
}

// Body
// 校验字符串.
// 将Body值赋给ptr, 然后执行Struct校验.
func (o *validation) Body(body []byte, ptr interface{}) error {
	if err := json.Unmarshal(body, ptr); err != nil {
		return err
	}
	return o.Struct(ptr)
}

// Json
// 校验字符串.
// 将Body值赋给ptr, 然后执行Struct校验.
func (o *validation) Json(text string, ptr interface{}) error {
	return o.Body([]byte(text), ptr)
}

// Struct
// 校验结构体.
func (o *validation) Struct(ptr interface{}) error {
	if e0 := o.Validator.Struct(ptr); e0 != nil {
		for _, e1 := range e0.(i18nValidator.ValidationErrors) {
			return errors.New(e1.Translate(o.Translator))
		}
	}
	return nil
}

// Iris
// 校验/IRIS Body.
func (o *validation) Iris(ctx iris.Context, ptr interface{}) error {
	if err := ctx.ReadJSON(ptr); err != nil {
		return err
	}
	return o.Struct(ptr)
}

// IrisForm
// 校验/IRIS Form.
func (o *validation) IrisForm(ctx iris.Context, ptr interface{}) error {
	if err := ctx.ReadForm(ptr); err != nil {
		return err
	}
	return o.Struct(ptr)
}

// IrisQuery
// 校验/IRIS URL.
func (o *validation) IrisQuery(ctx iris.Context, ptr interface{}) error {
	if err := ctx.ReadQuery(ptr); err != nil {
		return err
	}
	return o.Struct(ptr)
}

// 初始化验证器.
func (o *validation) init() *validation {
	o.Validator = i18nValidator.New()

	// 1. 字段名称.
	//    通过反射的字段参数, 读取字段名称.
	o.Validator.RegisterTagNameFunc(func(field reflect.StructField) string {
		for _, k := range []string{"label", "title"} {
			if v := field.Tag.Get(k); v != "" {
				return v
			}
		}
		return field.Name
	})

	// 2. 中文语言.
	//    标准化语言翻译.
	var found bool
	if o.Translator, found = i18nTranslator.New(i18n.New(), i18n.New()).GetTranslator(Config.Lang); found {
		_ = i18nTranslations.RegisterDefaultTranslations(o.Validator, o.Translator)
	}

	// 3. 校验方法.
	//    注册验证器的自定义回调方法.
	validateRegister(o.Validator, o.Translator)
	return o
}
