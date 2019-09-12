package main

import (
	"fmt"

	"github.com/hmarf/sample_interface/sub"
)

func main() {
	sub := sub.NewSub()
	fmt.Println(sub.Sub())
}
