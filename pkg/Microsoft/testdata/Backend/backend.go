package main

import (
	"fmt"

	"github.com/jdkato/prose/tag"
	"github.com/jdkato/prose/tokenize"
)

func main() {
	text := []string{
		"I'm a backend developer.",
		"I'm a back end developer.",
		"I'm a back-end developer.",
		"The bug was on the back end.",
		"The bug was on the back-end.",
		"Back end code is stored here.",
		"The back-end is written in Java.",
		"You can talk about the back end.",
		"But not about a back end developer.",
		"You can talk about the back end.",
		"You can talk about the back end, but not about a back end developer."}
	tagger := tag.NewPerceptronTagger()

	for _, sent := range text {
		fmt.Println("---")
		words := tokenize.NewTreebankWordTokenizer().Tokenize(sent)
		for _, tok := range tagger.Tag(words) {
			fmt.Println(tok.Text, tok.Tag)
		}
	}
}
