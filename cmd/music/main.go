package main

import (
	"fmt"
	"os"

	"github.com/mattfenwick/algorithms/pkg/music/cli"
)

func main() {
	fmt.Printf("args: %+v\n", os.Args)
	cli.RunRootCommand()
}
