package twiliosendsms

import (
	sendsms "micro_service_phone/internal/send_sms"
	cacherepo "micro_service_phone/pkg/cache_repo"
	"micro_service_phone/pkg/file_logger"
)

const (
	minimalPhoneLength = 7
)

type TwilioSendSms struct {
	accountSid string
	authToken  string
	fromPhone  string
	logger     *file_logger.FileLogger
	cache      *cacherepo.ICacheRepository
}

// --------------------------------------------------------------------------------------
func NewTwilioSendSms(
	accountSid string,
	authToken string,
	fromPhone string,
	logger *file_logger.FileLogger,
	cache *cacherepo.ICacheRepository,
) sendsms.ISendSms {
	return &TwilioSendSms{
		accountSid: accountSid,
		authToken:  authToken,
		fromPhone:  fromPhone,
		logger:     logger,
		cache:      cache,
	}
}
