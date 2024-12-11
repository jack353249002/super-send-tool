package validate

import "github.com/go-playground/validator/v10"

func SuperSendErrorMessages(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			switch fieldErr.Field() {
			case "Online":
				if fieldErr.Tag() == "required" {
					errors["online"] = "Online 不能为空"
				}
			case "ID":
				if fieldErr.Tag() == "required" {
					errors["address"] = "ID 不能为空"
				}
			case "Address":
				if fieldErr.Tag() == "required" {
					errors["address"] = "地址不能为空"
				}
			case "isSSL":
				if fieldErr.Tag() == "required" {
					errors["isSSL"] = "是否开启ssl不能为空"
				}
			case "UserName":
				if fieldErr.Tag() == "required" {
					errors["fieldErr"] = "用户名不能为空"
				}
			case "Password":
				if fieldErr.Tag() == "required" {
					errors["Password"] = "密码不能为空"
				}
			}
		}
	}
	return errors
}
