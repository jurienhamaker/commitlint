package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	TYPE_EMPTY_EMPTY = parser.ParseConventionalCommit("subject").Type
	TYPE_EMPTY_WITH  = parser.ParseConventionalCommit("test: subject").Type
)

func TestTypeEmptyEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := typeEmptyValidator(
		TYPE_EMPTY_EMPTY,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeEmptyEmptyNeverShouldFail(t *testing.T) {
	_, level := typeEmptyValidator(
		TYPE_EMPTY_EMPTY,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeEmptyFilledAlwaysShouldFail(t *testing.T) {
	_, level := typeEmptyValidator(
		TYPE_EMPTY_WITH,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeEmptyFilledNeverShouldSucceed(t *testing.T) {
	_, level := typeEmptyValidator(
		TYPE_EMPTY_WITH,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
