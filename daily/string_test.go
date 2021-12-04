package daily

import (
	"fmt"
	"testing"
)

func Test_replaceSpace(t *testing.T) {
	s := "We are happy."
	buf := []byte(s)
	fmt.Println(buf)

	fmt.Println(replaceSpace(s))
}

func TestReverseLeftWords(t *testing.T) {
	s := reverseLeftWords("abcdefg", 2)
	fmt.Println(s)
}
