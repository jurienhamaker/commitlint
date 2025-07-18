package rules

import (
	"testing"

	"github.com/jurienhamaker/commitlint/parser"
	"github.com/jurienhamaker/commitlint/validation"
)

var (
	SUBJECT_CASE_NUMERIC            = parser.ParseConventionalCommit("feat: 1.0.0").Subject
	SUBJECT_CASE_LOWERCASE          = parser.ParseConventionalCommit("feat: subject").Subject
	SUBJECT_CASE_SENTENCECASE       = parser.ParseConventionalCommit("feat: Subject is in sentence case").Subject
	SUBJECT_CASE_PASCALCASE         = parser.ParseConventionalCommit("feat: Subject Is In Pascal Case").Subject
	SUBJECT_CASE_NUMBERS_AND_DASHES = parser.ParseConventionalCommit("feat(scope): Subject contains a 0 and some-dashes").Subject
)

func TestSubjectCaseNumericAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_NUMERIC,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseNumericNeverLowercaseShouldSucceed(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_NUMERIC,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseNumericAlwaysUppercaseShouldSucceed(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_NUMERIC,
		validation.ValidationStateError,
		true,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseNumericNeverUppercaseShouldSucceed(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_NUMERIC,
		validation.ValidationStateError,
		false,
		[]string{"uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseLowercaseAlwaysLowercaseShouldSucceed(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseLowercaseAlwaysLowercaseOrUppercaseShouldSucceed(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"lowercase", "uppercase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseLowercaseNeverLowercaseShouldFail(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_LOWERCASE,
		validation.ValidationStateError,
		false,
		[]string{"lowercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseLowercaseAlwaysUppercaseOrSentencecaseShouldFail(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_LOWERCASE,
		validation.ValidationStateError,
		true,
		[]string{"uppercase", "sentencecase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseSentencecaseAlwaysSentencecaseShouldSucceed(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_SENTENCECASE,
		validation.ValidationStateError,
		true,
		[]string{"sentencecase"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCasePascalcaseAlwaysSentencecaseOrLowercaseShouldFail(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_PASCALCASE,
		validation.ValidationStateError,
		true,
		[]string{"sentencecase", "lowercase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseSentencecaseAlwaysKebabcaseShouldFail(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_SENTENCECASE,
		validation.ValidationStateError,
		true,
		[]string{"kebabcase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseSentencecaseNeverSentencecaseShouldFail(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_SENTENCECASE,
		validation.ValidationStateError,
		false,
		[]string{"sentencecase"},
	)

	if level != validation.ValidationStateError {
		t.Errorf(`Expected level to equal "error" got "%s"`, validation.ValidationStateName[level])
	}
}

func TestSubjectCaseNumbersAndDashesNeverPascalCaseOrUpperCaseShouldSucceed(t *testing.T) {
	_, level := subjectCaseValidator(
		SUBJECT_CASE_NUMBERS_AND_DASHES,
		validation.ValidationStateError,
		false,
		[]string{"pascal-case", "upper-case"},
	)

	if level != validation.ValidationStateSuccess {
		t.Errorf(`Expected level to equal "success" got "%s"`, validation.ValidationStateName[level])
	}
}
