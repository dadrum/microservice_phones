package apilayercountries

import "errors"

// --------------------------------------------------------------------------------------
// Get cached countries
func (a *ApiLayerCountries) GetCountries() ([]byte, error) {

	if a.cache == nil {
		return nil, errors.New("there is no countries in cache")
	}

	return a.cache, nil
}
