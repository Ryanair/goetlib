package http

import "strings"

func Clean(rawBody string) string {
	str_ := strings.ReplaceAll(rawBody, "\n", "")
	str := strings.ReplaceAll(str_, "\t", "")
	return str
}
