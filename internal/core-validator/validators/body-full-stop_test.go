package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	BODY_FULL_STOP_EMPTY   = parser.ParseConventionalCommit("feat:").Body
	BODY_FULL_STOP_WITH    = parser.ParseConventionalCommit("test: subject\n\nbody.").Body
	BODY_FULL_STOP_WITHOUT = parser.ParseConventionalCommit("test: subject\n\nbody").Body
)

func TestBodyFullStopEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_FULL_STOP_EMPTY,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyFullBodyStopEmptyNeverShouldSucceed(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_FULL_STOP_EMPTY,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyFullStopWithAlwaysShouldSucceed(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_FULL_STOP_WITH,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyFullStopWithNeverShouldFail(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_FULL_STOP_WITH,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyFullStopWithoutAlwaysShouldFail(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_FULL_STOP_WITHOUT,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyFullStopWithoutNeverShouldSucceed(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_FULL_STOP_WITHOUT,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
