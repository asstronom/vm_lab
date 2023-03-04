package main

import (
	"fmt"

	"example.com/pkg/cpu"
)

func main() {
	_ = cpu.NewCPU(nil)
	fmt.Println("Hello")
}
