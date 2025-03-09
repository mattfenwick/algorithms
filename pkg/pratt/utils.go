package pratt

import "github.com/sirupsen/logrus"

func Must0(err error) {
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}
}

func Must[A any](a A, err error) A {
	if err != nil {
		logrus.Fatalf("%s", err.Error())
	}
	return a
}
