package main

import (
	"fmt"
	"os"
)

func main()  {
	fmt.Print("Hello World")
	if len(os.Args) > 1 {
		fmt.Print(os.Args[1])
	}
	os.Exit(0)
}
