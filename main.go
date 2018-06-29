package main

import (
	"fmt"
	"os"

	"golang.org/x/sync/errgroup"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	var g errgroup.Group
	g.Go(func() error { return nil })
	return g.Wait()
}
