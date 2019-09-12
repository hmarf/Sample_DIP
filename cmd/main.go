package main

import (
	"github.com/hmarf/sample_interface/down"
	"github.com/hmarf/sample_interface/up"
)

func main() {
	down := down.NewDown()
	up := up.NewUp(down)
	up.Up()
}
