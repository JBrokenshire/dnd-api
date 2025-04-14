package security

import "strings"

var badCharacters = `'"`
var badWords = []string{
	" select",
	"--",
	"1=1",
	`'1'='1'`,
	"src=",
	"script",
	"iframe",
	"oastify",
}

func IsURLBad(url string) bool {

	// Should always have a URL.
	if url == "" {
		return true
	}

	// Check for backtick
	if strings.Contains(url, "`") {
		return true
	}

	// check for bad characters
	if strings.ContainsAny(url, badCharacters) {
		return true
	}

	url = strings.ToLower(url)

	for _, badWord := range badWords {
		// Skip script check on /scribe-transcriptions routes
		if badWord == "script" && strings.Contains(url, "/scribe-transcriptions") {
			continue
		}
		if strings.Contains(url, badWord) {
			return true
		}
	}

	return false
}
