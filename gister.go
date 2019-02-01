package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/google/go-github/github"
)

type gistfm map[github.GistFilename]github.GistFile

// GistFile represents an unit in a gist
type GistFile struct {
	name   string
	source io.Reader
}

// NewGistFile instantiates a new GistFile instance
func NewGistFile(name string, r io.Reader) *GistFile {
	return &GistFile{
		name:   name,
		source: r,
	}
}

// Gister creates a gist with a list of files
type Gister struct {
	client      *github.Client
	description string
	files       []*GistFile
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
		files:       make([]*GistFile, 0),
		public:      true,
	}, nil
}

// Add adds a file to the gister
func (g *Gister) Add(name string, r io.Reader) {
	g.files = append(g.files, NewGistFile(name, r))
}

// Describe adds a description to the gist
func (g *Gister) Describe(s string) {
	g.description = s
}

// SetPrivate makes the gist private
func (g *Gister) SetPrivate(v bool) {
	g.public = !v
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

	for _, f := range g.files {
		content, err := ioutil.ReadAll(f.source)

		if err != nil {
			return nil, fmt.Errorf("read failed - %v", err)
		}

		data := string(content)

		fxs[github.GistFilename(f.name)] = gistFile(f.name, data)
	}
	return fxs, nil
}

func gistFile(name string, content string) github.GistFile {
	return github.GistFile{
		Filename: &name,
		Content:  &content,
	}
}
