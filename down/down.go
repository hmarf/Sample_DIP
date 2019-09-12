package down

import "fmt"

type downstruct struct {
}

type DownInterface interface {
	Down()
}

func NewDown() DownInterface {
	return &downstruct{}
}

func (up *downstruct) Down() {
	fmt.Println("down")
}
