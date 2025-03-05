package code

import (
	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/totp"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var listSort string

func ListTOTP(password string) error {
	if err := utils.ValidatePassword(password); err != nil {
		return err
	}

	accounts, err := storage.LoadEncrypted(password)
	if err != nil {
		return err
	}

	totp.ShowTOTPList(*accounts, listSort)
	return nil
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "Liste todas as contas e códigos TOTP",
	Long:  "Exiba todas as contas cadastradas com seus códigos TOTP atualizados em tempo real.",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := utils.PromptPassword()
		utils.HandleError(err)

		utils.HandleError(ListTOTP(password))
	},
}

func init() {
	CodeCmd.AddCommand(listCmd)
	listCmd.Flags().StringVarP(&listSort, "sort", "s", "", "Define a ordem de exibição: 'name' para ordenar por nome ou 'created' para ordenar por data de criação.")
}
