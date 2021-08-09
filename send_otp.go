package vietguys

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/Selly-Modules/mongodb"
)

// SendOTPResult ...
type SendOTPResult struct {
	Carrier   string
	Error     int
	ErrorCode int
	MsgID     string
	Message   string
	Log       string
}

// SendOTP to phone number
// Phone format: 84935123456
// Content must have "%s", for otp code substitution
func (s Service) SendOTP(phone, ip, contentFormat string) error {
	// Generate otp code
	otpCode := randomOTPCode()

	// Format content
	content := fmt.Sprintf(contentFormat, otpCode)

	// Just remove char "+" if existed
	strings.Replace(phone, "+", "", 1)

	// Check that phone or ip is not over quota
	canSend := s.checkCanSend(phone, ip)
	if canSend {
		return errors.New("ip or phone has reached over limited quota per day")
	}

	// Create payload
	params := url.Values{}
	params.Add("u", s.User)
	params.Add("pwd", s.Password)
	params.Add("from", s.From)
	params.Add("json", "1")
	params.Add("phone", phone)
	params.Add("sms", content)
	payload := strings.NewReader(params.Encode())

	// Create request
	client := s.Client
	req, err := http.NewRequest(http.MethodPost, s.Endpoint, payload)
	if err != nil {
		return err
	}

	// Add necessary headers
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	// Call
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	// Make sure close body
	defer res.Body.Close()

	// Ready body
	body, err := ioutil.ReadAll(res.Body)
	rawResult := string(body)
	fmt.Println(rawResult)
	if err != nil {
		fmt.Println("error : ", err.Error())
		return err
	}

	var result SendOTPResult
	if err = json.Unmarshal(body, &result); err != nil {
		return err
	}

	go func() {
		// Save log to db
		log := Log{
			ID:          mongodb.NewObjectID(),
			Carrier:     result.Carrier,
			Type:        SMSTypeOTP,
			Code:        otpCode,
			IsCodeValid: true,
			PhoneNumber: phone,
			IP:          ip,
			Content:     content,
			CreatedAt:   now(),
			Success:     result.Error == 0,
			Result:      rawResult,
		}
		s.saveLog(log)
	}()

	return nil
}
