package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:  "phoenix",
	Long: "A collection of command line utilities and shorthands for nix development. ",
	Args: cobra.NoArgs,
}

var version string

type Variables struct {
	Version string
}

func Initialize(vars Variables) {
	version = vars.Version
	rootCmd.Version = version
}

func init() {
	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	cobra.EnableCommandSorting = true

	rootCmd.AddCommand(initCommand)
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}

	help, _ := rootCmd.Flags().GetBool("help")
	ver, _ := rootCmd.Flags().GetBool("version")

	if help || ver {
		os.Exit(0)
	}
}
