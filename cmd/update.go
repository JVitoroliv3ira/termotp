package cmd

import "github.com/spf13/cobra"

var updateVersion string

var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "Atualize o TermOTP",
	Long:  "Verifique e instale a versão mais recente do TermOTP disponível ou escolha uma versão específica para instalar.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	rootCmd.AddCommand(updateCmd)
	updateCmd.Flags().StringVarP(&updateVersion, "version", "v", "", "Especifique a versão do TermOTP para atualizar (ex: v1.0.0)")
}
