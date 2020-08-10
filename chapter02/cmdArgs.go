package main

import (
	"fmt"
	"os"
)

func main() {
	var args, space string
	for i := 1; i < len(os.Args); i++ {
		args += space + os.Args[i]
		space = " "
	}
	fmt.Println(args)
}
