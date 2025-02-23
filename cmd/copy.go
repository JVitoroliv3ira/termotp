package cmd

import "github.com/spf13/cobra"

var copyCmd = &cobra.Command{
	Use:   "copy",
	Short: "Copia um código TOTP sem exibir",
	Long:  "Gera um código TOTP e copia automaticamente para a área de transferência.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func init() {
	rootCmd.AddCommand(copyCmd)
	copyCmd.Flags().String("account", "", "Nome da conta (ex: gitlab)")
}
