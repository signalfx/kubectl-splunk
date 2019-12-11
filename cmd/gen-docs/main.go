package main

import (
	"github.com/signalfx/kubectl-signalfx/pkg/cli"
	"log"

	"github.com/spf13/cobra/doc"
)

func main() {
	err := doc.GenMarkdownTree(cli.RootCmd, "docs")
	if err != nil {
		log.Fatal(err)
	}
}
