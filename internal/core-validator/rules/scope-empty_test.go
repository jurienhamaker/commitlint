package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	SCOPE_EMPTY_EMPTY = parser.ParseConventionalCommit("feat:").Scope
	SCOPE_EMPTY_WITH  = parser.ParseConventionalCommit("test(scope): subject").Scope
)

func TestScopeEmptyEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := scopeEmptyValidator(
		SCOPE_EMPTY_EMPTY,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeEmptyEmptyNeverShouldFail(t *testing.T) {
	_, level := scopeEmptyValidator(
		SCOPE_EMPTY_EMPTY,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeEmptyFilledAlwaysShouldFail(t *testing.T) {
	_, level := scopeEmptyValidator(
		SCOPE_EMPTY_WITH,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeEmptyFilledNeverShouldSucceed(t *testing.T) {
	_, level := scopeEmptyValidator(
		SCOPE_EMPTY_WITH,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
