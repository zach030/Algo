package dp

import (
	"fmt"
	"testing"
)

func Test_strStr(t *testing.T) {
	a:=strStr("hello","ll")
	b:=kmp("abeababeabf","abeabf")
	fmt.Println(a,b)
}
