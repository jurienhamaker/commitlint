package rules

import (
	"fmt"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	SCOPE_MIN_LENGTH_SHORT_SCOPE = "a"
	SCOPE_MIN_LENGTH_LONG_SCOPE  = "ab"
)

var (
	SCOPE_MIN_LENGTH_EMPTY = parser.ParseConventionalCommit("feat: no scope").Scope
	SCOPE_MIN_LENGTH_SHORT = parser.ParseConventionalCommit(fmt.Sprintf("feat(%s): short scope", SCOPE_MIN_LENGTH_SHORT_SCOPE)).Scope
	SCOPE_MIN_LENGTH_LONG  = parser.ParseConventionalCommit(fmt.Sprintf("feat(%s): long scope", SCOPE_MIN_LENGTH_LONG_SCOPE)).Scope
)

func TestScopeMinLengthEmptyShouldSucceed(t *testing.T) {
	_, level := scopeMinLengthValidator(
		SCOPE_MIN_LENGTH_EMPTY,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeMinLengthShortShouldFail(t *testing.T) {
	_, level := scopeMinLengthValidator(
		SCOPE_MIN_LENGTH_SHORT,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeMinLengthLongShouldSucceed(t *testing.T) {
	_, level := scopeMinLengthValidator(
		SCOPE_MIN_LENGTH_LONG,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
