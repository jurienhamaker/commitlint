package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	HEADER_TRIM_EMPTY = parser.ParseConventionalCommit("").Header
	HEADER_TRIM_VALID = parser.ParseConventionalCommit("test: header").Header

	HEADER_TRIM_END   = parser.ParseConventionalCommit("test: header ").Header
	HEADER_TRIM_START = parser.ParseConventionalCommit(" test: header").Header
	HEADER_TRIM_BOTH  = parser.ParseConventionalCommit(" test: header ").Header

	HEADER_TRIM_END_TAB   = parser.ParseConventionalCommit("test: header\t").Header
	HEADER_TRIM_START_TAB = parser.ParseConventionalCommit("\ttest: header").Header
	HEADER_TRIM_BOTH_TAB  = parser.ParseConventionalCommit("\ttest: header\t").Header

	HEADER_TRIM_END_MIXED   = parser.ParseConventionalCommit("test: header \t").Header
	HEADER_TRIM_START_MIXED = parser.ParseConventionalCommit("\t test: header").Header
	HEADER_TRIM_BOTH_MIXED  = parser.ParseConventionalCommit("\t test: header\t ").Header
)

func TestHeaderTrimEmptyShouldSucceed(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_EMPTY,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimValidShouldSucceed(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_VALID,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimEndShouldFail(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_END,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimStartShouldFail(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_START,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimBothShouldFail(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_BOTH,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimEndTabShouldFail(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_END_TAB,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimStartTabShouldFail(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_START_TAB,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimBothTabShouldFail(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_BOTH_TAB,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimEndMixedShouldFail(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_END_MIXED,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimStartMixedShouldFail(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_START_MIXED,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderTrimBothMixedShouldFail(t *testing.T) {
	_, level := headerTrimValidator(
		HEADER_TRIM_BOTH_MIXED,
		validation.ValidationStateError,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
