package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	"github.com/laetificat/sftpchanger/src/config"
	"github.com/laetificat/sftpchanger/src/core"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lists all the available SFTP server configs",
	Long: `
Shows a list of all the available SFTP server configurations.`,
	Run: func(cmd *cobra.Command, args []string) {
		vscodeDirPath := filepath.Join(config.Cfg.ForestAppLocation, ".vscode")

		configList, err := core.GetAll(vscodeDirPath)
		if err != nil {
			fmt.Println(err)
		}

		writer := bufio.NewWriter(os.Stdout)
		core.ShowList(configList, writer)
		writer.Flush()
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
