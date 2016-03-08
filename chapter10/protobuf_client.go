package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	pb "github.com/Masterminds/go-in-practice/chapter10/userpb"
	"github.com/golang/protobuf/proto"
)

func main() {
	res, err := http.Get("http://localhost:8080")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer res.Body.Close()

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body: %v", err)
		os.Exit(1)
	}

	var u pb.User
	err = proto.Unmarshal(b, &u)
	if err != nil {
		fmt.Println("Error decoding response body: %v", err)
		os.Exit(1)
	}

	fmt.Println(u.GetName())
	fmt.Println(u.GetId())
	fmt.Println(u.GetEmail())
}
