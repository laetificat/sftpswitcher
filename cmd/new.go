package cmd

import (
	"bufio"
	"fmt"
	"os"

	"github.com/laetificat/sftpchanger/src/core"
	"github.com/spf13/cobra"
)

var configName string

var newCmd = &cobra.Command{
	Use:   "new",
	Short: "Creates a new config file for a project",
	Long: `
Creates a new config file with the given project name,
if no name is given it will try to guess by using the current directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		reader := bufio.NewReader(os.Stdin)
		result, err := core.CreateConfig(reader, configName)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(result)
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	newCmd.PersistentFlags().StringVarP(&configName, "name", "n", "", "the name to use for the config file")
}
