package cmd

import (
	"fmt"

	"github.com/JVitoroliv3ira/termotp/internal/update"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/spf13/cobra"
)

var updateVersion string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualize o TermOTP",
	Long:  "Verifique e instale a versão mais recente do TermOTP disponível ou escolha uma versão específica para instalar.",
	Run: func(cmd *cobra.Command, args []string) {
		u := update.NewUpdater()
		installed, err := u.Update(updateVersion)
		utils.HandleError(err)

		if installed {
			if updateVersion == "" {
				fmt.Println("TermOTP foi atualizado para a versão mais recente com sucesso!")
			} else {
				fmt.Printf("TermOTP foi atualizado para a versão %s com sucesso!\n", updateVersion)
			}
		} else {
			if updateVersion == "" {
				fmt.Println("TermOTP já está na versão mais recente.")
			} else {
				fmt.Printf("A versão %s já estava instalada, portanto nenhuma ação foi realizada.\n", updateVersion)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateVersion, "version", "v", "", "Especifique a versão do TermOTP para atualizar (ex: v1.0.0)")
}
