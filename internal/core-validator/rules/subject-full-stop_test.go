package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	SUBJECT_FULL_STOP_EMPTY   = parser.ParseConventionalCommit("test:").Subject
	SUBJECT_FULL_STOP_WITH    = parser.ParseConventionalCommit("test: subject.").Subject
	SUBJECT_FULL_STOP_WITHOUT = parser.ParseConventionalCommit("test: subject").Subject
)

func TestSubjectFullStopEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := subjectFullStopValidator(
		SUBJECT_FULL_STOP_EMPTY,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectFullSubjectStopEmptyNeverShouldSucceed(t *testing.T) {
	_, level := subjectFullStopValidator(
		SUBJECT_FULL_STOP_EMPTY,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectFullStopWithAlwaysShouldSucceed(t *testing.T) {
	_, level := subjectFullStopValidator(
		SUBJECT_FULL_STOP_WITH,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectFullStopWithNeverShouldFail(t *testing.T) {
	_, level := subjectFullStopValidator(
		SUBJECT_FULL_STOP_WITH,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectFullStopWithoutAlwaysShouldFail(t *testing.T) {
	_, level := subjectFullStopValidator(
		SUBJECT_FULL_STOP_WITHOUT,
		validation.ValidationStateError,
		true,
		".",
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectFullStopWithoutNeverShouldSucceed(t *testing.T) {
	_, level := subjectFullStopValidator(
		SUBJECT_FULL_STOP_WITHOUT,
		validation.ValidationStateError,
		false,
		".",
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
