package code

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/totp"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var generateName string

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Gere um código TOTP",
	Long:  "Gere um código TOTP válido para autenticação em dois fatores (2FA).",
	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleError(utils.ValidateServiceName(generateName))

		password, err := utils.PromptPassword()
		utils.HandleError(err)
		utils.HandleError(utils.ValidatePassword(password))

		accounts, err := storage.LoadEncrypted(password)
		utils.HandleError(err)
		account, err := accounts.GetAccount(generateName)
		utils.HandleError(err)
		code, secondsRemaining, err := totp.GenerateTOTP(account.Secret)
		utils.HandleError(err)

		fmt.Printf("Código TOTP para a conta [%s]: %s\n", account.Name, code)
		fmt.Printf("Este código é válido por %d %s.\n", secondsRemaining, utils.Pluralize("segundo", "segundos", secondsRemaining))
	},
}

func init() {
	CodeCmd.AddCommand(generateCmd)
	generateCmd.Flags().StringVarP(&generateName, "name", "n", "", "Nome da conta (ex: gitlab)")
}
