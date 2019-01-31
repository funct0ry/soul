package main

import (
	"fmt"
	"os"
	"path/filepath"

	flag "github.com/spf13/pflag"
)

var (
	helpRequested    = flag.BoolP("help", "h", false, "Show this message.")
	versionRequested = flag.BoolP("version", "v", false, "Print the version.")

	description = flag.StringP("description", "d", "", "Adds a description to your gist.")
)

func main() {
	flag.Usage = Usage
	flag.Parse()

	if *helpRequested {
		flag.Usage()
		os.Exit(0)
	}

	if *versionRequested {
		fmt.Fprintf(os.Stdout, "soul v%s\n", Version)
		os.Exit(0)
	}

	gstr, err := NewGister()

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}

	gstr.Describe(*description)

	if flag.NArg() < 1 {
		// Read from stdin
		gstr.Add("gistfile.txt", os.Stdin)
	}

	for _, arg := range flag.Args() {
		fileName := filepath.Base(arg)
		f, err := os.Open(arg)

		if err != nil {
			fmt.Fprintln(os.Stderr, "ERROR: opening file", err)
			continue
		}

		gstr.Add(fileName, f)
	}

	g, err := gstr.Save()

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println(g.GetHTMLURL())
}
