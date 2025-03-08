package account

import (
	"github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Gerencie contas vinculadas a TOTP",
	Long:  "Adicione, remova ou renomeie contas vinculadas ao TOTP para geração de códigos de autenticação em dois fatores.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
