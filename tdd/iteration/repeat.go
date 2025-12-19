package iteration

import "strings"

func Repeat(s string, repeatCount int) string {
	var repeated strings.Builder
	for i := 0; i < repeatCount; i++ {
		repeated.WriteString(s)
	}
	return repeated.String()
}
