package parser

import (
	"fmt"
	"regexp"
	"slices"
	"strings"
)

var (
	baseFormatRegex       = regexp.MustCompile(`(?is)^(?:(?P<category>[^\(!:]+)(?:\((?P<scope>[^\)]+)\))?(?P<breaking>!)?: (?P<subject>[^\n\r]+))[\n\r]?(?P<remainder>.*)`)
	bodyFooterFormatRegex = regexp.MustCompile(`(?isU)^(?:(?P<body>.*))?[\n\r]?(?P<footer>(?-U:(?:[\w\-]+(?:: | #).*|(?i:BREAKING CHANGE:.*))+))`)
	footerFormatRegex     = regexp.MustCompile(`(?s)^(?P<footer>(?i:(?:[\w\-]+(?:: | #).*|(?i:BREAKING CHANGE:.*))+))`)
)

func regExMapper(match []string, expectedFormatRegex *regexp.Regexp, result map[string]string) {
	for i, name := range expectedFormatRegex.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = match[i]
		}
	}
}

// ParseConventionalCommit takes a commits message and parses it into usable blocks.
func ParseConventionalCommit(message string) (commit *ConventionalCommit) {
	match := baseFormatRegex.FindStringSubmatch(message)

	parts := strings.Split(message, "\n")
	if len(match) == 0 {
		parts = append(parts, "")
		return &ConventionalCommit{
			Raw:      message,
			Category: "",
			Major:    strings.Contains(parts[1], "BREAKING CHANGE"),
			Header:   strings.Trim(parts[0], "\n"),
			Subject:  strings.Trim(parts[0], "\n"),
			Body:     strings.Join(parts[1:], "\n"),
		}
	}

	result := make(map[string]string)
	regExMapper(match, baseFormatRegex, result)

	// split the remainder into body & footer
	match = bodyFooterFormatRegex.FindStringSubmatch(result["remainder"])
	if len(match) > 0 {
		regExMapper(match, bodyFooterFormatRegex, result)
	} else {
		result["body"] = result["remainder"]
	}

	if slices.Contains(MajorCategories, result["category"]) {
		result["breaking"] = "!"
	}

	footers := []string{}
	for v := range strings.SplitSeq(result["footer"], "\n") {
		v = strings.Trim(v, "\n")
		if !footerFormatRegex.MatchString(v) && len(footers) > 0 {
			footers[len(footers)-1] += fmt.Sprintf("\n%s", v)
			continue
		}
		footers = append(footers, v)
	}

	for i := range footers {
		footers[i] = strings.Trim(footers[i], "\n")
		if footers[i] == "" { // Remove the element at index i from footers.
			copy(footers[i:], footers[i+1:])   // Shift a[i+1:] left one index.
			footers[len(footers)-1] = ""       // Erase last element (write zero value).
			footers = footers[:len(footers)-1] // Truncate slice.
		}
	}

	if len(footers) == 0 {
		footers = nil
	}

	commit = &ConventionalCommit{
		Raw:      message,
		Category: result["category"],
		Scope:    result["scope"],
		Major:    result["breaking"] == "!" || strings.Contains(result["footer"], "BREAKING CHANGE"),
		Subject:  result["subject"],
		Header:   strings.Trim(parts[0], "\n"),
		Body:     result["body"],
		Footer:   footers,
	}

	if commit.Major {
		return commit
	}

	if slices.Contains(MinorCategories, result["category"]) {
		commit.Minor = true
		return commit
	}

	if slices.Contains(PatchCategories, result["category"]) {
		commit.Patch = true
		return commit
	}

	return commit
}

func ParseConventionalCommits(messages []string) (commits ConventionalCommits) {
	for _, message := range messages {
		commits = append(commits, ParseConventionalCommit(message))
	}

	return
}
