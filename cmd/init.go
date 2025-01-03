package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"phoenix/lib"

	"github.com/spf13/cobra"
)

var initCommand = &cobra.Command{
	Use:   "init",
	Short: "Initialize a project directory",
	Args:  cobra.NoArgs,
}

func init() {
	f := initCommand.Flags()

	var (
		force         bool
		initDirectory string
	)

	f.BoolVarP(&force, "force", "f", false, "override existing files")
	f.StringVarP(&initDirectory, "directory", "d", "", "target directory")
	initCommand.Run = func(cmd *cobra.Command, args []string) {
		initDir, err := filepath.Abs(initDirectory)
		if err != nil {
			lib.Err(err)
		}
		if info, err := os.Stat(initDir); err == nil && info.IsDir() && !force {
			lib.Err("could not initialize: " + lib.Cyan(initDir) + " already exists")
		}

		lib.Mkdir(initDir)
		fmt.Println(lib.Green("project ready") + " at " + lib.Cyan(initDir))
	}
}
