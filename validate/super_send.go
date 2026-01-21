package validate

import "github.com/go-playground/validator/v10"

func SuperSendErrorMessages(err error) map[string]string {
	errors := make(map[string]string)
	if validationErrors, ok := err.(validator.ValidationErrors); ok {
		for _, fieldErr := range validationErrors {
			switch fieldErr.Field() {
			case "SuperSendID":
				if fieldErr.Tag() == "required" {
					errors["superSendID"] = "SuperSendID 不能为空"
				}
			case "Token":
				if fieldErr.Tag() == "required" {
					errors["token"] = "token 不能为空"
				}
			case "Online":
				if fieldErr.Tag() == "required" {
					errors["online"] = "Online 不能为空"
				}
			case "ID":
				if fieldErr.Tag() == "required" {
					errors["id"] = "ID 不能为空"
				}
			case "Address":
				if fieldErr.Tag() == "required" {
					errors["address"] = "地址不能为空"
				}
			case "IsSSL":
				if fieldErr.Tag() == "required" {
					errors["isSSL"] = "是否开启ssl不能为空"
				}
			case "UserName":
				if fieldErr.Tag() == "required" {
					errors["UserName"] = "用户名不能为空"
				}
			case "Password":
				if fieldErr.Tag() == "required" {
					errors["Password"] = "密码不能为空"
				}
			case "Title":
				if fieldErr.Tag() == "required" {
					errors["Title"] = "标题不能为空"
				}
			case "Body":
				if fieldErr.Tag() == "required" {
					errors["Body"] = "内容不能为空"
				}
			case "SendID":
				if fieldErr.Tag() == "required" {
					errors["send_id"] = "发送id不能为空"
				}
			case "SendServerID":
				if fieldErr.Tag() == "required" {
					errors["send_server_id"] = "发送服务器id不能为空"
				}
			case "TimeZone":
				if fieldErr.Tag() == "required" {
					errors["time_zone"] = "时区不能为空"
				}

			}

		}
	}
	return errors
}
