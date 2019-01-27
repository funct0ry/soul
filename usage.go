package main

import (
	"os"
	"text/template"
)

// UsageTemplate is a text template for the usage message
const UsageTemplate = `soul version {{.Version}}
Usage: soul [options] FILE*

        --login                      Authenticate gist on this computer.
    -f, --filename [NAME.EXTENSION]  Specify filename and syntax type.
    -p, --private                    Create a private gist.
    -d, --description DESCRIPTION    Adds a description to your gist.
    -l, --list [USER]                List all gists for user.
    -r, --read ID [FILENAME]         Read a gist and print out the contents.
        --delete [ URL | ID ]        Delete a gist
    -h, --help                       Show this message.
    -v, --version                    Print the version.
`

// Usage prints the help message to stderr
func Usage() {
	t := template.Must(template.New("letter").Parse(UsageTemplate))
	t.Execute(os.Stderr, struct {
		Version string
	}{
		Version: Version,
	})
}
