package up

import (
	"fmt"

	"github.com/hmarf/sample_interface/down"
)

func Up() {
	fmt.Println("up")
	down.Down()
}
