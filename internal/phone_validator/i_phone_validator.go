package phonevalidator

type ValidateResult struct {
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

type IPhoneValidator interface {
	// a method to validate phone.
	Validate(phone string) ([]byte, error)
}
