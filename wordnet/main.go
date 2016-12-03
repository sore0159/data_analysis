package main

import "fmt"
import "github.com/fluhus/gostuff/nlp/wordnet"

func main() {
	wn, err := wordnet.Parse("/usr/local/WordNet-3.0/dict")
	if err != nil {
		fmt.Println("PARSE ERROR:", err)
		return
	}
	fmt.Println("LEN:", len(wn.Synset))
	cn := wn.Search("cat")["n"][0]
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
