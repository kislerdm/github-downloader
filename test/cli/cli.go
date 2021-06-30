/*
Copyright © 2021 Dmitry Kisler <admin@dkisler.com>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cli

import (
	"fmt"

	"github.com/serverlessml/cli/cli/configure"
	"github.com/urfave/cli/v2"
)

// Run runs the CLI.
func Run(args []string) error {
	appName := "ServerlesML"

	app := cli.App{
		Name:                 appName,
		Usage:                fmt.Sprintf("%s CLI", appName),
		Description:          fmt.Sprintf("%s CLI to operate serlessML framework.", appName),
		Version:              "1.0.0b",
		Authors:              []*cli.Author{{Name: "Dmitry Kisler", Email: "admin@dkisler.com"}},
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{Name: "profile", Value: "default", Aliases: []string{"p"}, Usage: "use profile"},
		},
		Commands: []*cli.Command{
			{
				Name:  "configure",
				Usage: fmt.Sprintf("Configure %s", appName),
				Action: func(c *cli.Context) error {
					return configure.Run(c.String("profile"))
				},
			},
		},
	}

	return app.Run(args)
}
