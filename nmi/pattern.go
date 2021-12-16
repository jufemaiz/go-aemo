package nmi

import (
	"fmt"
	"regexp"
)

// Pattern for making regexes nicer with pattern matching.
type Pattern string

// NewPattern validates a new pattern, return invalid pattern if required.
func NewPattern(s string) (Pattern, error) {
	if _, err := regexp.Compile(s); err != nil {
		return Pattern(s), fmt.Errorf("pattern '%s': %w", s, ErrPatternInvalid)
	}

	return Pattern(s), nil
}

// Compile compiles the pattern, returning the regexp and an error.
func (p Pattern) Compile() (*regexp.Regexp, error) {
	r, err := regexp.Compile(p.String())
	if err != nil {
		return r, fmt.Errorf("Compile: %w", err)
	}

	return r, nil
}

// String returns the pattern as a string.
func (p Pattern) String() string {
	return string(p)
}

// Match returns true if the pattern matches a given string.
func (p Pattern) Match(s string) bool {
	regex, err := p.Compile()
	if err != nil {
		return false
	}

	return regex.MatchString(s)
}
