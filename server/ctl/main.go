package main

import (
	"fmt"

	"github.com/KobraKommander9/protobuf-language-server/server/ports/blank"
)

func main() {
	var n blank.Client
	n = Nothing{}
	fmt.Printf("server started %v\n", n)
}

type Nothing struct{}

func (Nothing) Close() error {
	return nil
}
