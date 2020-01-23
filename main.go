package main

import (
	"flag"
	"fmt"
)

var someFlag bool

func init() {
	flag.BoolVar(&someFlag, "f", false, "An example bool flag")
	flag.Parse()
}

func main() {
	fmt.Println(someFlag)
}
