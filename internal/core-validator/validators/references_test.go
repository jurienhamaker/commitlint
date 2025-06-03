package validators

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	REFERENCES_EMPTY       = parser.ParseConventionalCommit("").Raw
	REFERENCES_SINGLE      = parser.ParseConventionalCommit("test: subject\nFixes #123").Raw
	REFERENCES_MULTI       = parser.ParseConventionalCommit("test: subject\nFixes #123 and #456").Raw
	REFERENCES_NONE        = parser.ParseConventionalCommit("test: subject\nNo references").Raw
	REFERENCES_SINGLE_PRO  = parser.ParseConventionalCommit("test: subject\nPRO-1234").Raw
	REFERENCES_MULTI_MIXED = parser.ParseConventionalCommit("test: solves #1234\n\nAnd PRO-1234").Raw
	REFERENCES_SHA         = parser.ParseConventionalCommit("test: subject\nReverts baff22bb4").Raw
)

func TestReferencesEmptyAlwaysShouldSucceed(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_EMPTY,
		validation.ValidationStateError,
		true,
		[]string{"#"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesEmptyNeverShouldSucceed(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_EMPTY,
		validation.ValidationStateError,
		false,
		[]string{"#"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesSingleAlwaysShouldSucceed(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_SINGLE,
		validation.ValidationStateError,
		true,
		[]string{"#"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesSingleNeverShouldFail(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_SINGLE,
		validation.ValidationStateError,
		false,
		[]string{"#"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesMultiAlwaysShouldSucceed(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_MULTI,
		validation.ValidationStateError,
		true,
		[]string{"#"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesMultiNeverShouldFail(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_MULTI,
		validation.ValidationStateError,
		false,
		[]string{"#"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesNoneAlwaysShouldFail(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_NONE,
		validation.ValidationStateError,
		true,
		[]string{"#"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesNoneNeverShouldSucceed(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_NONE,
		validation.ValidationStateError,
		false,
		[]string{"#"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesSingleProAlwaysShouldSucceed(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_SINGLE_PRO,
		validation.ValidationStateError,
		true,
		[]string{"PRO-"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesSingleProNeverShouldFail(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_SINGLE_PRO,
		validation.ValidationStateError,
		false,
		[]string{"PRO-"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesMultiMixedAlwaysShouldSucceed(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_MULTI_MIXED,
		validation.ValidationStateError,
		true,
		[]string{"#", "PRO-"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesMultiMixedNeverShouldFail(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_MULTI_MIXED,
		validation.ValidationStateError,
		false,
		[]string{"#", "PRO-"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesShaAlwaysShouldSucceed(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_SHA,
		validation.ValidationStateError,
		true,
		[]string{"sha"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestReferencesShaAlwaysShouldFail(t *testing.T) {
	_, level := referencesValidator(
		REFERENCES_SHA,
		validation.ValidationStateError,
		false,
		[]string{"sha"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}
