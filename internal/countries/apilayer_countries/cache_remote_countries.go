package apilayercountries

import (
	"encoding/json"
	"io"
	"micro_service_phone/internal/countries"
	"net/http"
)

// --------------------------------------------------------------------------------------
// request data from remote server
func requestRemoteCountries(url, apiKey string) ([]byte, error) {
	// prepare http client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Set("apikey", apiKey)
	if err != nil {
		return nil, err
	}

	// do prepared request
	res, err := client.Do(req)

	// defer close body stream
	if res.Body != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return nil, err
	}

	// read all data from body stream
	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

// --------------------------------------------------------------------------------------
func parseReceivedData(body []byte) ([]byte, error) {
	// parse json data to `result` variable
	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	// make slice with expected capacity
	countriesSlice := make([]countries.CountryDetails, 0, len(result))

	// iterate received data
	for key, value := range result {
		detailsField := value.(map[string]interface{})
		country := countries.NewCountryDetails(
			key,
			detailsField["country_name"].(string),
			detailsField["dialling_code"].(string),
		)
		countriesSlice = append(countriesSlice, country)
	}

	return json.Marshal(countriesSlice)
}

// --------------------------------------------------------------------------------------
// a method to download data of countries and phone codes.
func (a *ApiLayerCountries) CacheRemoteCountries() error {
	// request data
	a.logger.Debugln("Requesting remote countries")
	body, err := requestRemoteCountries(hostUrl, a.apiKey)
	if err != nil {
		a.logger.Debugln(err.Error())
		return err
	}

	a.logger.Debugln("Parsing countries")
	reformat, _ := parseReceivedData(body)

	a.logger.Debugln("Cache countries")
	a.cache = reformat

	return nil

}
