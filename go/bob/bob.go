// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package bob should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package bob

import (
	"strings"
	"unicode"
)

// Hey should have a comment documenting it.
func Hey(remark string) string {
	for _, letter := range remark {
		if letter != ' ' && letter != '\t' && letter != '\n' {
			var hasCap = false
			for _, letter := range remark {
				if letter >= 'A' && letter <= 'Z' {
					hasCap = true
				} else if unicode.IsLetter(letter) {
					if remark[len(remark)-1] == '?' {
						return "Sure."
					}
					return "Whatever."
				}
			}
			if hasCap {
				if strings.ContainsAny(remark, "?") {
					return "Calm down, I know what I'm doing!"
				}
				return "Whoa, chill out!"
			} else {
				if remark[len(remark)-1] == '?' {
					return "Sure."
				}
				return "Whatever."
			}
		}
	}
	return "Fine. Be that way!"
}
