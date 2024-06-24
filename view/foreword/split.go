package foreword

import (
	"regexp"
)

var newLineRegex, _ = regexp.Compile("\r\n|\r|\n")

func Split(text string) []string {
	return newLineRegex.Split(text, -1)
}
