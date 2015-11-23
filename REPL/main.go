package main

import (
	"bufio"
	"fmt"
	"os"
)

type Word string
type Sentence []Word

type Noun uint // zero draft simplification
type Verb func(...Noun)Noun
type Monad Verb
type Dyad Verb
type Adverb func(...Verb)Verb

func main() {
	x := read()
	fmt.Println(rev(x))
}

func read() string {
	fmt.Print(" ")
	x, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	// ^ ignores reader errror!!!
	return x[:len(x)-1]
}

// Reverse a string with rev.
func rev(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}
