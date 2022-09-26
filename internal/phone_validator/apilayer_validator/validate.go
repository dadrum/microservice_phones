package apilayervalidator

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"regexp"
	"strings"
)

// --------------------------------------------------------------------------------------
// request phone validating by remote server
func requestRemotePhoneValidate(url, apiKey, digitsOfPhone string) ([]byte, error) {
	// prepare url
	targetUrl := fmt.Sprintf(url, digitsOfPhone)

	// prepare http client and request
	client := &http.Client{}
	req, err := http.NewRequest("GET", targetUrl, nil)
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
/* 	parse json and format DTO->Model */
func parseReceivedData(body []byte) ([]byte, error) {
	// parse json data to `result` variable
	var dto ValidateResultDto
	if err := json.Unmarshal(body, &dto); err != nil {
		return nil, err
	}

	// DTO->Model
	model := MapValidateResultDtoToModel(dto)

	// marshal struct to json
	return json.Marshal(model)
}

// --------------------------------------------------------------------------------------
// return only digits from sting
func filterDigits(phone string) string {
	re := regexp.MustCompile("[0-9]+")
	stringsSlice := re.FindAllString(phone, -1)
	return strings.Join(stringsSlice, "")
}

// --------------------------------------------------------------------------------------
// a method to download data of countries and phone codes.
func (a *ApiLayerPhoneValidator) Validate(phone string) ([]byte, error) {

	// prepare received phone number
	phoneDigits := filterDigits(phone)
	if len(phoneDigits) < minimalPhoneLength {
		return nil, errors.New("phone number is too short")
	}

	// check cached data for this phone
	cachePhoneKey := []byte(phoneCheckCacheKey + phoneDigits)
	cacheResult, err := (*a.cache).Get(cachePhoneKey)
	if err == nil && cacheResult != nil {
		// return cached value
		return cacheResult, nil
	}

	//request remote validate
	testResultBytes, err := requestRemotePhoneValidate(hostUrl, a.apiKey, phoneDigits)
	if err != nil {
		return nil, err
	}

	answerBytes, err := parseReceivedData(testResultBytes)
	if err != nil {
		return nil, err
	}

	// cache result
	err = (*a.cache).Set(cachePhoneKey, answerBytes)

	return answerBytes, err

}
