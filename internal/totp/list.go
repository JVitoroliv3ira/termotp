package totp

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/JVitoroliv3ira/termotp/internal/models"
)

func ShowTOTPList(storageData models.StorageData) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	for {
		clearScreen()
		fmt.Println("Códigos TOTP:")

		now := time.Now()
		secondsRemaining := 30 - (now.Unix() % 30)

		for _, account := range storageData.Accounts {
			code, err := GenerateTOTP(account.Secret)
			if err != nil {
				fmt.Printf("Erro ao gerar código para %s\n", account.Name)
				continue
			}
			fmt.Printf("[%s] %s (Expira em %d segundos)\n", account.Name, code, secondsRemaining)
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
