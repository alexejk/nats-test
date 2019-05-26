package main

import (
	"github.com/alexejk/nats-test/cmd"
	"github.com/sirupsen/logrus"
)

func main() {

	if err := cmd.RootCmd().Execute(); err != nil {
		logrus.Fatal(err)
	}

}
