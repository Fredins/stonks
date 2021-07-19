package fund

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Fredins/stonks/key"
)

func fundDetailsGetRequest(orderbookId int) []byte {
	res, err := http.Get(key.GetFundDetailsUrl(fmt.Sprint(orderbookId)))
	if err != nil {
		println(err)
	}
	responseJson, err := io.ReadAll(res.Body)
	if err != nil {
		println(err)
	}
	return responseJson
}

func unmarshalFundDetailsJson(responseJson []byte) FundDetails {
	var fundDetails FundDetails
	err := json.Unmarshal(responseJson, &fundDetails)
	if err != nil {
		fmt.Println(err)
	}
	return fundDetails
}

func QuerryFundDetailsJson(orderbookId int) []byte {
	return fundDetailsGetRequest(orderbookId)
}

func QuerryFundDetails(orderbookId int) FundDetails {
	responseJson := fundDetailsGetRequest(orderbookId)
	fundDetails := unmarshalFundDetailsJson(responseJson)
	return fundDetails
}
