package piscine_go

import (
	"github.com/01-edu/z01"
)

func LoafOfBread2(str string) string {
	counter := 0
	for _, r := range str {
		if r != ' ' {
			counter++
		}
	}
	if counter < 5 {
		return "Invalid Output\n"
	}

	result := ""
	group := ""
	nonSpaceCount := 0
	skipNext := false

	for _, r := range str {
		if r == ' ' {
			continue
		}

		if skipNext {
			skipNext = false
			continue
		}

		group += string(r)
		nonSpaceCount++

		if nonSpaceCount%5 == 0 {
			result += group + "\n"
			group = ""
			skipNext = true
		}
	}

	for _, r := range result {
		z01.PrintRune(r)
	}

	return result
}
