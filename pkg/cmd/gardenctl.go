package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

type GardenctlOptions struct {
}

func NewGardenctlCommand(options GardenctlOptions) *cobra.Command {
	gardenctlCommand := &cobra.Command{
		Use:   "gardenctl",
		Short: "Cuddle your garden.",
		Long:  `Place some nice, long text here later.`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}

	gardenctlCommand.AddCommand(NewPlantCommand(options))

	return gardenctlCommand
}

func NewPlantCommand(options GardenctlOptions) *cobra.Command {
	var cmdEcho = &cobra.Command{
		Use:   "echo [string to echo]",
		Short: "Echo anything to the screen",
		Long: `echo is for echoing anything back.
	Echo works a lot like print, except it has a child command.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Echo: " + strings.Join(args, " "))
		},
	}

	return cmdEcho
}
