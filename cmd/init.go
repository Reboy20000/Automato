package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

var initCmd = &cobra.Command{
	Use:   "init <github-token>",
	Short: "Initialize Automato with a GitHub token",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		token := args[0]

		home, _ := os.UserHomeDir()
		dir := filepath.Join(home, ".automato")
		os.MkdirAll(dir, 0700)

		tokenFile := filepath.Join(dir, "token")
		err := os.WriteFile(tokenFile, []byte(token), 0600)
		if err != nil {
			fmt.Println("❌ Failed to save token")
			return
		}

		fmt.Println("✅ Automato initialized")
	},
}
