package cmd

import (
	"github.com/laetificat/sftpchanger/src/core"
	"fmt"
	"github.com/spf13/cobra"
)

var changeCmd = &cobra.Command{
	Use:   "change",
	Short: "Changes to the given config",
	Long: `
Changes the SFTP config to the given config, will figure out the correct one if nothing is given.`,
	Run: func(cmd *cobra.Command, args []string) {
		var configName string
		if len(args) > 0 {
			configName = args[0]
		}

		configName, err := core.ChangeConfig(configName)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Printf("Changed to config %s\n", configName)
	},
}

func init() {
	rootCmd.AddCommand(changeCmd)
}
