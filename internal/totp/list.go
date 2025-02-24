package totp

import (
	"fmt"
	"os"
	"os/signal"
	"sort"
	"time"

	"github.com/JVitoroliv3ira/termotp/internal/models"
	"github.com/fatih/color"
)

func ShowTOTPList(storageData models.StorageData) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	accountNames := make([]string, 0, len(storageData.Accounts))
	for name := range storageData.Accounts {
		accountNames = append(accountNames, name)
	}
	sort.Strings(accountNames)

	boldText := color.New(color.Bold)
	greenText := color.New(color.FgGreen)
	yellowText := color.New(color.FgYellow)
	redText := color.New(color.FgRed)
	cyanText := color.New(color.FgCyan)

	for {
		clearScreen()
		now := time.Now()
		secondsRemaining := 30 - (now.Unix() % 30)

		fmt.Println("\nCódigos TOTP:")

		for _, name := range accountNames {
			account := storageData.Accounts[name]
			code, err := GenerateTOTP(account.Secret)
			if err != nil {
				fmt.Printf("Erro ao gerar código para %s\n", account.Name)
				continue
			}

			timerColor := greenText
			if secondsRemaining <= 10 {
				timerColor = yellowText
			}
			if secondsRemaining <= 5 {
				timerColor = redText
			}

			timeText := fmt.Sprintf("(Expira em %d %s)", secondsRemaining, pluralize("segundo", "segundos", secondsRemaining))

			fmt.Printf("[%s] %s %s\n",
				cyanText.Sprint(name),
				boldText.Sprint(code),
				timerColor.Sprint(timeText),
			)
		}

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

func pluralize(singular, plural string, value int64) string {
	if value == 1 {
		return singular
	}
	return plural
}
