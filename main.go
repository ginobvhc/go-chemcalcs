package main

import (
	"fmt"
	"go-chemcalcs/models/fluid/compress"
)

func main() {
	fmt.Println(compress.VelSound(1.2, 18, 373))
}
