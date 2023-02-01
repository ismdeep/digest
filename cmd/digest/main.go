package main

import (
	"fmt"
	"os"

	"github.com/ismdeep/digest"
)

func help() {
	fmt.Println(`Usage: digest <password>

ABOUT:
    Repo: https://github.com/ismdeep/digest`)
}

func main() {
	if len(os.Args) < 2 {
		help()
		return
	}

	switch os.Args[1] {
	case "-h", "--help":
		help()
	default:
		fmt.Println(digest.Generate(os.Args[1]))
	}
}
