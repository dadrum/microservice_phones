package apilayercountries

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
	"micro_service_phone/internal/countries"
	"micro_service_phone/pkg/file_logger"
)

const (
	// in release https://api.apilayer.com/number_verification/countries
	hostUrl = "http://localhost/?countries"
)

type ApiLayerCountries struct {
	apiKey string
	logger *file_logger.FileLogger
	cache  []byte
}

// --------------------------------------------------------------------------------------
func NewApiLayerCountries(
	apiKey string,
	logger *file_logger.FileLogger,
) countries.ICountriesRepository {
	return &ApiLayerCountries{apiKey: apiKey, logger: logger}
}
