package twiliosendsms

import (
	"errors"
	"regexp"
	"strings"

	"github.com/sfreiberg/gotwilio"
)

// --------------------------------------------------------------------------------------
// return only digits from sting
func filterDigits(phone string) string {
	re := regexp.MustCompile("[0-9]+")
	stringsSlice := re.FindAllString(phone, -1)
	return strings.Join(stringsSlice, "")
}

// --------------------------------------------------------------------------------------
// a method to send sms
func (s *TwilioSendSms) Send(phone, message string) ([]byte, error) {

	// prepare received phone number
	phoneDigits := filterDigits(phone)
	if len(phoneDigits) < minimalPhoneLength {
		return nil, errors.New("phone number is too short")
	}

	twilio := gotwilio.NewTwilioClient(s.accountSid, s.authToken)

	smsResponse, exception, err := twilio.SendSMS(s.fromPhone, phone, message, "", "")
	if err != nil {
		return nil, err
	}
	if exception != nil {
		s.logger.Errorln("Twilio exception: " + exception.Error())
		return nil, errors.New("Twilio: " + exception.Message)
	}
	return []byte(smsResponse.Body), nil
}
