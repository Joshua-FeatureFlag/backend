package main

import (
	"fmt"

	"github.com/Joshua-FeatureFlag/proto/github.com/Joshua-FeatureFlag/proto/featureflag"
)

func main() {
	fmt.Println("Hello, World!")
	user := &featureflag.User{
		Id:   1,
		Name: "Alice",
	}
	fmt.Println(user)
}
