package main

import (
	"context"
	"fmt"
	"io/ioutil"

	"github.com/google/go-github/github"
)

type gistfm map[github.GistFilename]github.GistFile

// Gister creates a gist with a list of files
type Gister struct {
	client      *github.Client
	description string
	files       []string
	public      bool
}

// NewGister creates a new Gister
func NewGister() (*Gister, error) {
	c, err := githubClient()

	if err != nil {
		return nil, fmt.Errorf("failed to create github client - %v", err)
	}

	return &Gister{
		client:      c,
		description: "",
		files:       make([]string, 0),
		public:      false,
	}, nil
}

// Add adds a file to the gister
func (g *Gister) Add(name string) {
	g.files = append(g.files, name)
}

// Describe adds a description to the gist
func (g *Gister) Describe(s string) {
	g.description = s
}

// Save posts the gist to gist.github.com and returns the gist object
func (g *Gister) Save() (*github.Gist, error) {

	gfm, err := g.fileMap()

	if err != nil {
		return nil, err
	}

	gist, _, err := g.client.Gists.Create(context.Background(), &github.Gist{
		Description: &g.description,
		Files:       gfm,
		Public:      &g.public,
	})

	if err != nil {
		return nil, err
	}

	return gist, nil
}

func (g *Gister) fileMap() (gistfm, error) {
	fxs := make(gistfm)

	for _, fname := range g.files {
		content, err := ioutil.ReadFile(fname)

		if err != nil {
			return nil, fmt.Errorf("read failed - %v", err)
		}

		data := string(content)

		fxs[github.GistFilename(fname)] = gistFile(fname, data)
	}
	return fxs, nil
}

func gistFile(name string, content string) github.GistFile {
	return github.GistFile{
		Filename: &name,
		Content:  &content,
	}
}
