package vietguys

import "strings"

// VerifyOTP verify otp code is right or not
func (s Service) VerifyOTP(phone, otpCode string) bool {
	// Just remove char "+" if existed
	strings.Replace(phone, "+", "", 1)

	return s.checkOTP(phone, otpCode)
}
