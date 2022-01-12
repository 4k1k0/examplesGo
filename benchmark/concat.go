package main

import "strings"

func ConcatA(source []string) string {
	var result string
	for _, s := range source {
		result += s
	}
	return result
}

func ConcatB(source []string) string {
	var result strings.Builder
	for _, s := range source {
		_, _ = result.WriteString(s)
	}
	return result.String()
}
