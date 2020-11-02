package main

import (
	"github.com/golang/protobuf/proto"
	integers "github.com/tng016/go_with_examples/proto_encoding/ex2.1/proto"
	"github.com/tng016/go_with_examples/proto_encoding/helper"
)

func main() {
	m1 := &integers.SingleInteger{
		Integer: proto.Uint64(18789),
	}

	helper.PrintBytes(m1)
	helper.PrintBinary(m1)
}
