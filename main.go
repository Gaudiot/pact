package main

import (
	"fmt"
	"log"

	"github.com/pact/core"
)

func main() {
	err := core.LoadEnvConfig()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("Hello, World!")
}
