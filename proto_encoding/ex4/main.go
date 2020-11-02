package main

import (
	protostrings "github.com/tng016/go_with_examples/proto_encoding/ex4/proto"
	"github.com/tng016/go_with_examples/proto_encoding/helper"
	"github.com/golang/protobuf/proto"
)

func main() {
	m1 := &protostrings.ProtoString{
		S: proto.String("123"),
	}

	helper.PrintBytes(m1)
	helper.PrintBinary(m1)
}
