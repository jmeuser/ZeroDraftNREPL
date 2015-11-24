package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Word string

// A Sentence is a stack of Words.
type Sentence []Word

// Introduce a Word to the end of a Sentence.
func (s Sentence) intr(w Word) Sentence {
	return append(s, w)
}

// Eliminate the Word at the end of a Sentence.
func (s Sentence) elim() Sentence {
	return s[:len(s)-1]
}

// last Word of a Sentence.
func (s Sentence) last() Word {
	return s[len(s)-1]
}

func (s Sentence) String() string {
	var ss []string
	for _, w := range s {
		ss = append(ss, string(w))
	}
	return strings.Join(ss, " ")
}

type Noun uint // zero draft simplification
type Verb func(...Noun) Noun
type Monad Verb
type Dyad Verb
type Adverb func(...Verb) Verb

func main() {
	x := read()
	fmt.Println(x)
	var s Sentence
	s = s.intr("test")
	s = s.intr("super test")
	fmt.Println(s.last())
	s = s.elim()
	fmt.Println(s)
}

func read() Sentence {
	fmt.Print(" ")                                      // prompt
	in, _ := bufio.NewReader(os.Stdin).ReadString('\n') // read line from prompt
	//  ^ ignores reader errror!!!
	in = strings.TrimSpace(in)
	ss := regexp.MustCompile(" +").Split(in, -1) // assumes space separated words
	var out Sentence
	for _, s := range ss {
		out = out.intr(Word(s))
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
