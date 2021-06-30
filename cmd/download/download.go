package download

import (
	"net/url"
	"strings"

	githubClient "github.com/kislerdm/github-download/external"
)

type repoDetails struct {
	Owner, Repo, Branch, Input string
}

func newRepoDetails(urlDownload string) *repoDetails {
	uParse, err := url.Parse(urlDownload)
	if err != nil {
		return nil
	}
	pathParts := strings.Split(uParse.Path[1:], "/")
	return &repoDetails{
		Owner:  pathParts[0],
		Repo:   pathParts[1],
		Branch: pathParts[3],
		Input:  strings.Join(pathParts[4:], "/"),
	}
}

func Download(token, inputURL, output string, verbose bool) error {
	d := newRepoDetails(inputURL)
	return githubClient.
		New(token, d.Owner, d.Repo, d.Branch).
		DownloadAll(d.Input, output, verbose)
}
