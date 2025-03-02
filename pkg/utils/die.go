package utils

import "github.com/sirupsen/logrus"

func Die0(err error) {
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
}

func Die[A any](a A, err error) A {
	if err != nil {
		logrus.Fatalf("%+v", err)
	}
	return a
}
