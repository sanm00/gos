package main

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"gapi_scaffold/cmd"
)

func main() {
	app := &cli.App{
		Name:  "go-restful-scaffold",
		Usage: "生成 RESTful API 项目脚手架",
		Commands: []*cli.Command{
			cmd.NewCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
