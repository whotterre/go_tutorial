package ex_six_one

import (
	"bytes"
	"fmt"
)

type IntSet struct {
	words []uint64
}

func main() {
	x := IntSet{}
	x.Has(32)
}

// Check if a bit vector has an item
func (s *IntSet) Has(x int) bool {
	word, bit := x/64, uint(x%64)
	return (word < len(s.words)) && (s.words[word]&(1<<bit) != 0)
}

// Adds a number to the bit vector set
func (s *IntSet) Add(x int) {
	word, bit := x/64, uint(x%64)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.words[word] |= 1 << bit
}

// UnionWith set sets s to the union of s and t
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.words = append(s.words, tword)
		}

	}
}
//Prints out a set 
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < 64; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", 64*i+j)
			}
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
