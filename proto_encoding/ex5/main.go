package main

import (
	repeated "github.com/tng016/go_with_examples/proto_encoding/ex5/proto"
	"github.com/tng016/go_with_examples/proto_encoding/helper"
)

func main() {
	m1 := &repeated.RepeatedInteger{
		Integer:       []uint64{1, 2, 3},
		PackedInteger: []uint64{1, 2, 3},
	}

	helper.PrintBytes(m1)
	helper.PrintBinary(m1)
}
