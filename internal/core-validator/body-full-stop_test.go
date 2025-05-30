package corevalidator

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	BODY_EMPTY   = "feat:"
	BODY_WITH    = "test: subject\n\nbody."
	BODY_WITHOUT = "test: subject\n\nbody"
)

var (
	BODY_EMPTY_PARSED   = parser.ParseConventionalCommit(BODY_EMPTY)
	BODY_WITH_PARSED    = parser.ParseConventionalCommit(BODY_WITH)
	BODY_WITHOUT_PARSED = parser.ParseConventionalCommit(BODY_WITHOUT)
)

func TestEmptyBodyAlwaysShouldSucceed(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_EMPTY_PARSED.Body,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestEmptyBodyNeverShouldSucceed(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_EMPTY_PARSED.Body,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyWithAlwaysShouldSucceed(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_WITH_PARSED.Body,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyWithNeverShouldFail(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_WITH_PARSED.Body,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyWithoutAlwaysShouldFail(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_WITHOUT_PARSED.Body,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyWithoutNeverShouldSucceed(t *testing.T) {
	_, level := bodyFullStopValidator(
		BODY_WITHOUT_PARSED.Body,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
