package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

// A Word refers to an N event (action or object).
type Word string
func (w Word) isNoun() bool {
	re := "^[0-9]+$" // zero draft simplification
	b, _ := regexp.MatchString(re, string(w))
	return b
}

func (w Word) isVerb() bool {
	re := "^[!#$%&*+,-;<=>?@|]$"
	b, _ := regexp.MatchString(re, string(w))
	return b
}

func (w Word) isPronom() bool {
	re:= "^[A-Za-z]+$"
	b, _ := regexp.MatchString(re, string(w))
	return b
}

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

// VERY ineffecient where for sentences.
func (s Sentence) where(Q func(Word)bool) []int {
	var N []int
	for i, w := range s {
		if Q(w) {
			N = append(N,i)
		}
	}
	return N
}

func (s Sentence) String() string {
	var ss []string
	for _, w := range s {
		ss = append(ss, string(w))
	}
	return strings.Join(ss, " ")
}

type Noun Word
type Verb Word
type Adverb Word
type Pronom Word
type Copula Word

func main() {
	x := read()
	fmt.Println(x)
	fmt.Println("Words:",len(x))
	fmt.Println("Nouns:",x.where(Word.isNoun))
	fmt.Println("Verbs:", x.where(Word.isVerb))
	fmt.Println("Pronoms:", x.where(Word.isPronom))
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
