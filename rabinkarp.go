package rabinkarp

import "bytes"
import "fmt"

type Needle struct {
	text   []byte
	length int
	hash   uint
	m      uint
}

const prime uint = 101
const UINT_MAX = ^uint(0)

func hash(text []byte) (h uint, m uint) {
	var tmp uint = 1
	m = 1
	h = 0
	for i := len(text) - 1; i >= 0; i-- {
		tmp = m
		h = (h + ((uint(text[i]) * m) % UINT_MAX)) % UINT_MAX
		m = (m * prime) % UINT_MAX
	}
	return h, tmp
}

func update(hash uint, m uint, remove byte, add byte) uint {
	hash = (hash - (((m * uint(remove)) % UINT_MAX) % UINT_MAX)) % UINT_MAX
	hash = (hash * prime) % UINT_MAX
	hash = (hash + uint(add)) % UINT_MAX
	return hash
}

func NewNeedle(needle []byte) *Needle {
	n := new(Needle)
	n.text = needle
	n.length = len(needle)
	n.hash, n.m = hash(needle)
	return n
}

func (self *Needle) Index(haystack []byte) (index int) {
	fmt.Println("Looking for: ", self.hash)
	l := len(haystack)
	fmt.Println("Length: ", l)
	max := l - self.length
	fmt.Println("Max: ", max)
	if l < self.length {
		return -1
	}
	h, _ := hash(haystack[0:self.length])
	for i := 0; i < max; i++ {
		fmt.Println("Current hash: ", h)
		if h == self.hash && bytes.Compare(haystack[i:i+self.length], self.text) == 0 {
			return i
		}
		h = update(h, self.m, haystack[i], haystack[i+self.length])
	}
	return -1
}

/*
Return the index of the first instance of `needle` in `haystack`,
or -1 if not found.
*/
func Index(haystack, needle string) int { return NewNeedle([]byte(needle)).Index([]byte(haystack)) }
