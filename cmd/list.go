package cmd

import "github.com/spf13/cobra"

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as contas e seus códigos TOTP",
	Long:  "Exibe todas as contas cadastradas, atualizando os códigos TOTP em tempo real.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
