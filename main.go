package main

import (
	"log"
	"os"

	"github.com/kislerdm/github-download/cli"
)

func main() {
	if err := cli.Run(os.Args); err != nil {
		if err.Error() != "" {
			log.Fatalln(err)
		}
	}
}
