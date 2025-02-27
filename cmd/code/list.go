package code

import (
	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/totp"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Lista todas as contas e seus códigos TOTP",
	Long:  "Exibe todas as contas cadastradas, atualizando os códigos TOTP em tempo real.",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := utils.PromptPassword()
		utils.HandleError(err)
		utils.HandleError(utils.ValidatePassword(password))

		accounts, err := storage.LoadAccounts(password)
		utils.HandleError(err)

		totp.ShowTOTPList(accounts)
	},
}

func init() {
	CodeCmd.AddCommand(listCmd)
}
