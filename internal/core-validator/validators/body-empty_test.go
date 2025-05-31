package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	BODY_EMPTY_EMPTY = parser.ParseConventionalCommit("feat:").Body
	BODY_EMPTY_WITH  = parser.ParseConventionalCommit("test: subject\n\nbody.").Body
)

func TestBodyEmptyEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := bodyEmptyValidator(
		BODY_EMPTY_EMPTY,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyEmptyEmptyNeverShouldFail(t *testing.T) {
	_, level := bodyEmptyValidator(
		BODY_EMPTY_EMPTY,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyEmptyFilledAlwaysShouldFail(t *testing.T) {
	_, level := bodyEmptyValidator(
		BODY_EMPTY_WITH,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyEmptyFilledNeverShouldSucceed(t *testing.T) {
	_, level := bodyEmptyValidator(
		BODY_EMPTY_WITH,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
