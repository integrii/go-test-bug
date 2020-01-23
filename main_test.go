package main

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	fmt.Println(os.Args)
}

func TestFlag(t *testing.T) {
	t.Log("some test running")
}
