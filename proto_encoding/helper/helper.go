package helper

import (
	"fmt"

	"github.com/golang/protobuf/proto"
)

func PrintBytes(m proto.Message) {
	fmt.Println("message:", m)

	bytes, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}
	fmt.Println(bytes)
}

func PrintBinary(m proto.Message) {
	fmt.Println("m:", m)

	bytes, err := proto.Marshal(m)
	if err != nil {
		panic(err)
	}

	for _, n := range bytes {
		fmt.Printf("% 08b", n) // prints 00000000 11111101
	}
	fmt.Println("")
}
