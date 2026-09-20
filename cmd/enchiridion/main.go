// Command enchiridion mirrors a Notion knowledge base as greppable markdown.
package main

import (
	"fmt"
	"os"
)

// version is injected at build time via -ldflags (see .goreleaser.yml).
var version = "dev"

// exitUnimplemented marks commands that are specced but not built yet.
const exitUnimplemented = 2

func main() {
	os.Exit(run(os.Args[1:]))
}

func run(args []string) int {
	if len(args) > 0 && args[0] == "version" {
		fmt.Printf("enchiridion %s\n", version)

		return 0
	}

	fmt.Fprintln(os.Stderr, "sync: not implemented yet — see specs/implementation-plan.md")

	return exitUnimplemented
}
