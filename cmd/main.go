/*===========================================================================*\
*  MIT License Copyright (c) 2022 Preston Doster <preston.doster@genoq.com>   *
*                                                                             *
*                                888                            888    888    *
*                                888                            888    888    *
*                                888                            888    888    *
*   .d88b.   8888b.  888d888 .d88888  .d88b.  88888b.   .d8888b 888888 888    *
*  d88P"88b     "88b 888P"  d88" 888 d8P  Y8b 888 "88b d88P"    888    888    *
*  888  888 .d888888 888    888  888 88888888 888  888 888      888    888    *
*  Y88b 888 888  888 888    Y88b 888 Y8b.     888  888 Y88b.    Y88b.  888    *
*   "Y88888 "Y888888 888     "Y88888  "Y8888  888  888  "Y8888P  "Y888 888    *
*       888                                                                   *
*  Y8b d88P                                                                   *
*   "Y88P"                                                                    *
*                                                                             *
\*===========================================================================*/

package main

import (
	"fmt"
	"os"
	"time"

	"github.com/preskton/gardenctl"
	"github.com/preskton/gardenctl/internal/plant"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

var cfg = &AppOptions{}

type AppOptions struct {
	verbose bool
}

func main() {
	/* Change version to -V */
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"V"},
		Usage:   "The version of the program.",
	}
	app := &cli.App{
		Name:     gardenctl.Name,
		Version:  gardenctl.Version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			&cli.Author{
				Name:  gardenctl.AuthorName,
				Email: gardenctl.AuthorEmail,
			},
		},
		Copyright: gardenctl.Copyright,
		HelpName:  gardenctl.Copyright,
		Usage:     "Cuddle your garden. 🌱",
		UsageText: `service <options> <flags> 
A longer sentence, about how exactly to use this program`,
		Commands: []*cli.Command{
			&cli.Command{},
		},
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:        "verbose",
				Aliases:     []string{"v"},
				Destination: &cfg.verbose,
			},
		},
		EnableBashCompletion: true,
		HideHelp:             false,
		HideVersion:          false,
		Before: func(c *cli.Context) error {
			Preloader()
			fmt.Fprintf(c.App.Writer, gardenctl.Banner())
			return nil
		},
		After: func(c *cli.Context) error {
			// Destruct
			return nil
		},
		Action: func(c *cli.Context) error {

			//
			plantObject := plant.NewPlant()
			return plantObject.Run()
			//

		},
	}
	app.Run(os.Args)
}

// Preloader will run for ALL commands, and is used
// to initalize the runtime environments of the program.
func Preloader() {
	/* Flag parsing */
	if cfg.verbose {
		logrus.SetLevel(logrus.InfoLevel)
	} else {
		logrus.SetLevel(logrus.WarnLevel)
	}
}
