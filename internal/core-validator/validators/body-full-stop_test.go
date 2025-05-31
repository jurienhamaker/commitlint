package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	BODY_FULL_STOP_EMPTY   = "feat:"
	BODY_FULL_STOP_WITH    = "test: subject\n\nbody."
	BODY_FULL_STOP_WITHOUT = "test: subject\n\nbody"
)

var (
	BODY_FULL_STOP_EMPTY_PARSED   = parser.ParseConventionalCommit(BODY_FULL_STOP_EMPTY)
	BODY_FULL_STOP_WITH_PARSED    = parser.ParseConventionalCommit(BODY_FULL_STOP_WITH)
	BODY_FULL_STOP_WITHOUT_PARSED = parser.ParseConventionalCommit(BODY_FULL_STOP_WITHOUT)
)

func TestBodyFullStopEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_FULL_STOP_EMPTY_PARSED.Body,
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
		BODY_FULL_STOP_EMPTY_PARSED.Body,
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
		BODY_FULL_STOP_WITH_PARSED.Body,
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
		BODY_FULL_STOP_WITH_PARSED.Body,
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
		BODY_FULL_STOP_WITHOUT_PARSED.Body,
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
		BODY_FULL_STOP_WITHOUT_PARSED.Body,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
