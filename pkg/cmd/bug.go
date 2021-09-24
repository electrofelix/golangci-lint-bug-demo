package cmd

// Integrates with code generated from the swagger spec

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

type Lister struct {
	logger   *logrus.Logger
	maxCount int
}

func (l *Lister) Emit() {
	fmt.Println("vim-go")

	log := logrus.New()
	log.Info("some message")
}
