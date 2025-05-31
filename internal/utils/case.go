package utils

import (
	"regexp"
	"strings"

	"github.com/iancoleman/strcase"
)

type Case int

const (
	UpperCase Case = iota
	LowerCase
	CamelCase
	SnakeCase
	PascalCase
	KebabCase
	SentenceCase
)

var CaseStringToCase = map[string]Case{
	"upper-case": UpperCase,
	"uppercase":  UpperCase,

	"lower-case": LowerCase,
	"lowercase":  LowerCase,

	"camel-case": CamelCase,
	"camelcase":  CamelCase,

	"snake-case": SnakeCase,
	"snakecase":  SnakeCase,

	"pascal-case": PascalCase,
	"pascalcase":  PascalCase,

	"kebab-case": KebabCase,
	"kebabcase":  KebabCase,

	"sentence-case": SentenceCase,
	"sentencecase":  SentenceCase,
}

func UpperCaseFirst(input string) string {
	first := input[0:1]
	rest := input[1:]
	return strings.ToUpper(first) + rest
}

func ToCase(input string, caseType Case) string {
	switch caseType {
	case UpperCase:
		return strings.ToUpper(input)
	case LowerCase:
		return strings.ToLower(input)
	case CamelCase:
		return strcase.ToLowerCamel(input)
	case SnakeCase:
		return strcase.ToSnake(input)
	case PascalCase:
		return strcase.ToCamel(input)
	case KebabCase:
		return strcase.ToKebab(input)
	case SentenceCase:
		return UpperCaseFirst(input)
	}

	return input
}

var replacer = strings.NewReplacer("'", "", "\"", "", "`", "")

func EnsureCase(input string, caseType Case) bool {
	input = replacer.Replace(input)
	input = strings.TrimRight(input, ".") // When we have a body full stop we want to remove it.
	transformed := ToCase(input, caseType)

	match, _ := regexp.MatchString(`\d`, transformed)
	if transformed == "" || match {
		return true
	}

	return input == transformed
}
