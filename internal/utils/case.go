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

func ToUpperCaseFirst(input string) string {
	first := input[0:1]
	rest := input[1:]
	return strings.ToUpper(first) + rest
}

func ToSentenceCase(input string) string {
	first := input[0:1]
	rest := input[1:]
	return strings.ToUpper(first) + strings.ToLower(rest)
}

func ToPascalCase(input string) string {
	reg := regexp.MustCompile(`\w+\W?`)
	splitted := reg.FindAllString(input, -1)
	newSplitted := []string{}

	for _, v := range splitted {
		newSplitted = append(newSplitted, ToUpperCaseFirst(v))
	}

	return strings.Join(newSplitted, "")
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
		return ToPascalCase(input)
	case KebabCase:
		return strcase.ToKebab(input)
	case SentenceCase:
		return ToSentenceCase(input)
	}

	return input
}

var replacer = strings.NewReplacer("'", "", "\"", "", "`", "")

func EnsureCase(input string, caseType Case) bool {
	input = replacer.Replace(input)
	input = strings.TrimRight(input, ".") // When we have a body full stop we want to remove it.
	transformed := ToCase(input, caseType)

	if transformed == "" {
		return true
	}

	return input == transformed
}
