package vietguys

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

const (
	otpValidMinute = 30
)

// Check phone or ip can send sms
func (s Service) checkCanSend(phone, ip string) bool {
	var (
		canSend      = true
		startOfToday = startOfDate(time.Now())
	)

	// Check phone number first
	if len(phone) > 0 && s.PhoneMaxSendPerDay > 0 {
		total, _ := s.DB.CountDocuments(bgCtx, bson.M{
			"phoneNumber": phone,
			"createdAt": bson.M{
				"$gte": startOfToday,
			},
		})
		canSend = total > int64(s.PhoneMaxSendPerDay)
	}

	// Check ip, but only check if can send
	if canSend && len(ip) > 0 && s.IPMaxSendPerDay > 0 {
		total, _ := s.DB.CountDocuments(bgCtx, bson.M{
			"ip": ip,
			"createdAt": bson.M{
				"$gte": startOfToday,
			},
		})
		canSend = total > int64(s.IPMaxSendPerDay)
	}

	return canSend
}

// Check otp right or not
func (s Service) checkOTP(phone, otpCode string) bool {
	var (
		timeAgo = timeBeforeNowInMin(otpValidMinute)
	)

	total, _ := s.DB.CountDocuments(bgCtx, bson.M{
		"phoneNumber": phone,
		"code":        otpCode,
		"isCodeValid": true,
		"createdAt": bson.M{
			"$gte": timeAgo,
		},
	})
	isValid := total > 0

	// If success, set code valid to false
	if isValid {
		s.DB.UpdateOne(bgCtx, bson.M{
			"code": otpCode,
		}, bson.M{
			"$set": bson.M{
				"isCodeValid": false,
			},
		})
	}
	return isValid
}
