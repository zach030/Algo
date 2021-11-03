package skiplist

import (
	"bytes"
)

type SkipList struct {
	header *Element
}

type Entry struct {
	Key   []byte
	Value []byte
}

type Element struct {
	entry  *Entry
	levels []*Element
	score  float64
}

func newElement(score float64, entry *Entry, level int) *Element {
	return &Element{
		levels: make([]*Element, level),
		entry:  entry,
		score:  score,
	}
}

var (
	ElemA = &Element{entry: &Entry{Key: []byte("1"), Value: []byte("va")}, levels: make([]*Element, 3), score: calcScore([]byte("1"))}
	ElemB = &Element{entry: &Entry{Key: []byte("2"), Value: []byte("vb")}, levels: make([]*Element, 1), score: calcScore([]byte("2"))}
	ElemC = &Element{entry: &Entry{Key: []byte("3"), Value: []byte("vc")}, levels: make([]*Element, 2), score: calcScore([]byte("3"))}
	ElemD = &Element{entry: &Entry{Key: []byte("4"), Value: []byte("vd")}, levels: make([]*Element, 1), score: calcScore([]byte("4"))}
	ElemE = &Element{entry: &Entry{Key: []byte("5"), Value: []byte("ve")}, levels: make([]*Element, 3), score: calcScore([]byte("5"))}
	ElemF = &Element{entry: &Entry{Key: []byte("6"), Value: []byte("vf")}, levels: make([]*Element, 1), score: calcScore([]byte("6"))}
	ElemG = &Element{entry: &Entry{Key: []byte("7"), Value: []byte("vg")}, levels: make([]*Element, 2), score: calcScore([]byte("7"))}
	ElemH = &Element{entry: &Entry{Key: []byte("8"), Value: []byte("vh")}, levels: make([]*Element, 1), score: calcScore([]byte("8"))}
	ElemI = &Element{entry: &Entry{Key: []byte("9"), Value: []byte("vi")}, levels: make([]*Element, 3), score: calcScore([]byte("9"))}
)

func NewSkipList() *SkipList {
	return &SkipList{header: ElemA}
}

func initRelation() {
	ElemA.levels[0] = ElemB
	ElemA.levels[1] = ElemC
	ElemA.levels[2] = ElemE

	ElemB.levels[0] = ElemC

	ElemC.levels[0] = ElemD
	ElemC.levels[1] = ElemE

	ElemD.levels[0] = ElemE

	ElemE.levels[0] = ElemF
	ElemE.levels[1] = ElemG
	ElemE.levels[2] = ElemI

	ElemF.levels[0] = ElemG

	ElemG.levels[0] = ElemH
	ElemG.levels[1] = ElemI

	ElemH.levels[0] = ElemI
}

// a(1)--------------------e(5)--------------------i(9)
// a(1)--------c(3)--------e(5)--------g(7)--------i(9)
// a(1)--b(2)--c(3)--d(4)--e(5)--f(6)--g(7)--h(8)--i(9)

func (s *SkipList) Search(key []byte) *Entry {
	prev := s.header
	i := len(s.header.levels) - 1
	score := calcScore(key)
	for i >= 0 {
		for next := prev.levels[i]; next != nil; next = prev.levels[i] {
			if comp := compare(score, key, next); comp <= 0 {
				if comp == 0 {
					return next.entry
				}
				break
			}
			prev = next
		}
		i--
	}
	return nil
}

func (s *SkipList) Add(entry *Entry) error {
	return nil
}

func calcScore(key []byte) (score float64) {
	var hash uint64
	l := len(key)
	if l > 8 {
		l = 8
	}
	for i := 0; i < l; i++ {
		shift := uint(64 - 8 - 8*i)
		hash |= uint64(key[i]) << shift
	}
	score = float64(hash)
	return
}

func compare(score float64, key []byte, next *Element) int {
	if score == next.score {
		return bytes.Compare(key, next.entry.Key)
	}

	if score < next.score {
		return -1
	} else {
		return 1
	}
}
