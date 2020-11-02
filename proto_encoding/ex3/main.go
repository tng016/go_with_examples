package main

import (
	integers "github.com/tng016/go_with_examples/proto_encoding/ex3/proto"
	"github.com/tng016/go_with_examples/proto_encoding/helper"
	"github.com/golang/protobuf/proto"
)

func main() {
	m1 := &integers.SignedInteger{
		Integer:  proto.Int32(-1),
		SInteger: proto.Int32(-1),
	}

	helper.PrintBytes(m1)
	helper.PrintBinary(m1)
}
