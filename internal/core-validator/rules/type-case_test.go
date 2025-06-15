package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	TYPE_CASE_EMPTY        = parser.ParseConventionalCommit("type").Type
	TYPE_CASE_LOWERCASE    = parser.ParseConventionalCommit("test: lowercase").Type
	TYPE_CASE_SENTENCECASE = parser.ParseConventionalCommit("Test: Type is in sentence case").Type
	TYPE_CASE_KEBABCASE    = parser.ParseConventionalCommit("test-type: Type is in sentence case").Type
)

func TestTypeCaseNumericAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_EMPTY,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseNumericNeverLowercaseShouldSucceed(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_EMPTY,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseNumericAlwaysUppercaseShouldSucceed(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_EMPTY,
		validation.ValidationStateError,
		true,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseNumericNeverUppercaseShouldSucceed(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_EMPTY,
		validation.ValidationStateError,
		false,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseLowercaseAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseLowercaseAlwaysLowercaseOrUppercaseShouldSucceed(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"lowercase", "uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseLowercaseNeverLowercaseShouldFail(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_LOWERCASE,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseLowercaseAlwaysUppercaseOrSentencecaseShouldFail(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"uppercase", "sentencecase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseSentenceCaseAlwaysSentencecaseShouldSucceed(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_SENTENCECASE,
		validation.ValidationStateError,
		true,
		[]string{"sentencecase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseSentenceCaseAlwaysKebabcaseShouldFail(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_SENTENCECASE,
		validation.ValidationStateError,
		true,
		[]string{"kebabcase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseSentenceCaseNeverSentenceCaseShouldFail(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_SENTENCECASE,
		validation.ValidationStateError,
		false,
		[]string{"sentencecase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseKebabCaseAlwaysKebabcaseShouldSucceed(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_KEBABCASE,
		validation.ValidationStateError,
		true,
		[]string{"kebabcase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeCaseKebabCaseAlwaysUppercaseOrSentencecaseShouldFail(t *testing.T) {
	_, level := typeCaseValidator(
		TYPE_CASE_KEBABCASE,
		validation.ValidationStateError,
		true,
		[]string{"uppercase", "sentencecase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
