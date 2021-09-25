package main

import (
	"fmt"
	// go get
//go get -u github.com/gobuffalo/packr/...
	"github.com/gobuffalo/packr"
)

func main() {
	box := packr.NewBox("../filehandling")
	data := box.String("test.txt")
	fmt.Println("Contents of file:", data)
}