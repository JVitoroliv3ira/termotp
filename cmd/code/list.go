package code

import (
	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/totp"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste todas as contas e códigos TOTP",
	Long:  "Exiba todas as contas cadastradas com seus códigos TOTP atualizados em tempo real.",
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
