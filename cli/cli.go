package cli

import (
	_ "embed"
	"fmt"
	"os"
	"strings"

	"github.com/kislerdm/github-download/cmd/download"
	"github.com/urfave/cli/v2"
)

//go:embed VERSION
var version string

const (
	appName  = "github-download"
	template = `{{.Name}} - the app to download a blob/file, or tree/directory from github repo.

USAGE: {{.Name}} [options] https://github.com/OWNER/REPO/tree/BRANCH/PATH_TO_DOWNLOAD

OPTIONS:
	{{range $index, $option := .VisibleFlags}}{{if $index}}
	{{end}}{{$option}}{{end}}

VERSION: {{.Version}}
LICENSE: {{.Copyright}}
AUTHOR{{with $length := len .Authors}}{{if ne 1 $length}}S:{{printf "\n"}}{{else}}:{{end}}{{end}}{{range $index, $author := .Authors}}{{if $index}}{{end}} {{$author}}{{printf "\n"}}{{end}}
`
)

// Run runs the CLI.
func Run(args []string) error {
	cli.VersionPrinter = func(c *cli.Context) {
		fmt.Printf("version=%s\n", c.App.Version)
	}

	app := cli.App{
		Name:                  appName,
		CustomAppHelpTemplate: template,
		Version:               strings.TrimSpace(version),
		Authors:               []*cli.Author{{Name: "Dmitry Kisler", Email: "admin@dkisler.com"}},
		EnableBashCompletion:  true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "token",
				Value:   os.Getenv("GITHUB_TOKEN"),
				Aliases: []string{"t"},
				Usage:   "github token, default is taken from envvar 'GITHUB_TOKEN'",
			},
			&cli.StringFlag{
				Name:    "output",
				Value:   "/tmp/",
				Aliases: []string{"o"},
				Usage:   "prefix where to store downloaded files to",
			},
			&cli.BoolFlag{
				Name:    "verbose",
				Value:   false,
				Aliases: []string{"b"},
				Usage:   "verbose output",
			},
		},
		Copyright: "MIT",
		Action: func(c *cli.Context) error {
			if c.NArg() == 0 {
				return fmt.Errorf("URL to download must be provided as an argument")
			}
			return download.Download(
				c.String("token"),
				c.Args().Get(0),
				c.String("output"),
				c.Bool("verbose"),
			)
		},
	}
	return app.Run(args)
}
