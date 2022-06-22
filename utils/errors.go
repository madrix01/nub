package utils

import (
	"fmt"
	"os"
)

func MakeError(err string) {
	fmt.Println("[ERROR] " + err)
	os.Exit(1)
}
