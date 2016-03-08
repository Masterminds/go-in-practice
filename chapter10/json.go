package main

import (
	"fmt"
	"os"

	"github.com/Masterminds/go-in-practice/chapter10/user"
	"github.com/ugorji/go/codec"
)

func main() {
	jh := new(codec.JsonHandle)
	u := &user.User{
		Name:  "Inigo Montoya",
		Email: "inigo@montoya.example.com",
	}

	var out []byte
	err := codec.NewEncoderBytes(&out, jh).Encode(&u)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(string(out))

	var u2 user.User
	err = codec.NewDecoderBytes(out, jh).Decode(&u2)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(u2)
}
