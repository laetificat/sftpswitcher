package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// rootCmd represents the changesftp command
var rootCmd = &cobra.Command{
	Use:   "changesftp",
	Short: "SFTP switcher allows you to easily switch SFTP config for VSCode",
	Long: `
SFTP switcher is made for the SFTP extension for VSCode.

Sometimes you have a project that is used in multiple workspaces, but the
SFTP extension does not support multiple servers in a single configuration
file. With this command you can save multiple configurations and easily
switch between them.

Use list to show all the available configurations in a list.
Use change [config name] to change to the given config
Use new -n [config name] to create a new config with the given name`,
}

/*
Execute runs the command.
*/
func Execute() error {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
	}

	return nil
}
