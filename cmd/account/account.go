package account

import (
	"github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Gerencie suas contas associadas a tokens TOTP",
	Long:  "Comandos para adicionar, remover e atualizar contas utilizadas na geração de códigos TOTP.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
