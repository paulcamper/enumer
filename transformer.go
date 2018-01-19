package main

import (
	"strings"

	"unicode/utf8"

	"github.com/fatih/camelcase"
)

func transform(src, delim string) string {
	entries := camelcase.Split(src)
	if len(entries) <= 1 {
		return strings.ToLower(src)
	}

	result := strings.ToLower(entries[0])
	for i := 1; i < len(entries); i++ {
		result += delim + strings.ToLower(entries[i])
	}
	return result
}

func toCamelCaseLower(src string) string {
	f := firstLower(src)

	return string(f + src[1:])
}

func toSnakeCase(src string) string {
	return transform(src, "_")
}

func toSnakeCaseUpper(src string) string {
	return strings.ToUpper(toSnakeCase(src))
}

func toKebabCase(src string) string {
	return transform(src, "-")
}

func toKebabCaseUpper(src string) string {
	return strings.ToUpper(toKebabCase(src))
}

func first(src string) string {
	r, _ := utf8.DecodeRuneInString(src)
	return string(r)
}

func firstUpper(src string) string {
	return strings.ToUpper(first(src))
}

func firstLower(src string) string {
	return strings.ToLower(first(src))
}
