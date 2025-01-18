package main

import (
	"os"

	"github.com/mattfenwick/algorithms/pkg/music/cli"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.Infof("args: %+v\n", os.Args)
	cli.RunRootCommand()
}
