package cmd

import (
	"os"

	"github.com/hcd233/aris-api-tmpl/internal/tool/lintstatic"
	"github.com/spf13/cobra"
)

var lintCmd = &cobra.Command{
	Use:   "lint",
	Short: "Lint Command Group",
	Long:  `Lint command group for code quality checks.`,
}

var staticLintCmd = &cobra.Command{
	Use:   "static",
	Short: "Run static analysis",
	Long:  `Run go vet and golangci-lint if it is installed.`,
	Run: func(_ *cobra.Command, args []string) {
		result := lintstatic.Run(args)
		result.Write(os.Stdout)
		if result.Err != nil {
			os.Exit(1)
		}
	},
}

func init() {
	lintCmd.AddCommand(staticLintCmd)
	rootCmd.AddCommand(lintCmd)
}
