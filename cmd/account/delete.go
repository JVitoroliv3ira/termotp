package account

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var deleteName string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Exclui uma conta TOTP",
	Long:  "Remove uma conta TOTP do armazenamento.",
	Run: func(cmd *cobra.Command, args []string) {
		utils.HandleError(utils.ValidateServiceName(deleteName))

		password, err := utils.PromptPassword()
		utils.HandleError(err)
		utils.HandleError(utils.ValidatePassword(password))

		storage.DeleteAccount(deleteName, password)
		fmt.Printf("\nConta '%s' removida com sucesso!\n", deleteName)
	},
}

func init() {
	AccountCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&deleteName, "name", "n", "", "Nome da conta a ser removida (ex: gitlab)")
}
