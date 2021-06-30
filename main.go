package main

import (
	"flag"
	"log"
	"net/url"
	"os"
	"path"
	"runtime"
	"strings"

	client "github.com/kislerdm/github-download/external"
)

type RepoDetails struct {
	Owner, Repo, Branch, Input string
}

func newRepoDetails(urlDownload string) *RepoDetails {
	uParse, err := url.Parse(urlDownload)
	if err != nil {
		return nil
	}
	pathParts := strings.Split(uParse.Path[1:], "/")
	return &RepoDetails{
		Owner:  pathParts[0],
		Repo:   pathParts[1],
		Branch: pathParts[3],
		Input:  strings.Join(pathParts[4:], "/"),
	}
}

type Args struct {
	Token        string
	OutputPrefix string
	Details      *RepoDetails
}

var args Args

func init() {
	var input string

	_, filename, _, _ := runtime.Caller(0)
	outputPrfixDefault := path.Dir(filename)

	flag.StringVar(&args.Token, "token", os.Getenv("GITHUB_TOKEN"), "github token, default it taken from envvar")
	flag.StringVar(&args.OutputPrefix, "output", outputPrfixDefault, "prefix where to store downloaded files to")
	flag.StringVar(&input, "input", "", "url to the object/blob/tree to download from github")
	flag.Parse()
	if args.Token == "" {
		log.Fatalln("github token must be provided as 'GITHUB_TOKEN' envvar, or a the flag -token")
	}
	if input == "" {
		log.Fatalln("path to download not specified")
	}
	args.Details = newRepoDetails(input)
}

func main() {
	repoClient := client.New(args.Token, args.Details.Owner, args.Details.Repo, args.Details.Branch)
	if errs := repoClient.DownloadAll(args.Details.Input, args.OutputPrefix); len(errs) > 0 {
		log.Fatalln(errs)
	}
}
