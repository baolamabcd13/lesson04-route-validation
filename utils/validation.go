package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

func HandleValidationErrors(err error) gin.H {
	if validationError, ok := err.(validator.ValidationErrors); ok {
		errors := make(map[string]string)

		for _, e := range validationError {
			root := strings.Split(e.Namespace(), ".")[0]

			rawPath := strings.TrimPrefix(e.Namespace(), root+".")

			parts := strings.Split(rawPath, ".")

			for i, part := range parts {
				if strings.Contains(part, "[") {
					idx := strings.Index(part, "[")
					base := camelToSnake(part[:idx]) // tu 0 den -> [
					index := part[idx:]
					parts[i] = base + index
				} else {
					parts[i] = camelToSnake(part)
				}
			}

			filepath := strings.Join(parts, ".")

			switch e.Tag() {
			case "gt":
				errors[filepath] = fmt.Sprintf("%s phải lớn hơn %s", filepath, e.Param())
			case "lt":
				errors[filepath] = fmt.Sprintf("%s phải nhỏ hơn %s", filepath, e.Param())
			case "uuid":
				errors[filepath] = fmt.Sprintf("%s phải là uuid hợp lệ", filepath)
			case "slug":
				errors[filepath] = fmt.Sprintf("%s chỉ được chưa chữ thuờng, số, dấu gạch ngang hoặc dấu chấm", filepath)
			case "min":
				errors[filepath] = fmt.Sprintf("%s phải nhiều hơn %s ký tự", filepath, e.Param())
			case "max":
				errors[filepath] = fmt.Sprintf("%s phải ít hơn %s ký tự", filepath, e.Param())
			case "oneof":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[filepath] = fmt.Sprintf("%s phải là một trong các giá trị: %s", filepath, allowedValues)
			case "required":
				errors[filepath] = fmt.Sprintf("%s là bắt buộc", filepath)
			case "search":
				errors[filepath] = fmt.Sprintf("%s chỉ được chưa chữ thuờng, in hoa, số và khoảng trắng", filepath)
			case "gte":
				errors[filepath] = fmt.Sprintf("%s phải lớn hơn hoặc bằng %s", filepath, e.Param())
			case "lte":
				errors[filepath] = fmt.Sprintf("%s phải nhỏ hơn hoặc bằng %s", filepath, e.Param())
			case "email":
				errors[filepath] = fmt.Sprintf("%s phải đúng định dạng là email", filepath)
			case "datetime":
				errors[filepath] = fmt.Sprintf("%s phải theo đúng định dạng YYYY-MM-DD", filepath)
			case "min_int":
				errors[filepath] = fmt.Sprintf("%s phải có giá trị lớn hơn %s", filepath, e.Param())
			case "max_int":
				errors[filepath] = fmt.Sprintf("%s phải có giá trị bé hơn %s", filepath, e.Param())
			case "file_ext":
				allowedValues := strings.Join(strings.Split(e.Param(), " "), ",")
				errors[filepath] = fmt.Sprintf("%s chỉ cho phép nhũng file có extension là: %s", filepath, allowedValues)
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

	v.RegisterValidation("min_int", func(fl validator.FieldLevel) bool {
		minStr := fl.Param()
		//base
		//10: hệ thập phân decimal
		//16: hệ thâp lục phân hex -> FF = 255
		//2: hệ nhị phân binary -> 1010 = 10
		minVal, err := strconv.ParseInt(minStr, 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() >= minVal
	})

	v.RegisterValidation("max_int", func(fl validator.FieldLevel) bool {
		maxStr := fl.Param()
		maxVal, err := strconv.ParseInt(maxStr, 10, 64)
		if err != nil {
			return false
		}
		return fl.Field().Int() <= maxVal
	})

	//file_ext=jpg,png
	v.RegisterValidation("file_ext", func(fl validator.FieldLevel) bool {
		filename := fl.Field().String()
		allowedStr := fl.Param()
		if allowedStr == "" {
			return false
		}

		allowedExt := strings.Fields(allowedStr)
		ext := strings.TrimPrefix(strings.ToLower(filepath.Ext(filename)), ".")

		for _, allowed := range allowedExt {
			if ext == strings.ToLower(allowed) {
				return true
			}
		}
		return false
	})

	return nil
}
