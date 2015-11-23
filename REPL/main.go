package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Word string
type Sentence []Word

type Noun uint // zero draft simplification
type Verb func(...Noun) Noun
type Monad Verb
type Dyad Verb
type Adverb func(...Verb) Verb

func main() {
	x := read()
	fmt.Println(x)
}

func read() Sentence {
	fmt.Print(" ") // prompt
	in, _ := bufio.NewReader(os.Stdin).ReadString('\n') // read line from prompt
	// ^ ignores reader errror!!!
	in = strings.TrimSpace(in)
	ss := regexp.MustCompile(" +").Split(in,-1) // assumes space separated words
	var out Sentence
	for _, s := range ss {
		out = append(out, Word(s))
	}
	return out
}

// Reverse a string with rev.
func rev(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
