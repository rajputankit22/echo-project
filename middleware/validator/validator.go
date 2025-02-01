package validator

import (
	"echo-project/constant"
	"strings"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTraslations "github.com/go-playground/validator/v10/translations/en"
)

// ValidateInterface is an interface for the Validate function
type ValidateInterface interface {
	Validate(input interface{}) error
}

// Validator struct
type Validator struct {
	validator *validator.Validate
	trans     ut.Translator
}

// ValidationError struct
type ValidationError struct {
	Err    error
	Msg    string
	Fields map[string]string
}

// Construnctor function for Validator
func NewValidatorAdapter() ValidateInterface {
	localTrans := en.New()
	uni := ut.New(localTrans, localTrans)
	trans, _ := uni.GetTranslator("en")

	validatorObj := &Validator{
		validator: validator.New(),
		trans:     trans,
	}

	setupCustomValidator(validatorObj)
	setupCustomMessages(validatorObj)

	return validatorObj
}

// Validate function validates the input
func (v *Validator) Validate(input interface{}) error {
	err := v.validator.Struct(input)
	if err == nil {
		return nil
	}
	return v.newValidatorError(err)
}

// Error function returns the error message
func (ve *ValidationError) Error() string {
	return ve.Msg
}

// GetFields function returns the fields
func (v *Validator) newValidatorError(err error) *ValidationError {
	if err == nil {

		return &ValidationError{
			Err: nil,
			Msg: "success",
		}
	}
	switch err.(type) {
	case validator.ValidationErrors:
		errs, _ := err.(validator.ValidationErrors)
		return v.formValidatorErrors(errs)
	default:
		return &ValidationError{
			Err: err,
			Msg: err.Error(),
		}
	}
}

// formValidatorErrors function returns the form validation errors
func (v *Validator) formValidatorErrors(errs validator.ValidationErrors) *ValidationError {
	var (
		msg       string
		fieldErrs map[string]string
	)
	fieldErrs = make(map[string]string)

	for _, e := range errs {
		fieldErrs[strings.ToLower(e.Field())] = strings.ToLower(e.Translate(v.trans))
		msg = fieldErrs[strings.ToLower(e.Field())]
	}

	ve := &ValidationError{
		Err:    errs,
		Msg:    msg,
		Fields: fieldErrs,
	}
	return ve
}

// setupCustomValidator function sets up the custom validator
func setupCustomValidator(validatorObj *Validator) {
	validatorObj.validator.RegisterValidation("customValidation", validatorObj.customValidation)
}

// setupCustomMessages function sets up the custom messages
func setupCustomMessages(validatorObj *Validator) {
	enTraslations.RegisterDefaultTranslations(validatorObj.validator, validatorObj.trans)
	registerTranslation(validatorObj, "required", constant.RequiredValidateMessage)
	registerTranslation(validatorObj, "customValidation", constant.RequiredValidateMessage)
}

// registerTranslation function registers the translation
func registerTranslation(validatorObj *Validator, tag string, message string) {
	validatorObj.validator.RegisterTranslation(tag, validatorObj.trans, func(ut ut.Translator) error {
		return ut.Add(tag, message, true)
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T(tag, fe.Field(), fe.Param())
		return t
	})
}

// customValidation function validates the custom validation
func (v *Validator) customValidation(fl validator.FieldLevel) bool {
	str, ok := fl.Field().Interface().(string)
	if !ok {
		return false
	}
	if len(str) == 0 {
		return true
	}
	return false
}
