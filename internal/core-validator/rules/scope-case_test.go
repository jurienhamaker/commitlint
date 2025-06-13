package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	SCOPE_CASE_EMPTY              = parser.ParseConventionalCommit("test: subject").Scope
	SCOPE_CASE_LOWERCASE          = parser.ParseConventionalCommit("test(scope): subject").Scope
	SCOPE_CASE_KEBABCASE          = parser.ParseConventionalCommit("test(scope-kebab): subject").Scope
	SCOPE_CASE_SENTENCECASE       = parser.ParseConventionalCommit("test(Scope Start): subject").Scope
	SCOPE_CASE_MULTIPLE_LOWERCASE = parser.ParseConventionalCommit("test(lower/case,multiple): subject").Scope
	SCOPE_CASE_MULTIPLE_CAMELCASE = parser.ParseConventionalCommit("test(camelCase\\multipleScope, Options): subject").Scope
)

func TestScopeCaseEmptyAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_EMPTY,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseLowercaseAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseLowercaseNeverLowercaseShouldFail(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_LOWERCASE,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseLowercaseAlwaysUppercaseShouldFail(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseKebabcaseAlwaysKebabcaseShouldSucceed(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_KEBABCASE,
		validation.ValidationStateError,
		true,
		[]string{"kebabcase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseKebabcaseNeverKebabcaseShouldFail(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_KEBABCASE,
		validation.ValidationStateError,
		false,
		[]string{"kebabcase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseSentencecaseAlwaysSentencecaseShouldSucceed(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_SENTENCECASE,
		validation.ValidationStateError,
		true,
		[]string{"sentencecase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseSentencecaseNeverSentencecaseShouldFail(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_SENTENCECASE,
		validation.ValidationStateError,
		false,
		[]string{"sentencecase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseLowercaseMultipleAlwaysLowercaseOrSentencecaseShouldSucceed(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_MULTIPLE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"lowercase", "sentencecase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseLowercaseMultipleNeverLowercaseOrSentencecaseShouldFail(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_MULTIPLE_LOWERCASE,
		validation.ValidationStateError,
		false,
		[]string{"lowercase", "sentencecase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseCamelcaseMultipleAlwaysCamelcaseShouldFail(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_MULTIPLE_CAMELCASE,
		validation.ValidationStateError,
		true,
		[]string{"camelcase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestScopeCaseCamelcaseMultipleAlwaysCamelcaseOrSentencecaseShouldSucceed(t *testing.T) {
	_, level := scopeCaseValidator(
		SCOPE_CASE_MULTIPLE_CAMELCASE,
		validation.ValidationStateError,
		true,
		[]string{"camelcase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
