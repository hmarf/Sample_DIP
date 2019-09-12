package up

import (
	"fmt"

	"github.com/hmarf/sample_interface/down"
)

type upstruct struct {
	down down.DownInterface
}

type UpInterface interface {
	Up()
}

func NewUp(down down.DownInterface) UpInterface {
	return &upstruct{down: down}
}

func (up *upstruct) Up() {
	fmt.Println("up")
	up.down.Down()
}
