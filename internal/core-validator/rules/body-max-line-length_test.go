package rules

import (
	"fmt"
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

const (
	BODY_MAX_LINE_LENGTH_SHORT_BODY = "a"
	BODY_MAX_LINE_LENGTH_LONG_BODY  = "ab"
)

var (
	BODY_MAX_LINE_LENGTH_EMPTY       = parser.ParseConventionalCommit("test: subject").Body
	BODY_MAX_LINE_LENGTH_SHORT       = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", BODY_MAX_LINE_LENGTH_SHORT_BODY)).Body
	BODY_MAX_LINE_LENGTH_LONG        = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s", BODY_MAX_LINE_LENGTH_LONG_BODY)).Body
	BODY_MAX_LINE_LENGTH_SHORT_MULTI = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s\n%s\n%s", BODY_MAX_LINE_LENGTH_SHORT_BODY, BODY_MAX_LENGTH_SHORT_BODY, BODY_MAX_LENGTH_SHORT_BODY)).Body
	BODY_MAX_LINE_LENGTH_LONG_MULTI  = parser.ParseConventionalCommit(fmt.Sprintf("test: subject\n%s\n%s\n%s", BODY_MAX_LINE_LENGTH_SHORT_BODY, BODY_MAX_LENGTH_LONG_BODY, BODY_MAX_LENGTH_SHORT_BODY)).Body
)

func TestBodyMaxLineLengthEmptyShouldSucceed(t *testing.T) {
	_, level := bodyMaxLineLengthValidator(
		BODY_MAX_LINE_LENGTH_EMPTY,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyMaxLineLengthShortShouldSucceed(t *testing.T) {
	_, level := bodyMaxLineLengthValidator(
		BODY_MAX_LINE_LENGTH_SHORT,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyMaxLineLengthLongShouldFail(t *testing.T) {
	_, level := bodyMaxLineLengthValidator(
		BODY_MAX_LINE_LENGTH_LONG,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyMaxLineLengthShortMultiShouldSucceed(t *testing.T) {
	_, level := bodyMaxLineLengthValidator(
		BODY_MAX_LINE_LENGTH_SHORT_MULTI,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestBodyMaxLineLengthLongMultiShouldFail(t *testing.T) {
	_, level := bodyMaxLineLengthValidator(
		BODY_MAX_LINE_LENGTH_LONG_MULTI,
		validation.ValidationStateError,
		1,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
