package main

import (
	"fmt"
	"os"

	"github.com/kaiiy/pretex/lib"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: pretex <tex-file-path>")
		return
	}

	srcFile := os.Args[1]

	srcText, err := os.ReadFile(srcFile)

	if err != nil {
		fmt.Println("Failed to read file:\n", err)
		return
	}

	destText := lib.Parse(string(srcText))

	os.WriteFile("./out.tex", []byte(destText), 0777)
}
