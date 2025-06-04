package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	HEADER_FULL_STOP_WITH    = parser.ParseConventionalCommit("test: subject.\n\nheader.").Header
	HEADER_FULL_STOP_WITHOUT = parser.ParseConventionalCommit("test: subject\n\nheader").Header
)

func TestHeaderFullStopWithAlwaysShouldSucceed(t *testing.T) {
	_, level := headerFullStopValidator(
		HEADER_FULL_STOP_WITH,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderFullStopWithNeverShouldFail(t *testing.T) {
	_, level := headerFullStopValidator(
		HEADER_FULL_STOP_WITH,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderFullStopWithoutAlwaysShouldFail(t *testing.T) {
	_, level := headerFullStopValidator(
		HEADER_FULL_STOP_WITHOUT,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestHeaderFullStopWithoutNeverShouldSucceed(t *testing.T) {
	_, level := headerFullStopValidator(
		HEADER_FULL_STOP_WITHOUT,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
