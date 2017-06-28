package main

import (
	"os"

	"github.com/urfave/cli"

	"gopkg.in/go-playground/validator.v9"
)

var version string
var validate *validator.Validate
var apiBaseURL = "https://www.refuges.info/api"
var output string

func main() {
	app := cli.NewApp()
	validate = validator.New()

	app.Version = version
	app.Usage = "Refuges.info API CLI tool"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "output",
			Usage:       "Give a file name to save output",
			Destination: &output,
		},
	}

	app.Commands = []cli.Command{
		bboxCmd,
	}

	app.Run(os.Args)
}
