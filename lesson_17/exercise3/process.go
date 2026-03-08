package exercise3

import "regexp"

func processItemsSlow(items []string) []string {
	result := make([]string, 0, len(items))
	for _, s := range items {
		re := regexp.MustCompile(`\d+`)
		if re.MatchString(s) {
			result = append(result, re.ReplaceAllString(s, "NUM"))
		}
	}
	return result
}

var digitRe = regexp.MustCompile(`\d+`)

func processItemsFast(items []string) []string {
	result := make([]string, 0, len(items))
	for _, s := range items {
		if digitRe.MatchString(s) {
			result = append(result, digitRe.ReplaceAllString(s, "NUM"))
		}
	}
	return result
}
