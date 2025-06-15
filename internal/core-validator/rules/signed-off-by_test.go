package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	SIGNED_OFF_BY_EMPTY                         = parser.ParseConventionalCommit("test: subject\n").Footer
	SIGNED_OFF_BY_WITH                          = parser.ParseConventionalCommit("test: subject\nBody\nSigned-off-by: signer").Footer
	SIGNED_OFF_BY_WITHOUT                       = parser.ParseConventionalCommit("test: subject\nBody\nfooter").Footer
	SIGNED_OFF_BY_IN_SUBJECT                    = parser.ParseConventionalCommit("test: Signed-off-by: signer\nBody\nfooter").Footer
	SIGNED_OFF_BY_WITH_COMMENTS_AND_EMPTY_LINES = parser.ParseConventionalCommit("test: subject\nbody\nSigned-off-by: signer\n \n#comment\n#More comments\nNot-signed-off-by: apple").Footer
)

func TestSignedOffByEmptyAlwaysShouldFail(t *testing.T) {
	_, level := signedOffByValidator(
		SIGNED_OFF_BY_EMPTY,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSignedOffByEmptyNeverShouldSucceed(t *testing.T) {
	_, level := signedOffByValidator(
		SIGNED_OFF_BY_EMPTY,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSignedOffByWithAlwaysShouldSucceed(t *testing.T) {
	_, level := signedOffByValidator(
		SIGNED_OFF_BY_WITH,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSignedOffByWithNeverShouldFail(t *testing.T) {
	_, level := signedOffByValidator(
		SIGNED_OFF_BY_WITH,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSignedOffByWithoutAlwaysShouldFail(t *testing.T) {
	_, level := signedOffByValidator(
		SIGNED_OFF_BY_WITHOUT,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSignedOffByWithoutNeverShouldSucceed(t *testing.T) {
	_, level := signedOffByValidator(
		SIGNED_OFF_BY_WITHOUT,
		validation.ValidationStateError,
		false,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSignedOffByInSubjectAlwaysShouldFail(t *testing.T) {
	_, level := signedOffByValidator(
		SIGNED_OFF_BY_IN_SUBJECT,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSignedOffByWithCommentsAndEmptyLinesAlwaysShouldSucceed(t *testing.T) {
	_, level := signedOffByValidator(
		SIGNED_OFF_BY_WITH_COMMENTS_AND_EMPTY_LINES,
		validation.ValidationStateError,
		true,
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
