package main

import (
	"fmt"
	"strconv"
)

func main() {
	type op struct {
		typ int
		arg string
	}
	var str string
	fmt.Scan(&str)
	var ops int
	fmt.Scan(&ops)
	opArr := make([]op, 0)
	for i := 0; i < ops; i++ {
		var (
			typ int
			arg string
		)
		fmt.Scan(&typ, &arg)
		opArr = append(opArr, op{typ: typ, arg: arg})
	}
	for _, op := range opArr {
		switch op.typ {
		case 1:
			str = fmt.Sprintf("%s%v", str, op.arg)
		case 2:
			pos, _ := strconv.Atoi(op.arg)
			fmt.Println(query(str, pos-1))
		}
	}
}

func query(str string, pos int) int {
	size := len(str)
	step := size
	found := false
	for i := pos + 1; i < size; i++ {
		if str[i] == str[pos] {
			found = true
			if i-pos < step {
				step = i - pos
			}
		}
	}
	for i := pos - 1; i >= 0; i-- {
		if str[i] == str[pos] {
			found = true
			if pos-i < step {
				step = pos - i
			}
		}
	}
	if found {
		return step
	}
	return -1
}
