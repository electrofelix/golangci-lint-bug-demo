package main

import (
	"fmt"

	"github.com/electrofelix/golangci-lint-bug-demo/pkg/cmd"
)

func main() {
	fmt.Println("vim-go") // nolint: forbidigo

	l := cmd.Lister{}
	l.Emit()
}
