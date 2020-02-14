package main

import (
	"fmt"

	"golang.org/x/tools/imports"
)

func main() {
	fmt.Println(imports.LocalPrefix)
}
