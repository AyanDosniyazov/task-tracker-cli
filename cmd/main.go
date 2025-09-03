package main

import (
	"fmt"
	"os"
)

func main() {
	for arg := range os.Args {
		fmt.Println(os.Args[arg])
	}
}
