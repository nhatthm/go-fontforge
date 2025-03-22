package fontforge

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	semver "github.com/Masterminds/semver/v3"
)

// semVerRegex is the regular expression used to parse a semantic version.
const semVerRegex string = `v?([0-9]+)(\.[0-9]+)?(\.[0-9]+)?` +
	`(-([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?` +
	`(\+([0-9A-Za-z\-]+(\.[0-9A-Za-z\-]+)*))?`

var versionRegex = regexp.MustCompile("^" + semVerRegex + "$")

func newVersion(v string) (*semver.Version, error) {
	m := versionRegex.FindStringSubmatch(v)
	if m == nil {
		return nil, semver.ErrInvalidSemVer
	}

	metadata := m[8]
	pre := m[5]
	minor := uint64(0)
	patch := uint64(0)

	var err error

	major, err := strconv.ParseUint(m[1], 10, 64)
	if err != nil {
		return nil, fmt.Errorf("error parsing version segment: %w", err)
	}

	if m[2] != "" {
		minor, err = strconv.ParseUint(strings.TrimPrefix(m[2], "."), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing version segment: %w", err)
		}
	}

	if m[3] != "" {
		patch, err = strconv.ParseUint(strings.TrimPrefix(m[3], "."), 10, 64)
		if err != nil {
			return nil, fmt.Errorf("error parsing version segment: %w", err)
		}
	}

	return semver.NewVersion(semver.New(major, minor, patch, pre, metadata).String())
}

func parseVersion(v string) *semver.Version {
	// Sanitize the version.
	v, _, _ = strings.Cut(v, ";")
	v = strings.TrimSpace(v)
	v = buildPattern.ReplaceAllString(v, "+$1")

	r, _ := newVersion(v) //nolint: errcheck

	return r
}
