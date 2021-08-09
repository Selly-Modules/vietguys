package vietguys

import (
	"math/rand"
	"time"
)

// Constant ...
const (
	otpCharacters = "123456789"
	otpLength     = 6
)

// randomIntBetweenRange ...
func randomIntBetweenRange(min int, max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max-min) + min
}

// randomOTPCode ...
func randomOTPCode() string {
	var code string
	var length = len(otpCharacters)
	for i := 0; i < otpLength; i++ {
		// Random character index
		randIndex := randomIntBetweenRange(0, length)
		code += string(otpCharacters[randIndex])
	}
	return code
}
