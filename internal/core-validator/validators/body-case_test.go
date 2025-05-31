package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	BODY_CASE_EMPTY = parser.ParseConventionalCommit("test: subject").Body
	BODY_CASE_LOWER = parser.ParseConventionalCommit("test: subject\nbody").Body
	BODY_CASE_MIXED = parser.ParseConventionalCommit("test:  subject\nBody").Body
	BODY_CASE_UPPER = parser.ParseConventionalCommit("test: subject\nBODY").Body
)

func TestBodyCaseEmptyAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := bodyCaseValidator(
		BODY_CASE_EMPTY,
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
		BODY_CASE_EMPTY,
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
		BODY_CASE_EMPTY,
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
		BODY_CASE_EMPTY,
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
		BODY_CASE_LOWER,
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
		BODY_CASE_LOWER,
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
		BODY_CASE_MIXED,
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
		BODY_CASE_MIXED,
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
		BODY_CASE_MIXED,
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
		BODY_CASE_MIXED,
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
		BODY_CASE_UPPER,
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
		BODY_CASE_UPPER,
		validation.ValidationStateError,
		true,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
