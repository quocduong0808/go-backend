package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
)

func GenerateRandomOTP() (string, error) {
	const otpLength = 6
	const digits = "0123456789"
	otp := make([]byte, otpLength)

	for i := 0; i < otpLength; i++ {
		// Generate a secure random index in the range of digits
		num, err := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random OTP: %v", err)
		}
		otp[i] = digits[num.Int64()]
	}

	return string(otp), nil
}

func Hash(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return hex.EncodeToString(bs)
}
