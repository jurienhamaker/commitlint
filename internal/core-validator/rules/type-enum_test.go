package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	TYPE_ENUM_EMPTY = parser.ParseConventionalCommit("subject").Type
	TYPE_ENUM_TEST  = parser.ParseConventionalCommit("test: subject").Type
	TYPE_ENUM_FEAT  = parser.ParseConventionalCommit("feat: subject").Type
)

func TestTypeEnumEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := typeEnumValidator(
		TYPE_ENUM_EMPTY,
		validation.ValidationStateError,
		true,
		[]string{"test"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeEnumTestAlwaysTestShouldSucceed(t *testing.T) {
	_, level := typeEnumValidator(
		TYPE_ENUM_TEST,
		validation.ValidationStateError,
		true,
		[]string{"test"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeEnumTestNeverTestShouldFail(t *testing.T) {
	_, level := typeEnumValidator(
		TYPE_ENUM_TEST,
		validation.ValidationStateError,
		false,
		[]string{"test"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeEnumTestNeverFeatOrChoreShouldSucceed(t *testing.T) {
	_, level := typeEnumValidator(
		TYPE_ENUM_TEST,
		validation.ValidationStateError,
		false,
		[]string{"feat", "chore"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeEnumFeatAlwaysTestShouldFail(t *testing.T) {
	_, level := typeEnumValidator(
		TYPE_ENUM_FEAT,
		validation.ValidationStateError,
		true,
		[]string{"test"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestTypeEnumFeatAlwaysFeatOrChoreShouldSucceed(t *testing.T) {
	_, level := typeEnumValidator(
		TYPE_ENUM_FEAT,
		validation.ValidationStateError,
		true,
		[]string{"feat", "chore"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
