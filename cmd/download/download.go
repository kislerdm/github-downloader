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
