package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationError {
			switch e.Tag() {
			case "gt":
				errors[e.Field()] = e.Field() + " phải lớn hơn giá trị tối thiểu"
			case "uuid":
				errors[e.Field()] = e.Field() + " phải là uuid hợp lệ"
			case "slug":
				errors[e.Field()] = e.Field() + " chỉ được chưa chữ thuờng, số, dấu gạch ngang hoặc dấu chấm"
			case "min":
				errors[e.Field()] = fmt.Sprintf("%s phải nhiều hơn %s ký tự", e.Field(), e.Param())
			case "max":
				errors[e.Field()] = fmt.Sprintf("%s phải ít hơn %s ký tự", e.Field(), e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[e.Field()] = fmt.Sprintf("%s phải là một trong các giá trị: %s", e.Field(), allowedValues)
			case "required":
				errors[e.Field()] = e.Field() + " là bắt buộc"
			case "search":
				errors[e.Field()] = e.Field() + " chỉ được chưa chữ thuờng, in hoa, số và khoảng trắng"
			case "gte":
				errors[e.Field()] = e.Field() + " phải lớn hơn hoặc bằng giá trị tối thiêu"
			case "lte":
				errors[e.Field()] = e.Field() + " phải nhỏ hơn hoặc bằng giá trị tối thiêu"
			case "email":
				errors[e.Field()] = e.Field() + " phải đúng định dạng là email"
			case "datetime":
				errors[e.Field()] = e.Field() + " phải theo đúng định dạng YYYY-MM-DD"
			}
		}

		return gin.H{"error": errors}

	}
	return gin.H{"error": "Yêu cầu không hợp lệ" + err.Error()}
}

func RegisterValidators() error {
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if !ok {
		return fmt.Errorf("failed to get validator engine")
	}

	var slugRegex = regexp.MustCompile(`^[a-z0-9]+(?:[-.][a-z0-9]+)*$`)
	v.RegisterValidation("slug", func(fl validator.FieldLevel) bool {
		return slugRegex.MatchString(fl.Field().String())
	})

	var searchgRegex = regexp.MustCompile(`^[a-zA-Z0-9\s]+$`)
	v.RegisterValidation("search", func(fl validator.FieldLevel) bool {
		return searchgRegex.MatchString(fl.Field().String())
	})
	return nil
}
