package main

import (
	"context"
	"errors"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

func githubClient() (*github.Client, error) {
	token := os.Getenv("GITHUB_AUTH_TOKEN")
	if token == "" {
		return nil, errors.New("authorization token not set")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	tc := oauth2.NewClient(ctx, ts)
	return github.NewClient(tc), nil
}
