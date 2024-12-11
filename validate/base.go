package validate

func GetAllErrorStr(errs map[string]string) string {
	msg := ""
	for _, val := range errs {
		msg += val + ","
	}
	runes := []rune(msg)
	if len(runes) > 0 {
		runes = runes[:len(runes)-1]
	}
	return string(runes)
}
