/* Copyright (c) 2021 Dmitry Kisler <dkisler.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
*/

package external

import (
	"context"
	"fmt"
	"log"
	"path"
	"strings"

	"github.com/google/go-github/v36/github"
	"golang.org/x/oauth2"
)

var ctx = context.Background()

type Client struct {
	Client *github.Client
	Owner  string
	Name   string
	Branch string
}

// New instantiate the client to interact with github API
// token: github API token
// owner: github repo owner
// repoName: github repo name
// branch: repo branch
func New(token, owner, repoName, branch string) *Client {
	c := github.NewClient(nil)
	if token != "" {
		ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
		c = github.NewClient(oauth2.NewClient(ctx, ts))

	}
	return &Client{
		Client: c,
		Owner:  owner,
		Name:   repoName,
		Branch: branch,
	}
}

func (c *Client) getSHALastCommit() string {
	refList, _, _ := c.Client.Git.ListMatchingRefs(ctx, c.Owner, c.Name, nil)
	for _, ref := range refList {
		if strings.HasSuffix(ref.GetRef(), c.Branch) {
			return ref.GetObject().GetSHA()
		}
	}
	return ""
}

func (c *Client) getTreeSHAByCommitSHA(sha string) (tree *github.Tree, err error) {
	commit, _, err := c.Client.Git.GetCommit(ctx, c.Owner, c.Name, sha)
	if err != nil {
		return
	}
	shaTree := commit.GetTree().GetSHA()
	tree, _, err = c.Client.Git.GetTree(ctx, c.Owner, c.Name, shaTree, true)
	return
}

// DownloadAll downloads all blobs by its path/prefix, or tree
func (c *Client) DownloadAll(pathDownload, prefixOutput string, verbose bool) (errs Errors) {
	commitSHA := c.getSHALastCommit()
	if commitSHA == "" {
		return Errors{fmt.Errorf("commits for branch '%s' not found", c.Branch)}
	}

	tree, err := c.getTreeSHAByCommitSHA(commitSHA)
	if err != nil {
		return Errors{err}
	}

	var routine func(chan error) = func(ch chan error) {
		for _, entry := range tree.Entries {
			if entry.GetType() != "tree" {
				p := entry.GetPath()
				if strings.HasPrefix(p, pathDownload) {
					if err := c.Download(p, prefixOutput); err != nil {
						ch <- err
					}
					if verbose {
						log.Printf("Download [%v bytes] %s", entry.GetSize(), p)
					}
				}
			}
		}
		close(ch)
	}

	ch := make(chan error)
	go routine(ch)
	for c := range ch {
		errs = append(errs, c)
	}
	return errs
}

// Download downloads the file (blob) from github repo
func (c *Client) Download(pathDownload, prefixOutput string) error {
	_, resp, err := c.Client.Repositories.DownloadContents(ctx, c.Owner, c.Name, pathDownload, nil)
	if err != nil {
		return err
	}
	if resp.StatusCode != 200 {
		return fmt.Errorf(resp.Status)
	}
	pathOut := path.Join(prefixOutput, pathDownload)
	if err := Write(resp.Body, pathOut); err != nil {
		return err
	}
	return nil
}
