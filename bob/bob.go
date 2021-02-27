package bob

import (
	"regexp"
	"strings"
)

// answers stores Bob's answers
var answers = map[string]string {
	"questions": 			"Sure.",
	"yell": 				"Whoa, chill out!",
	"yellAndQuestion":		"Calm down, I know what I'm doing!",
	"address":				"Fine. Be that way!",
	"other":				"Whatever.",
}


// Hey
func Hey(remark string) string {

	if isQuestion(remark) && isYell(remark) {
		return answers["yellAndQuestion"]
	} else if isYell(remark) {
		return answers["yell"]
	} else if isQuestion(remark) {
		return answers["questions"]
	} else if isSilent(remark) {
		return answers["address"]
	} else {
		return answers["other"]
	}
}

// isYell returns if all chars are upper case
func isYell(remark string) bool {
	patternUpper, _ := regexp.Compile(`[[:upper:]]{2,}`)
	patternLower, _ := regexp.Compile(`[[:lower:]]+`)
	hasUpper := patternUpper.MatchString(remark)
	hasLower := patternLower.MatchString(remark)
	if hasLower {
		return false
	} else if hasUpper {
		return true
	}
	return false
}

// isQuestion removes whitespaces at the end of the string
func isQuestion(remark string) bool {
	remark = strings.TrimSpace(remark)
	return strings.HasSuffix(remark,"?")
}

// isSilent
func isSilent(remark string) bool {
	pattern, _ := regexp.Compile(`^[[:^alnum:]]*$`)
	return pattern.MatchString(remark)
}