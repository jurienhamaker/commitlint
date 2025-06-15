package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	SUBJECT_EMPTY_EMPTY = parser.ParseConventionalCommit("test:").Subject
	SUBJECT_EMPTY_WITH  = parser.ParseConventionalCommit("test: subject").Subject
)

func TestSubjectEmptyEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := subjectEmptyValidator(
		SUBJECT_EMPTY_EMPTY,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectEmptyEmptyNeverShouldFail(t *testing.T) {
	_, level := subjectEmptyValidator(
		SUBJECT_EMPTY_EMPTY,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectEmptyFilledAlwaysShouldFail(t *testing.T) {
	_, level := subjectEmptyValidator(
		SUBJECT_EMPTY_WITH,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectEmptyFilledNeverShouldSucceed(t *testing.T) {
	_, level := subjectEmptyValidator(
		SUBJECT_EMPTY_WITH,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
