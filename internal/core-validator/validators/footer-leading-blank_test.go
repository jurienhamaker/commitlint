package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	FOOTER_LEADING_BLANK_SIMPLE       = parser.ParseConventionalCommit("test: subject")
	FOOTER_LEADING_BLANK_BODY         = parser.ParseConventionalCommit("test: subject\nBody")
	FOOTER_LEADING_BLANK_TRAILING     = parser.ParseConventionalCommit("test: subject\nBody\n\n")
	FOOTER_LEADING_BLANK_WITHOUT      = parser.ParseConventionalCommit("test: subject\nBody\nBREAKING CHANGE: something important")
	FOOTER_LEADING_BLANK_WITHOUT_BODY = parser.ParseConventionalCommit("test: subject\n\nBREAKING CHANGE: something imortant")
	FOOTER_LEADING_BLANK_WITH         = parser.ParseConventionalCommit("test: subject\nbody\n\nBREAKING CHANGE: something important")
	FOOTER_LEADING_BLANK_WITH_MULTI   = parser.ParseConventionalCommit("test: subject\nmulti\nline\nbody\n\nBREAKING CHANGE: something important")
)

func TestFooterLeadingBlankSimpleAlwaysShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_SIMPLE.Body,
		FOOTER_LEADING_BLANK_SIMPLE.Footer,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankSimpleNeverShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_SIMPLE.Body,
		FOOTER_LEADING_BLANK_SIMPLE.Footer,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankBodyAlwaysShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_BODY.Body,
		FOOTER_LEADING_BLANK_BODY.Footer,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankBodyNeverShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_BODY.Body,
		FOOTER_LEADING_BLANK_BODY.Footer,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankTrailingAlwaysShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_TRAILING.Body,
		FOOTER_LEADING_BLANK_TRAILING.Footer,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankTrailingNeverShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_TRAILING.Body,
		FOOTER_LEADING_BLANK_TRAILING.Footer,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankWithoutAlwaysShouldFail(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_WITHOUT.Body,
		FOOTER_LEADING_BLANK_WITHOUT.Footer,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankWithoutNeverShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_WITHOUT.Body,
		FOOTER_LEADING_BLANK_WITHOUT.Footer,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankWithoutBodyAlwaysShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_WITHOUT_BODY.Body,
		FOOTER_LEADING_BLANK_WITHOUT_BODY.Footer,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankWithoutBodyNeverShouldFail(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_WITHOUT_BODY.Body,
		FOOTER_LEADING_BLANK_WITHOUT_BODY.Footer,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankWithAlwaysShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_WITH.Body,
		FOOTER_LEADING_BLANK_WITH.Footer,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankWithNeverShouldFail(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_WITH.Body,
		FOOTER_LEADING_BLANK_WITH.Footer,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankWithMultiAlwaysShouldSucceed(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_WITH_MULTI.Body,
		FOOTER_LEADING_BLANK_WITH_MULTI.Footer,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestFooterLeadingBlankWithMultiNeverShouldFail(t *testing.T) {
	_, level := footerLeadingBlankValidator(
		FOOTER_LEADING_BLANK_WITH_MULTI.Body,
		FOOTER_LEADING_BLANK_WITH_MULTI.Footer,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
