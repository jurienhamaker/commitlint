package rules

import (
	"fmt"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	SUBJECT_MIN_LENGTH_SHORT_SUBJECT = "a"
	SUBJECT_MIN_LENGTH_LONG_SUBJECT  = "ab"
)

var (
	SUBJECT_MIN_LENGTH_EMPTY = parser.ParseConventionalCommit("test: ").Subject
	SUBJECT_MIN_LENGTH_SHORT = parser.ParseConventionalCommit(fmt.Sprintf("test: %s", SUBJECT_MIN_LENGTH_SHORT_SUBJECT)).Subject
	SUBJECT_MIN_LENGTH_LONG  = parser.ParseConventionalCommit(fmt.Sprintf("test: %s", SUBJECT_MIN_LENGTH_LONG_SUBJECT)).Subject
)

func TestSubjectMinLengthEmptyShouldSucceed(t *testing.T) {
	_, level := subjectMinLengthValidator(
		SUBJECT_MIN_LENGTH_EMPTY,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectMinLengthShortShouldFail(t *testing.T) {
	_, level := subjectMinLengthValidator(
		SUBJECT_MIN_LENGTH_SHORT,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectMinLengthLongShouldSucceed(t *testing.T) {
	_, level := subjectMinLengthValidator(
		SUBJECT_MIN_LENGTH_LONG,
		validation.ValidationStateError,
		2,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
