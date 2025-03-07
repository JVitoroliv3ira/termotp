package account

import (
	"errors"
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/storage"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var renameOldName, renameNewName string

func RenameAccount(oldName, newName, password string) error {
	if len(oldName) < 3 {
		return errors.New("informe corretamente o nome da conta que deseja renomear (mínimo 3 caracteres)")
	}

	if err := utils.ValidateServiceName(newName); err != nil {
		return err
	}

	if err := utils.ValidatePassword(password); err != nil {
		return err
	}

	data, err := storage.LoadEncrypted(password)
	if err != nil {
		return err
	}

	account, err := data.GetAccount(oldName)
	if err != nil {
		return err
	}

	account.Name = newName

	if err := data.AddAccount(*account); err != nil {
		return err
	}

	if err := data.DeleteAccount(oldName); err != nil {
		return err
	}

	if err := storage.SaveEncrypted(data, password); err != nil {
		return err
	}

	fmt.Printf("\nConta '%s' renomeada com sucesso para '%s'!\n", oldName, newName)
	return nil
}

var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Renomeia uma conta TOTP",
	Long:  "Renomeia uma conta TOTP no armazenamento, alterando seu identificador.",
	Run: func(cmd *cobra.Command, args []string) {
		password, err := utils.PromptPassword()
		utils.HandleError(err)

		utils.HandleError(RenameAccount(renameOldName, renameNewName, password))
	},
}

func init() {
	AccountCmd.AddCommand(renameCmd)
	renameCmd.Flags().StringVarP(&renameOldName, "old-name", "o", "", "Nome atual da conta (ex: gitlab)")
	renameCmd.Flags().StringVarP(&renameNewName, "new-name", "n", "", "Novo nome para a conta (ex: gitlab-personal)")
}
