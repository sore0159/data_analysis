// Wordnet 3.0 used from
// https://wordnet.princeton.edu/
package main

import (
	"fmt"
	"github.com/fluhus/gostuff/nlp/wordnet"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: wordnet WORD [a|v|r|n]")
		return
	}
	fmt.Println("Loading wordnet...")
	wn, err := wordnet.Parse("/usr/local/WordNet-3.0/dict")
	if err != nil {
		fmt.Println("PARSE ERROR:", err)
		return
	}
	fmt.Println("Wordnet loaded!\nSYNSET TOTAL LEN:", len(wn.Synset))
	t := "n"
	if len(os.Args) == 2 {
		fmt.Println("Using default word-type 'noun'")
	} else if os.Args[2] == "n" || os.Args[2] == "v" || os.Args[2] == "r" || os.Args[2] == "a" {
		t = os.Args[2]
	} else {
		fmt.Println("Usage: wordnet WORD [a|v|r|n]")
		return
	}
	list := wn.Search(os.Args[1])[t]
	if len(list) == 0 {
		fmt.Println("No matches found!")
		return
	}
	fmt.Println(len(list), " synsets found!  Using first:")
	cn := list[0]
	fmt.Println("OFFSET: ", cn.Offset)
	fmt.Println("POS:", cn.Pos)
	fmt.Println("WORD:", cn.Word)
	fmt.Println("POINTER:", cn.Pointer)
	fmt.Println("FRAME:", cn.Frame)
	fmt.Println("GLOSS:", cn.Gloss)
	fmt.Println("Example:", cn.Example)

	pt := cn.Pointer[0]
	fmt.Println("SYBOL:", pt.Symbol)
	fmt.Println("SYNSET (TARGET):", pt.Synset)
	fmt.Println("SOURCE INDEX:", pt.Source)
	fmt.Println("TARGET INDEX:", pt.Target)

}
