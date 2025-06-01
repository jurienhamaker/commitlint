package parser

import (
	"sort"
	"testing"
)

func TestParseConventionalCommit(t *testing.T) {
	const commit = `feat: add new feature`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = "feat"
	if parsedCommit.Category != expected {
		t.Errorf(`Expected parsedCommit.Category to equal %v got %v`, expected, parsedCommit.Category)
	}
}

func TestParseConventionalCommitBreakingChange(t *testing.T) {
	const commit = `feat: add new feature

BREAKING CHANGE: A new breaking change`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = "feat"
	if parsedCommit.Category != expected && parsedCommit.Major {
		t.Errorf(`Expected parsedCommit.Category to equal %v got %v`, expected, parsedCommit.Category)
	}
}

func TestParseConventionalCommitBreakingChangeShortcut(t *testing.T) {
	const commit = `feat!: add new feature`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = "feat"
	if parsedCommit.Category != expected && parsedCommit.Major {
		t.Errorf(`Expected parsedCommit.Category to equal %v got %v`, expected, parsedCommit.Category)
	}
}

func TestParseConventionalCommitWithScope(t *testing.T) {
	const commit = `feat(my scope): add new feature`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = "my scope"
	if parsedCommit.Scope != expected {
		t.Errorf(`Expected parsedCommit.Scope to equal %v got %v`, expected, parsedCommit.Scope)
	}
}

func TestParseConventionalCommitEmptyBodyAndFooter(t *testing.T) {
	const commit = `feat: add new feature`
	parsedCommit := ParseConventionalCommit(commit)
	if parsedCommit.Body != "" {
		t.Errorf(`Expected parsedCommit.Body to be empty got %v`, parsedCommit.Body)
	}
	if len(parsedCommit.Footer) > 0 {
		t.Errorf(`Expected parsedCommit.Footer to be empty got %+#v`, parsedCommit.Footer)
	}
}

func TestParseConventionalCommitEmptyFooter(t *testing.T) {
	const commit = `feat: add new feature
Description of the new feature`
	parsedCommit := ParseConventionalCommit(commit)
	if len(parsedCommit.Footer) > 0 {
		t.Errorf(`Expected parsedCommit.Footer to be empty got %+#v`, parsedCommit.Footer)
	}
}

func TestParseConventionalCommitBreakingChangeWithScope(t *testing.T) {
	const commit = `feat(my scope): add new feature

BREAKING CHANGE: A new breaking change`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = "my scope"
	if parsedCommit.Scope != expected {
		t.Errorf(`Expected parsedCommit.Scope to equal %v got %v`, expected, parsedCommit.Scope)
	}
}

func TestParseConventionalCommitBreakingChangeShortcutWithScope(t *testing.T) {
	const commit = `feat(my scope)!: add new feature`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = "my scope"
	if parsedCommit.Scope != expected {
		t.Errorf(`Expected parsedCommit.Scope to equal %v got %v`, expected, parsedCommit.Scope)
	}
}

func TestParseConventionalCommitWithBody(t *testing.T) {
	const commit = `feat: add new feature
Description of the new feature`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = "Description of the new feature"
	if parsedCommit.Body != expected {
		t.Errorf(`Expected parsedCommit.Body to equal %v got %v`, expected, parsedCommit.Body)
	}
}

func TestParseConventionalCommitWithMultiLineBody(t *testing.T) {
	const commit = `feat: add new feature
Description of the new feature
more details
even more details`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = `Description of the new feature
more details
even more details`
	if parsedCommit.Body != expected {
		t.Errorf("Expected parsedCommit.Body to equal \n%v \n\ngot:\n\n%v", expected, parsedCommit.Body)
	}
}

func TestParseConventionalCommitBodyExcludesFooter(t *testing.T) {
	const commit = `feat: add new feature
Description of the new feature
more details

Closes #1`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = `Description of the new feature
more details
`
	if parsedCommit.Body != expected {
		t.Errorf("Expected parsedCommit.Body to equal \n%v \n\ngot:\n\n%v", expected, parsedCommit.Body)
	}
}

func TestParseConventionalCommitWithOutBody(t *testing.T) {
	const commit = `feat: add new feature`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = ""
	if parsedCommit.Body != expected {
		t.Errorf(`Expected parsedCommit.Body to equal %v got %v`, expected, parsedCommit.Body)
	}
}

func TestParseConventionalCommitFooter(t *testing.T) {
	const commit = `feat: add new feature
Description of the new feature

Reviewed-by: Z`

	parsedCommit := ParseConventionalCommit(commit)
	const expected = `Reviewed-by: Z`

	if len(parsedCommit.Footer) == 0 || parsedCommit.Footer[0] != expected {
		t.Errorf(`Expected parsedCommit.Footer to equal %v got %v`, expected, parsedCommit.Footer)
	}
}

func TestParseConventionalCommitFooterWithNoBody(t *testing.T) {
	const commit = `feat: add new feature

BREAKING CHANGE: A new breaking change`

	parsedCommit := ParseConventionalCommit(commit)
	const expected = `BREAKING CHANGE: A new breaking change`

	if len(parsedCommit.Footer) == 0 || parsedCommit.Footer[0] != expected {
		t.Errorf(`Expected parsedCommit.Footer to equal %v got %v`, expected, parsedCommit.Footer)
	}
}

func TestParseConventionalCommitMultipleFooter(t *testing.T) {
	const commit = `feat: add new feature

Description of the new feature
more details
even more details

BREAKING CHANGE: A new breaking change
Reviewed-by: Z
Closes #42
`

	parsedCommit := ParseConventionalCommit(commit)
	expected := []string{`BREAKING CHANGE: A new breaking change`, `Reviewed-by: Z`, `Closes #42`}

	if len(parsedCommit.Footer) != len(expected) {
		t.Fatalf(`Expected parsedCommit.Footer to equal %v got %v`, expected, parsedCommit.Footer)
	}
	for i, v := range expected {
		if parsedCommit.Footer[i] != v {
			t.Errorf(`Expected parsedCommit.Footer[%d] to equal %s got %s`, i, v, parsedCommit.Footer[i])
		}
	}
}

func TestParseConventionalCommitMultipleFooterAndLineBreaks(t *testing.T) {
	const commit = `feat: add new feature

Description of the new feature
more details
even more details

BREAKING CHANGE: A new breaking change

Reviewed-by: Z



Closes #42
`

	parsedCommit := ParseConventionalCommit(commit)
	expected := []string{`BREAKING CHANGE: A new breaking change`, `Reviewed-by: Z`, `Closes #42`}

	if len(parsedCommit.Footer) != len(expected) {
		t.Fatalf(`Expected parsedCommit.Footer to equal %v got %v`, expected, parsedCommit.Footer)
	}
	for i, v := range expected {
		if parsedCommit.Footer[i] != v {
			t.Errorf(`Expected parsedCommit.Footer[%d] to equal %s got %s`, i, v, parsedCommit.Footer[i])
		}
	}
}

func TestParseConventionalCommitWithMultiLineBodyWithBreaks(t *testing.T) {
	const commit = `feat: add new feature

Description of the new feature
more details
even more details`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = `
Description of the new feature
more details
even more details`
	if parsedCommit.Body != expected {
		t.Errorf("Expected parsedCommit.Body to equal \n%v \n\ngot:\n\n%v", expected, parsedCommit.Body)
	}
}

func TestParseConventionalCommitFooterWithMultiLineBodyWithBreaks(t *testing.T) {
	const commit = `feat: add new feature

Description of the new feature
more details
even more details


Closes: #42`
	parsedCommit := ParseConventionalCommit(commit)
	const expected = `Closes: #42`
	if len(parsedCommit.Footer) == 0 || parsedCommit.Footer[0] != expected {
		t.Errorf("Expected parsedCommit.Footer to equal \n`%v`\n\ngot:\n`%v`", expected, parsedCommit.Footer[0])
	}
}

// https://github.com/conventional-commits/conventionalcommits.org/issues/313
func TestParseConventionalCommitIssue313MultiLine(t *testing.T) {
	const commit = `feat: add new feature

Description of the new feature
more details
even more details

BREAKING CHANGE: You no longer have to use alertTitle and children to display your alert content. 
Use children only instead. This is because we found out that people actually display the "content" of the alert in alertTitle and not in both alertTitle and children (description). 
Because of that, we're consolidating both. This also opens up opportunities to have other components used within an alert component like an accordion for Progressive Disclosure design patterns.

Closes #42
`
	parsedCommit := ParseConventionalCommit(commit)
	expected := []string{
		`BREAKING CHANGE: You no longer have to use alertTitle and children to display your alert content. 
Use children only instead. This is because we found out that people actually display the "content" of the alert in alertTitle and not in both alertTitle and children (description). 
Because of that, we're consolidating both. This also opens up opportunities to have other components used within an alert component like an accordion for Progressive Disclosure design patterns.`,
		`Closes #42`,
	}

	if len(parsedCommit.Footer) != len(expected) {
		t.Errorf("Expected parsedCommit.Footer to equal \n`%v`\n\ngot:\n`%v`", expected, parsedCommit.Footer)
	}
	for i, v := range expected {
		if parsedCommit.Footer[i] != v {
			t.Errorf(`Expected parsedCommit.Footer[%d] to equal %s got %s`, i, v, parsedCommit.Footer[i])
		}
	}
}

// https://github.com/conventional-commits/conventionalcommits.org/issues/313
func TestParseConventionalCommitIssue313MultiLineWithBreaks(t *testing.T) {
	const commit = `feat: add new feature

Description of the new feature
more details
even more details

BREAKING CHANGE: You no longer have to use alertTitle and children to display your alert content. 

Use children only instead. This is because we found out that people actually display the "content" of the alert in alertTitle and not in both alertTitle and children (description).

Because of that, we're consolidating both. This also opens up opportunities to have other components used within an alert component like an accordion for Progressive Disclosure design patterns.

Closes #42
`
	parsedCommit := ParseConventionalCommit(commit)
	expected := []string{
		`BREAKING CHANGE: You no longer have to use alertTitle and children to display your alert content. 

Use children only instead. This is because we found out that people actually display the "content" of the alert in alertTitle and not in both alertTitle and children (description).

Because of that, we're consolidating both. This also opens up opportunities to have other components used within an alert component like an accordion for Progressive Disclosure design patterns.`,
		`Closes #42`,
	}
	if len(parsedCommit.Footer) != len(expected) {
		t.Errorf("Expected parsedCommit.Footer to equal \n`%v`\n\ngot:\n`%v`", expected, parsedCommit.Footer)
	}
	for i, v := range expected {
		if parsedCommit.Footer[i] != v {
			t.Errorf(`Expected parsedCommit.Footer[%d] to equal %s got %s`, i, v, parsedCommit.Footer[i])
		}
	}
}

func TestParseConventionalCommits(t *testing.T) {
	commits := []string{
		`feat: add new feature`,
		`fix: bug fix`,
	}
	parsedCommits := ParseConventionalCommits(commits)
	if parsedCommits.Len() != 2 {
		t.Errorf("Expected 2 items, got %d", parsedCommits.Len())
	}
}

func TestConventionalCommitsSortingNoBreakingChanges(t *testing.T) {
	commits := []string{
		`fix: bug fix`,
		`chore: foo bar`,
		`feat: add new feature`,
	}
	const expected = "feat"
	parsedCommits := ParseConventionalCommits(commits)
	sort.Sort(parsedCommits)
	if parsedCommits[0].Category != expected {
		t.Error("Failed to sort with no breaking changes")
		for _, c := range parsedCommits {
			t.Error(c.Category)
		}
	}
}

func TestConventionalCommitsSortingWithBreakingChanges(t *testing.T) {
	commits := []string{
		`fix: bug fix`,
		`chore: foo bar`,
		`feat: some new feature that breaks things

BREAKING CHANGE: bla bla
`,
		`feat: add new feature`,
	}
	parsedCommits := ParseConventionalCommits(commits)
	sort.Sort(parsedCommits)
	if parsedCommits[0].Category != "feat" && parsedCommits[0].Subject != "some new feature that breaks things" {
		t.Error("Failed to sort with no breaking changes")
	}
}

func TestConventionalCommitsSortingWithBreakingChangesShortcut(t *testing.T) {
	commits := []string{
		`fix: bug fix`,
		`chore: foo bar`,
		`feat!: some new feature that breaks things`,
		`feat: add new feature`,
	}
	parsedCommits := ParseConventionalCommits(commits)
	sort.Sort(parsedCommits)
	if parsedCommits[0].Category != "feat" && parsedCommits[0].Subject != "some new feature that breaks things" {
		t.Error("Failed to sort with no breaking changes")
	}
}
