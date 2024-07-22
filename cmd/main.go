package main

import (
	"fmt"
	"os"

	"github.com/mattfenwick/algorithms/pkg/taxes"
)

func main() {
	fmt.Printf("args: %+v\n", os.Args)
	taxes.RunRootCommand()
}
