package word

import (
	"fmt"
	"log"
	"testing"
)

func TestOne(t *testing.T) {
	log.Println("TEST ONE")
}

func TestTwo(t *testing.T) {
	log.Println("Loading wordnet...")
	wn, err := GetWordNet()
	if err != nil {
		log.Println("PARSE ERROR:", err)
		return
	}
	log.Println("Wordnet loaded!\nSYNSET TOTAL LEN:", len(wn.Synset))
	word := "cat"
	kind := "n"
	log.Println("Testing word ", word, "; Kind: ", kind)
	list := wn.Search(word)[kind]
	if len(list) == 0 {
		log.Println("No matches found!")
		return
	}
	log.Println(len(list), " synsets found!  Using first:")
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
