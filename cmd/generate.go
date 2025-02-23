package cmd

import (
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gera um código TOTP",
	Long:  "Gera um código TOTP válido para autenticação de dois fatores.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(generateCmd)
	generateCmd.Flags().String("account", "", "Nome da conta (ex: gitlab)")
}
