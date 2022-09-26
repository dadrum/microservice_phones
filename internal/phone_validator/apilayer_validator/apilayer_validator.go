package apilayervalidator

/*
###  apilayer.com ###
Free Plan
$ 0.00 / mo
20 Requests / Daily
100 Requests / Monthly
Free for Lifetime
No Credit Card Required
*/

import (
	"micro_service_phone/internal/phone_validator"
	cacherepo "micro_service_phone/pkg/cache_repo"
	"micro_service_phone/pkg/file_logger"
)

const (
	// in release https://api.apilayer.com/number_verification/validate?number=%v
	hostUrl = "http://localhost/?check_valid=%v"

	phoneCheckCacheKey = "phone_check_"
	minimalPhoneLength = 7
)

type ApiLayerPhoneValidator struct {
	apiKey string
	logger *file_logger.FileLogger
	cache  *cacherepo.ICacheRepository
}

// --------------------------------------------------------------------------------------
func NewApiLayerPhoneValidator(
	apiKey string,
	logger *file_logger.FileLogger,
	cache *cacherepo.ICacheRepository,
) phonevalidator.IPhoneValidator {
	return &ApiLayerPhoneValidator{apiKey: apiKey, logger: logger, cache: cache}
}
