package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	v "github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	"strconv"
	"strings"
	"unicode"
)

const minPasswordLength = 12

type CustomValidator struct {
	validator *v.Validate
	uni       *ut.UniversalTranslator
	trans     ut.Translator
}

func NewCustomValidator(validator *v.Validate) *CustomValidator {

	cv := &CustomValidator{validator: validator}

	translator := en.New()
	cv.uni = ut.New(translator, translator)
	trans, _ := cv.uni.GetTranslator("translator")
	cv.trans = trans

	_ = enTranslations.RegisterDefaultTranslations(validator, trans)

	// Add the custom rules in for validation
	cv.addCustomRules()

	return cv
}

func (cv *CustomValidator) Validate(i interface{}) error {
	err := cv.validator.Struct(i)
	if err == nil {
		return nil
	}
	var errs v.ValidationErrors
	errors.As(err, &errs)
	msg := ""
	for _, valErrorTranslation := range errs.Translate(cv.trans) {
		if msg != "" {
			msg += ", "
		}
		msg += valErrorTranslation
	}
	return errors.New(msg)
}

// addCustomRules adds the custom rules to the validator
func (cv *CustomValidator) addCustomRules() {

	if err := cv.validator.RegisterValidation("passLen", ValidatePasswordLength); err != nil {
		panic(err)
	}
	if err := cv.validator.RegisterValidation("passComplexity", ValidatePasswordComplexity); err != nil {
		panic(err)
	}
	if err := cv.validator.RegisterValidation("maxFieldSize", ValidateMaxFieldSize); err != nil {
		panic(err)
	} // Used to check the total length of fields like lists or maps after they're marshalled
	if err := cv.validator.RegisterValidation("appPackageName", ValidateAppPackageName); err != nil {
		panic(err)
	}
	if err := cv.validator.RegisterValidation("commaSeparatedDomains", ValidateCommaSeparatedDomains); err != nil {
		panic(err)
	}
	if err := cv.validator.RegisterValidation("proxyExclude", ValidateProxyExcludeList); err != nil {
		panic(err)
	}

	// Add translation
	cv.addTranslation(
		"passLen",
		fmt.Sprintf("{0} must be at least %v characters long", minPasswordLength),
	)
	cv.addTranslation(
		"passComplexity",
		"{0} must contains numbers, special characters and upper and lower case letters",
	)
	cv.addTranslation(
		"maxFieldSize", "{0} must be a maximum of {1} characters in length",
	)
	cv.addTranslation(
		"appPackageName", "{0} can't contain special characters apart from '.' (except for madepurpleos.com packages)",
	)
	cv.addTranslation(
		"commaSeparatedDomains", "{0} can't contain spaces or leading/trailing/successive commas",
	)
	cv.addTranslation(
		"proxyExclude", "{0} can't contain spaces or commas",
	)
}

func ValidatePasswordLength(fl v.FieldLevel) bool {
	return len(fl.Field().String()) >= minPasswordLength
}

func ValidatePasswordComplexity(fl v.FieldLevel) bool {

	password := fl.Field().String()
	var (
		hasUpper   = false
		hasLower   = false
		hasNumber  = false
		hasSpecial = false
	)

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUpper = true
		case unicode.IsLower(char):
			hasLower = true
		case unicode.IsNumber(char):
			hasNumber = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecial = true
		}
	}

	return hasUpper && hasLower && hasNumber && hasSpecial

}

func ValidateMaxFieldSize(fl v.FieldLevel) bool {
	maxStr := fl.Param()
	maxSize, err := strconv.Atoi(maxStr)
	if err != nil {
		return false
	}
	val, err := json.Marshal(fl.Field().Interface())
	if err != nil {
		return false
	}
	return len(string(val)) <= maxSize
}

func ValidateAppPackageName(fl v.FieldLevel) bool {

	packageName := fl.Field().String()

	if strings.HasPrefix(packageName, "madepurpleos.com") {
		return true
	}

	for _, char := range packageName {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '.' {
			return false
		}
	}

	return true
}

func ValidateCommaSeparatedDomains(fl v.FieldLevel) bool {

	domains := fl.Field().String()

	if strings.HasPrefix(domains, ",") ||
		strings.HasSuffix(domains, ",") ||
		strings.Contains(domains, " ") ||
		strings.Contains(domains, ",,") {
		return false
	}

	return true
}

func ValidateProxyExcludeList(fl v.FieldLevel) bool {

	excludeList := fl.Field().String()

	if strings.Contains(excludeList, " ") ||
		strings.Contains(excludeList, ",") {
		return false
	}

	return true
}

func (cv *CustomValidator) addTranslation(tag string, errMessage string) {
	registerFn := func(ut ut.Translator) error {
		return ut.Add(tag, errMessage, false)
	}
	transFn := func(ut ut.Translator, fe v.FieldError) string {
		param := fe.Param()
		tag := fe.Tag()

		t, err := ut.T(tag, fe.Field(), param)
		if err != nil {
			return fe.(error).Error()
		}
		return t
	}
	_ = cv.validator.RegisterTranslation(tag, cv.trans, registerFn, transFn)
}
