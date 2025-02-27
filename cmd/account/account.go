package account

import (
	"github.com/spf13/cobra"
)

var AccountCmd = &cobra.Command{
	Use:   "account",
	Short: "Gerencie contas vinculadas a TOTP",
	Long:  "Adicione, remova ou edite contas usadas na geração de códigos TOTP.",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
