package version

import (
	"github.com/spf13/cobra"
)

var applyVersion string

var applyCmd = &cobra.Command{
	Use:   "apply",
	Short: "Atualize a versão do TermOTP",
	Long:  "Atualize o TermOTP para uma versão específica disponível no GitHub Releases.",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	VersionCmd.AddCommand(applyCmd)
	applyCmd.Flags().StringVarP(&applyVersion, "version", "v", "", "Informe a nova versão a ser aplicada")
}
