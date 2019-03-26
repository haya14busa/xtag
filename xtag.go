package xtag

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/haya14busa/go-versionsort"
)

// FindLatest finds latest tag matched with give xtag (e.g. v1.2.x).
func FindLatest(xtag string, tags []string) (string, error) {
	if !strings.HasSuffix(xtag, ".x") {
		return "", fmt.Errorf("xtag must ends with '.x': %q", xtag)
	}
	pattern := "^" + strings.Replace(regexp.QuoteMeta(xtag), `\.x`, `(\.[0-9.]+)?`, -1) + "$"
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}
	latest := ""
	for _, tag := range tags {
		if re.MatchString(tag) && versionsort.Less(latest, tag) {
			latest = tag
		}
	}
	if latest != "" {
		return latest, nil
	}
	return "", fmt.Errorf("latest tag not found: %q", xtag)
}
