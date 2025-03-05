package account

import (
	"github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Gerencie contas vinculadas a TOTP",
	Long:  "Adicione ou remova contas para geração de códigos TOTP de autenticação em dois fatores.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
