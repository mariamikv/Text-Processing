package strategies

import (
	"regexp"
	"strings"
)

type RegexFilter struct {
	Pattern string
}

func (f *RegexFilter) ProcessText(text string) (string, error) {
	re := regexp.MustCompile(f.Pattern)
	filteredLines := make([]string, 0)
	for _, line := range strings.Split(text, "\n") {
		if re.MatchString(line) {
			filteredLines = append(filteredLines, line)
		}
	}
	return strings.Join(filteredLines, "\n"), nil
}
