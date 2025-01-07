package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"cryptogo.com/crypto/currency"
)

const apiUrl = "https://cex.io/api/ticker/%s/EUR"

func GetRate(cur string) (*currency.Rate, error) {
	var response CEXResponse
	res, err := http.Get(fmt.Sprintf(apiUrl, strings.ToUpper(cur)))

	if err != nil {
		return nil, err
	}
	if res.StatusCode == http.StatusOK {
		bodyBytes, err := io.ReadAll(res.Body)

		if err != nil {
			return nil, err
		}

		err = json.Unmarshal(bodyBytes, &response)

		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unsuccesful call to the API, status code: %v", res.StatusCode )
	}

	lastPrice, err := strconv.ParseFloat(response.Last, 64)
	if err != nil {
		return nil, err
	}
	rate := currency.Rate{Currency: cur, Price: lastPrice}

	return &rate, nil
}