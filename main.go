package main

import (
	"fmt"

	"github.com/myrachanto/pdfs/src"
)

func main() {
	err := src.GeneratePdf("./test.pdf")
	if err != nil {
		fmt.Println(err)
	}
}
