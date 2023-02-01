package main

import (
	"fmt"
	"os"

	"github.com/ismdeep/digest"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: digest-gen <password>")
		return
	}

	switch os.Args[1] {
	case "-h", "--help":
		fmt.Println("Usage: digest-gen <password>")
	default:
		fmt.Println(digest.Generate(os.Args[1]))
	}
}
