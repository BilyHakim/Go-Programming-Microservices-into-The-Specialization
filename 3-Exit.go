package main

import (
	"fmt"
	"os"
)

func main() {
	defer fmt.Println("Invoke with defer")
	fmt.Println("Before Existing")
	os.Exit(1)
}
