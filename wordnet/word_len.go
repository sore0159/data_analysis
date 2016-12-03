package main

import "unicode/utf8"

func WordLen(str string) int {
	return utf8.RuneCountInString(str)
}
