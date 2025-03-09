package main

import (
	"os"

	"github.com/mattfenwick/algorithms/pkg/pratt"
)

func main() {
	switch os.Args[1] {
	case "pratt":
		pratt.Run(os.Args[2])
	}
}
