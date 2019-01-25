package main

import (
	"fmt"
	"os"

	flag "github.com/spf13/pflag"
)

func main() {
	flag.Usage = Usage
	flag.Parse()

	gstr, err := NewGister()

	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(1)
	}

	for _, arg := range flag.Args() {
		gstr.Add(arg)
	}

	g, err := gstr.Save()

	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}

	fmt.Println(g.GetHTMLURL())
}
