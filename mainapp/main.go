package main

import (
	"flag"
	"fmt"
	clr "rameshify.com/cleaner/cleaner"
	"strings"
)

func main() {
	var root string
	var confirmFlag string
	confirm := true

	flag.StringVar(&root, "root", "./", "root to workspace/projects directory")
	flag.StringVar(&confirmFlag, "confirm", "y", "confirm delete?")

	flag.Parse()

	response := strings.ToLower(strings.TrimSpace(confirmFlag))

	if response == "n" || response == "no" {
		confirm = false
	}

	fmt.Println(root)

	cleaner := clr.Cleaner{
		Root:    root,
		Confirm: confirm,
		Names:   []string{"gone"},
		FileType:   clr.Directory,
	}

	cleaner.Start()
}
