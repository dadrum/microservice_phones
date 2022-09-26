package sendsms

type ISendSms interface {
	// a method to send sms
	Send(phone, message string) ([]byte, error)
}
