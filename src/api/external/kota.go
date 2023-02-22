package external

import (
	"encoding/json"
	"net/http"
	"shop-api/src/api/response"
)

func GetKotaByID(id string) (*response.Kota, error) {
	endpoint := "https://emsifa.github.io/api-wilayah-indonesia/api/regency/" + id + ".json"
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	body := &response.Kota{}
	json.NewDecoder(resp.Body).Decode(body)
	return body, nil
}
