package apilayervalidator

import phonevalidator "micro_service_phone/internal/phone_validator"

type ValidateResultDto struct {
	Carrier             string  `json:"carrier"`
	CountryCode         string  `json:"country_code"`
	CountryName         string  `json:"country_name"`
	CountryPrefix       string  `json:"country_prefix"`
	InternationalFormat string  `json:"international_format"`
	LineType            *string `json:"line_type"`
	LocalFormat         string  `json:"local_format"`
	Location            string  `json:"location"`
	Number              string  `json:"number"`
	Valid               bool    `json:"valid"`
}

func MapValidateResultDtoToModel(dto ValidateResultDto) phonevalidator.ValidateResult {
	return phonevalidator.ValidateResult{
		Carrier:             dto.Carrier,
		CountryCode:         dto.CountryCode,
		CountryName:         dto.CountryName,
		CountryPrefix:       dto.CountryPrefix,
		InternationalFormat: dto.InternationalFormat,
		LineType:            dto.LineType,
		LocalFormat:         dto.LocalFormat,
		Location:            dto.Location,
		Number:              dto.Number,
		Valid:               dto.Valid,
	}
}
