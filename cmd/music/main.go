package main

import (
	"fmt"
	"os"

	"github.com/mattfenwick/algorithms/pkg/music"
)

func main() {
	fmt.Printf("args: %+v\n", os.Args)
	music.RunRootCommand()
}
