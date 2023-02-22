package external

import (
	"encoding/json"
	"net/http"
	"shop-api/src/api/response"
)

func GetProvinsiByID(id string) (*response.Provinsi, error) {
	endpoint := "https://emsifa.github.io/api-wilayah-indonesia/api/province/" + id + ".json"
	resp, err := http.Get(endpoint)
	if err != nil {
		return nil, err
	}

	body := &response.Provinsi{}
	json.NewDecoder(resp.Body).Decode(body)
	return body, nil
}
