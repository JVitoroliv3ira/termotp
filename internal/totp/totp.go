package totp

import (
	"time"

	"github.com/pquerna/otp/totp"
)

var generateTOTPFunc = totp.GenerateCode

func GenerateTOTP(secret string) (string, int, error) {
	code, err := generateTOTPFunc(secret, time.Now())
	if err != nil {
		return "", 0, err
	}

	remainingTime := int(30 - (time.Now().Unix() % 30))

	return code, remainingTime, nil
}
