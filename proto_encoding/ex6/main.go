package main

import (
	embedded "github.com/tng016/go_with_examples/proto_encoding/ex6/proto"
	"github.com/tng016/go_with_examples/proto_encoding/helper"
	"github.com/golang/protobuf/proto"
)

func main() {
	m1 := &embedded.EmbeddedMessage{
		I: proto.Uint64(1),
		S: proto.String("123"),
		B: proto.Bool(true),
	}

	m2 := &embedded.TopMessage{
		M: m1,
	}

	helper.PrintBytes(m2)
	helper.PrintBinary(m2)
}
