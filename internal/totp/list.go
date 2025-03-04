package totp

import (
	"fmt"
	"os"
	"os/signal"
	"sort"
	"time"

	"github.com/JVitoroliv3ira/termotp/internal/models"
	"github.com/JVitoroliv3ira/termotp/internal/utils"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
)

func ShowTOTPList(storageData models.StorageData) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	accountNames := make([]string, 0, len(storageData.Accounts))
	for name := range storageData.Accounts {
		accountNames = append(accountNames, name)
	}
	sort.Strings(accountNames)

	bold := color.New(color.Bold)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed)

	for {
		clearScreen()
		fmt.Println(bold.Sprint("\n CÓDIGOS TOTP"))

		if len(accountNames) == 0 {
			fmt.Println("\nNenhuma conta cadastrada. Adicione uma conta para gerar códigos TOTP.")
			return
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{" CONTA ", " CÓDIGO ", " EXPIRA EM "})
		table.SetBorder(true)
		table.SetCenterSeparator("|")
		table.SetColumnSeparator("|")
		table.SetRowSeparator("-")
		table.SetAlignment(tablewriter.ALIGN_CENTER)

		table.SetColMinWidth(0, 15)
		table.SetColMinWidth(1, 10)
		table.SetColMinWidth(2, 15)
		table.SetAutoWrapText(false)
		table.SetHeaderAlignment(tablewriter.ALIGN_CENTER)
		table.SetColumnAlignment([]int{
			tablewriter.ALIGN_CENTER,
			tablewriter.ALIGN_CENTER,
			tablewriter.ALIGN_CENTER,
		})
		table.SetHeaderLine(true)
		table.SetRowLine(true)
		table.SetAutoFormatHeaders(false)

		for _, name := range accountNames {
			account := storageData.Accounts[name]
			code, secondsRemaining, err := GenerateTOTP(account.Secret)
			if err != nil {
				table.Append([]string{name, "Erro ao gerar", "-"})
				continue
			}

			timeText := fmt.Sprintf("%d %s", secondsRemaining, utils.Pluralize("segundo", "segundos", secondsRemaining))
			if secondsRemaining <= 5 {
				timeText = red.Sprint(timeText)
			} else if secondsRemaining <= 10 {
				timeText = yellow.Sprint(timeText)
			} else {
				timeText = green.Sprint(timeText)
			}

			table.Append([]string{
				fmt.Sprintf(" %-15s ", name),
				fmt.Sprintf(" %s ", bold.Sprint(code)),
				fmt.Sprintf(" %s ", timeText),
			})
		}

		table.Render()

		fmt.Println("\nAtualizando... Pressione Ctrl+C para sair.")

		select {
		case <-stop:
			fmt.Println("\nEncerrando...")
			return
		case <-time.After(1 * time.Second):
		}
	}
}

func clearScreen() {
	fmt.Print("\033[H\033[2J")
}
