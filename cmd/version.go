package cmd

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/version"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Exiba a versão do TermOTP",
	Long:  "Mostre a versão atual do TermOTP instalada no sistema.",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(version.GetVersion())
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
