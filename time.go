package vietguys

import "time"

//
// NOTE: due to unique timezone in server's code, all using time will be convert to HCM timezone (UTC +7)
// All functions generate time, must be call util functions here
// WARNING: don't accept call time.Now() directly
//

const timezoneHCM = "Asia/Ho_Chi_Minh"

// getHCMLocation ...
func getHCMLocation() *time.Location {
	l, _ := time.LoadLocation(timezoneHCM)
	return l
}

// Now ...
func now() time.Time {
	return time.Now().In(getHCMLocation())
}

// TimeBeforeNowInMin ...
func timeBeforeNowInMin(min int) time.Time {
	return time.Now().Add(time.Minute * time.Duration(min) * -1).In(getHCMLocation())
}

// StartOfDate ...
func startOfDate(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, getHCMLocation())
}
