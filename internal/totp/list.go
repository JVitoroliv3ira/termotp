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

var cacheTOTP = make(map[string]struct {
	Code           string
	ExpirationTime int64
})

func ShowTOTPList(storageData models.StorageData, sortOption string) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	if sortOption != "created" {
		sortOption = "name"
	}

	fmt.Print("\033[H\033[J")
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	for {
		fmt.Print("\033[H")
		fmt.Println("\nCÓDIGOS TOTP")

		if len(storageData.Accounts) == 0 {
			fmt.Println("\nNenhuma conta cadastrada. Adicione uma conta para gerar códigos TOTP.")
			return
		}

		accountNames := SortAccounts(storageData, sortOption)
		RenderTOTPTable(storageData, accountNames, sortOption)

		select {
		case <-stop:
			fmt.Println("\nEncerrando...")
			return
		case <-time.After(1 * time.Second):
		}
	}
}

func SortAccounts(storageData models.StorageData, sortOption string) []string {
	accountNames := make([]string, 0, len(storageData.Accounts))
	for name := range storageData.Accounts {
		accountNames = append(accountNames, name)
	}

	sort.Strings(accountNames)

	return accountNames
}

func RenderTOTPTable(storageData models.StorageData, accountNames []string, sortOption string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"CONTA", "CÓDIGO", "EXPIRA EM"})
	table.SetBorder(true)
	table.SetCenterSeparator("|")
	table.SetColumnSeparator("|")
	table.SetRowSeparator("-")
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetAutoWrapText(false)

	table.SetColMinWidth(0, 15)
	table.SetColMinWidth(1, 10)
	table.SetColMinWidth(2, 20)

	bold := color.New(color.Bold)
	green := color.New(color.FgGreen)
	yellow := color.New(color.FgYellow)
	red := color.New(color.FgRed)

	for _, name := range accountNames {
		account := storageData.Accounts[name]
		code, secondsRemaining := GetCachedTOTP(account.Secret, name)

		timeText := fmt.Sprintf("%d %s", secondsRemaining, utils.Pluralize("segundo", "segundos", secondsRemaining))

		if secondsRemaining <= 5 {
			timeText = red.Sprint(timeText)
		} else if secondsRemaining <= 10 {
			timeText = yellow.Sprint(timeText)
		} else {
			timeText = green.Sprint(timeText)
		}

		table.Append([]string{name, bold.Sprint(code), timeText})
	}

	table.Render()
	fmt.Printf("\nOrdenação atual: %s | Atualizando... Pressione Ctrl+C para sair.\n", sortOption)
}

func GetCachedTOTP(secret, name string) (string, int) {
	now := time.Now().Unix()

	if entry, exists := cacheTOTP[name]; exists && entry.ExpirationTime > now {
		return entry.Code, int(entry.ExpirationTime - now)
	}

	code, secondsRemaining, err := GenerateTOTP(secret)
	if err != nil {
		return "Erro", 30
	}

	expiration := now + int64(secondsRemaining)

	cacheTOTP[name] = struct {
		Code           string
		ExpirationTime int64
	}{
		Code:           code,
		ExpirationTime: expiration,
	}

	return code, secondsRemaining
}
