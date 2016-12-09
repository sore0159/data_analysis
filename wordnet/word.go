// Wordnet 3.0 used from
// https://wordnet.princeton.edu/
package word

import (
	"github.com/fluhus/gostuff/nlp/wordnet"
	"unicode/utf8"
)

func WordLen(str string) int {
	return utf8.RuneCountInString(str)
}

func GetWordNet() (*wordnet.WordNet, error) {
	return wordnet.Parse("/usr/local/WordNet-3.0/dict")
}
