package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	BODY_CASE_EMPTY = "test: subject"
	BODY_CASE_LOWER = "test: subject\nbody"
	BODY_CASE_MIXED = "test: subject\nBody"
	BODY_CASE_UPPER = "test: subject\nBODY"
)

var (
	BODY_CASE_EMPTY_PARSED = parser.ParseConventionalCommit(BODY_CASE_EMPTY)
	BODY_CASE_LOWER_PARSED = parser.ParseConventionalCommit(BODY_CASE_LOWER)
	BODY_CASE_MIXED_PARSED = parser.ParseConventionalCommit(BODY_CASE_MIXED)
	BODY_CASE_UPPER_PARSED = parser.ParseConventionalCommit(BODY_CASE_UPPER)
)

func TestBodyCaseEmptyAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_EMPTY_PARSED.Body,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseEmptyNeverLowercaseShouldSucceed(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_EMPTY_PARSED.Body,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseEmptyAlwaysUppercaseShouldSucceed(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_EMPTY_PARSED.Body,
		validation.ValidationStateError,
		true,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseEmptyNeverUppercaseShouldSucceed(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_EMPTY_PARSED.Body,
		validation.ValidationStateError,
		false,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseLowerNeverLowercaseShouldFail(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_LOWER_PARSED.Body,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseLowerAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_LOWER_PARSED.Body,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseMixedNeverLowercaseShouldSucceed(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_MIXED_PARSED.Body,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseMixedAlwaysLowercaseShouldFail(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_MIXED_PARSED.Body,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "errror" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseMixedNeverUppercaseShouldSucceed(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_MIXED_PARSED.Body,
		validation.ValidationStateError,
		false,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseMixedAlwaysUppercaseShouldFail(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_MIXED_PARSED.Body,
		validation.ValidationStateError,
		true,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseUpperNeverUppercaseShouldFail(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_UPPER_PARSED.Body,
		validation.ValidationStateError,
		false,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyCaseUpperAlwaysUppercaseShouldSucceed(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_UPPER_PARSED.Body,
		validation.ValidationStateError,
		true,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
