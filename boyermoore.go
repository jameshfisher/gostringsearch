package boyermoore

func compute_prefix(str string, size int) []int {
	result := make([]int, size)

	k := 0

	for q := 1; q < size; q++ {
		for ; k > 0 && str[k] != str[q]; k = result[k-1] {
		}
		if str[k] == str[q] {
			k++
		}
		result[q] = k
	}

	return result
}

func Reverse(str string) string {
	// inefficient...
	l := len(str)
	m := l - 1
	s := make([]uint8, l)
	for i := 0; i < l; i++ {
		s[i] = str[m-i]
	}
	return string(s)
}

type Indexer struct {
	// Contains everything that can be derived purely from the needle
	needle       string // The original needle we're searching for
	length       int
	badcharacter map[uint8]int
	goodsuffix   []int
}

func NewIndexer(needle string) *Indexer {
	jt := new(Indexer)

	jt.needle = needle

	// Calculate string sizes
	jt.length = len(needle)

	jt.badcharacter = make(map[uint8]int)

	// Calculate the bad character table
	// For all chars in needle, set the jump to its position
	for i := 0; i < jt.length; i++ {
		jt.badcharacter[needle[i]] = i
	}

	// Calculate the good suffix table
	jt.goodsuffix = make([]int, jt.length+1)

	reversed := Reverse(needle)

	prefix_normal := compute_prefix(needle, jt.length)
	prefix_reversed := compute_prefix(reversed, jt.length)

	last := jt.length - prefix_normal[jt.length-1]

	for i := 0; i < jt.length; i++ {
		jt.goodsuffix[i] = last

		j := jt.length - prefix_reversed[i]
		k := i - prefix_reversed[i] + 1

		if jt.goodsuffix[j] > k {
			jt.goodsuffix[j] = k
		}
	}

	return jt
}

func (self *Indexer) Index(haystack string) int {
	if self.length == 0 {
		return 0
	} // As per the current behaviour of Index

	haystack_len := len(haystack)
	if haystack_len == 0 {
		return -1
	} // And the needle has positive length; hence is not found in null string

	// We only know we need this by this point
	difference_in_length := haystack_len - self.length

	// Do the actual search
	for index := 0; index <= difference_in_length; {

		// Read from the end of the candidate match until we find non-match
		j := self.length
		for ; j > 0 && self.needle[j-1] == haystack[index+j-1]; j-- {
		}

		if j > 0 {
			// The candidate did not match

			k, ok := self.badcharacter[haystack[index+j-1]] // Amount to shift based on the character on which we failed
			if !ok {
				k = self.length
			}

			m := j - k - 1

			// Shift the index forward
			if k < j && m > self.goodsuffix[j] {
				index += m
			} else {
				index += self.goodsuffix[j]
			}
		} else {
			// The candidate matched: we're done!
			return index
		}
	}
	return -1
}

/*
Return the index of the first instance of `needle` in `haystack`,
or -1 if not found.
*/
func Index(haystack, needle string) int { return NewIndexer(needle).Index(haystack) }