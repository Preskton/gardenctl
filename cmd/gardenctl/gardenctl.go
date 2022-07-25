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
	"math/rand"
	"os"
	"time"

	"github.com/preskton/gardenctl/pkg/cmd"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var cfg = &AppOptions{}

type AppOptions struct {
	verbose bool
}

func main() {
	options := cmd.GardenctlOptions{}

	command := cmd.NewGardenctlCommand(options)

	run(command)
	os.Exit(1)
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

func run(command *cobra.Command) int {
	Preloader()
	rand.Seed(time.Now().UnixNano())

	if err := command.Execute(); err != nil {
		return 1
	}
	return 0
}
