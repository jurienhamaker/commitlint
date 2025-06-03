package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	HEADER_CASE_NUMERIC   = parser.ParseConventionalCommit("1.0.0").Header
	HEADER_CASE_LOWERCASE = parser.ParseConventionalCommit("header").Header
)

func TestHeaderCaseNumericAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := headerCaseValidator(
		HEADER_CASE_NUMERIC,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderCaseNumericNeverLowercaseShouldSucceed(t *testing.T) {
	_, level := headerCaseValidator(
		HEADER_CASE_NUMERIC,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderCaseNumericAlwaysUppercaseShouldSucceed(t *testing.T) {
	_, level := headerCaseValidator(
		HEADER_CASE_NUMERIC,
		validation.ValidationStateError,
		true,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderCaseNumericNeverUppercaseShouldSucceed(t *testing.T) {
	_, level := headerCaseValidator(
		HEADER_CASE_NUMERIC,
		validation.ValidationStateError,
		false,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderCaseLowercaseAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := headerCaseValidator(
		HEADER_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderCaseLowercaseAlwaysLowercaseOrUppercaseShouldSucceed(t *testing.T) {
	_, level := headerCaseValidator(
		HEADER_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"lowercase", "uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderCaseLowercaseNeverLowercaseShouldFail(t *testing.T) {
	_, level := headerCaseValidator(
		HEADER_CASE_LOWERCASE,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderCaseLowercaseAlwaysUppercaseOrSentencecaseShouldFail(t *testing.T) {
	_, level := headerCaseValidator(
		HEADER_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"uppercase", "sentencecase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
