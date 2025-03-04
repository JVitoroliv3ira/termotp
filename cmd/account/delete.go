package account

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

func DeleteAccount(name, password string) error {
	if err := utils.ValidateServiceName(name); err != nil {
		return err
	}
	if err := utils.ValidatePassword(password); err != nil {
		return err
	}

	accounts, err := storage.LoadEncrypted(password)
	if err != nil {
		return err
	}

	if err := accounts.DeleteAccount(name); err != nil {
		return err
	}

	if err := storage.SaveEncrypted(accounts, password); err != nil {
		return err
	}

	fmt.Printf("\nConta '%s' removida com sucesso!\n", name)
	return nil
}

var deleteName string

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Exclui uma conta TOTP",
	Long:  "Remove uma conta TOTP do armazenamento.",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := utils.PromptPassword()
		utils.HandleError(err)

		utils.HandleError(DeleteAccount(deleteName, password))
	},
}

func init() {
	AccountCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringVarP(&deleteName, "name", "n", "", "Nome da conta a ser removida (ex: gitlab)")
}
