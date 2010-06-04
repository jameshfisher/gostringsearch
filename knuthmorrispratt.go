package knuthmorrispratt

type Indexer struct {
	needle []byte
	length int
	table  []int
}

func NewIndexer(needle []byte) *Indexer {
	// Preprocess the needle, and return an Indexer for it.

	// Set the obvious properties first ...
	indexer := new(Indexer)
	indexer.needle = needle
	indexer.length = len(needle)
	indexer.table = make([]int, indexer.length)

	// Start building the table!

	// The first values are fixed
	indexer.table[0] = -1
	indexer.table[1] = 0

	pos := 2 // The current index into indexer.table we're computing
	cnd := 0 // Index into needle: next char of candidate substring

	for pos < indexer.length {
		if needle[pos-1] == needle[cnd] {
			// First case: the substring continues
			indexer.table[pos] = cnd + 1
			pos++
			cnd++
		} else if cnd > 0 {
			// Second case: the substring does not continue, but we can fall back
			cnd = indexer.table[cnd]
		} else {
			// Third case: we have run out of candidates.
			// Note: cnd == 0
			indexer.table[pos] = 0
			pos++
		}
	}
	return indexer
}

func (self *Indexer) Index(haystack []byte) (index int) {
	l := len(haystack)
	m := 0
	for i := 0; m+i < l; {
		if self.needle[i] == haystack[m+i] {
			i++
			if i == self.length {
				return m
			}
		} else {
			m += i - self.table[i]
			if self.table[i] > -1 {
				i = self.table[i]
			} else {
				i = 0
			}
		}
	}
	return -1
}

/*
Return the index of the first instance of `needle` in `haystack`,
or -1 if not found.
*/
func Index(haystack, needle string) int { return NewIndexer([]byte(needle)).Index([]byte(haystack)) }
