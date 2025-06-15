package rules

import (
	"fmt"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	SUBJECT_MAX_LENGTH_SHORT_SUBJECT = "a"
	SUBJECT_MAX_LENGTH_LONG_SUBJECT  = "ab"
)

var (
	SUBJECT_MAX_LENGTH_EMPTY = parser.ParseConventionalCommit("test: ").Subject
	SUBJECT_MAX_LENGTH_SHORT = parser.ParseConventionalCommit(fmt.Sprintf("test: %s", SUBJECT_MAX_LENGTH_SHORT_SUBJECT)).Subject
	SUBJECT_MAX_LENGTH_LONG  = parser.ParseConventionalCommit(fmt.Sprintf("test: %s", SUBJECT_MAX_LENGTH_LONG_SUBJECT)).Subject
)

func TestSubjectMaxLengthEmptyShouldSucceed(t *testing.T) {
	_, level := subjectMaxLengthValidator(
		SUBJECT_MAX_LENGTH_EMPTY,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectMaxLengthShortShouldSucceed(t *testing.T) {
	_, level := subjectMaxLengthValidator(
		SUBJECT_MAX_LENGTH_SHORT,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectMaxLengthLongShouldFail(t *testing.T) {
	_, level := subjectMaxLengthValidator(
		SUBJECT_MAX_LENGTH_LONG,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
