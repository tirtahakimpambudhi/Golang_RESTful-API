package helper

import "strconv"

func ParamsParseInt(params string) int {
	result, err := strconv.Atoi(params)
	PanicIFError(err)
	return result
}