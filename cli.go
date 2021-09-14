package main

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
)

const (
	appName    = "^--[fakeSMTPserver]--^"
	appVersion = "0.1"
)

func cliInitialize() {
	app := &cli.App{
		Name:  "fakeSMTPServer",
		Usage: "local development SMTP server",
		Action: func(c *cli.Context) error {
			fmt.Printf("%v v%v\ntype %v -help for help\n", appName, appVersion, os.Args[0])
			return nil
		},

		Commands: []*cli.Command{
			{
				Name:    "start",
				Aliases: []string{"s"},
				Usage:   "start listening",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "addr",
						Aliases: []string{"a"},
						Value:   ":2525",
						Usage:   "address:port",
					},

					&cli.StringFlag{
						Name:    "domain",
						Aliases: []string{"d"},
						Value:   "localhost",
						Usage:   "domain",
					},
				},

				Action: func(c *cli.Context) error {
					config := Config{
						Address:           c.String("addr"),
						Domain:            c.String("domain"),
						AllowInsecureAuth: true,
					}

					server := NewServer(config)
					return server.ListenAndServe()
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
