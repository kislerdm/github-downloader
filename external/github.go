package external

import (
	"context"
	"fmt"
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
	ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
	return &Client{
		Client: github.NewClient(oauth2.NewClient(ctx, ts)),
		Owner:  owner,
		Name:   repoName,
		Branch: branch,
	}
}

func (c *Client) getSHALastCommit(branch string) string {
	refList, _, _ := c.Client.Git.ListMatchingRefs(ctx, c.Owner, c.Name, nil)
	for _, ref := range refList {
		if strings.HasSuffix(ref.GetRef(), branch) {
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
func (c *Client) DownloadAll(pathDownload, prefixOutput string) (errs []error) {
	commitSHA := c.getSHALastCommit(c.Branch)
	if commitSHA == "" {
		return []error{fmt.Errorf("commits for branch '%s' not found", c.Branch)}
	}

	tree, err := c.getTreeSHAByCommitSHA(commitSHA)
	if err != nil {
		return []error{err}
	}

	var routine func(chan error) = func(ch chan error) {
		for _, entry := range tree.Entries {
			if entry.GetType() != "tree" {
				p := entry.GetPath()
				if strings.HasPrefix(p, pathDownload) {
					if err := c.Download(p, prefixOutput); err != nil {
						ch <- err
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
