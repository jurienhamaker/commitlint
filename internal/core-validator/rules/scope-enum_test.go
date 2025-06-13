package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	SCOPE_ENUM_EMPTY  = parser.ParseConventionalCommit("feat:").Scope
	SCOPE_ENUM_SINGLE = parser.ParseConventionalCommit("test(scope): subject").Scope
	SCOPE_ENUM_MULTI  = parser.ParseConventionalCommit("test(scope/multiple, yes): subject").Scope
)

func TestScopeEnumEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := scopeEnumValidator(
		SCOPE_ENUM_EMPTY,
		validation.ValidationStateError,
		true,
		[]string{"scope"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeEnumSingleAlwaysShouldSucceed(t *testing.T) {
	_, level := scopeEnumValidator(
		SCOPE_ENUM_SINGLE,
		validation.ValidationStateError,
		true,
		[]string{"scope"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeEnumSingleNeverShouldFail(t *testing.T) {
	_, level := scopeEnumValidator(
		SCOPE_ENUM_SINGLE,
		validation.ValidationStateError,
		false,
		[]string{"scope"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeEnumSingleAlwaysShouldFail(t *testing.T) {
	_, level := scopeEnumValidator(
		SCOPE_ENUM_SINGLE,
		validation.ValidationStateError,
		true,
		[]string{"neither", "will", "exist"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeEnumMultiAlwaysShouldSucceed(t *testing.T) {
	_, level := scopeEnumValidator(
		SCOPE_ENUM_MULTI,
		validation.ValidationStateError,
		true,
		[]string{"scope", "multiple", "yes"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
