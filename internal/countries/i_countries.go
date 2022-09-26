package countries

// country details with name, code and phone number
type CountryDetails struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	PhonePrefix string `json:"prefix"`
}

// country details constructor
func NewCountryDetails(Code, Name, PhonePrefix string) CountryDetails {
	return CountryDetails{
		Code:        Code,
		Name:        Name,
		PhonePrefix: PhonePrefix,
	}
}

type ICountriesRepository interface {
	// a method to download data about countries and phone codes.
	CacheRemoteCountries() error

	// Get cached countries
	GetCountries() ([]byte, error)
}
