package totp

import (
	"time"

	"github.com/pquerna/otp/totp"
)

func GenerateTOTP(secret string) (string, int, error) {
	code, err := totp.GenerateCode(secret, time.Now())
	if err != nil {
		return "", 0, err
	}

	remainingTime := int(30 - (time.Now().Unix() % 30))

	return code, remainingTime, nil
}
