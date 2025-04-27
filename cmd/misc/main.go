package main

import (
	"os"

	"github.com/mattfenwick/algorithms/pkg/pratt"
	"github.com/mattfenwick/algorithms/pkg/schema"
)

func main() {
	switch os.Args[1] {
	case "pratt":
		pratt.Run(os.Args[2])
	case "json":
		schema.Run(os.Args[2])
	}
}
