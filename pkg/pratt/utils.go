package pratt

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func Must0(err error) {
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
}

func Must[A any](a A, err error) A {
	if err != nil {
		fmt.Printf("%+v\n", err)
		logrus.Fatalf("%+v", err)
	}
	return a
}
