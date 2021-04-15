package fund

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/Fredins/stonks/key"
)

func fundListPostRequest() []byte {
	content := fundListRequestModel{
		StartIndex:                     0,
		IndexFund:                      false,
		SustainabilityProfile:          false,
		LowCo2:                         false,
		NoFossilFuelInvolvement:        false,
		RegionFilter:                   []string{},
		CountryFilter:                  []string{},
		AlignmentFilter:                []string{},
		IndustryFilter:                 []string{},
		FundTypeFilter:                 []string{},
		InterestTypeFilter:             []string{},
		SortField:                      "developmentFiveYears",
		SortDirection:                  "DESCENDING",
		Name:                           "",
		RecommendedHoldingPeriodFilter: []string{},
		CompanyFilter:                  []string{},
		ProductInvolvementsFilter:      []string{},
		RatingFilter:                   []string{},
		SustainabilityRatingFilter:     []string{},
		EnvironmentalRatingFilter:      []string{},
		SocialRatingFilter:             []string{},
		GovernanceRatingFilter:         []string{}}

	requestJson, err := json.Marshal(content)
	if err != nil {
		fmt.Println(err)
	}
	res, err := http.Post(key.GetFundListUrl(), "application/json", bytes.NewBuffer(requestJson))
	if err != nil {
		fmt.Println(err)
	}
	responseJson, err := io.ReadAll(res.Body)
	return responseJson
}

func unmarshalFundListJson(responseJson []byte) fundListResponseModel {
	var fundlist fundListResponseModel
	json.Unmarshal(responseJson, &fundlist)
	return fundlist
}

func QuerryFundListJson() []byte {
	return fundListPostRequest()
}

func QuerryFundList() fundListResponseModel {
	responseJson := fundListPostRequest()
	fundList := unmarshalFundListJson(responseJson)
	return fundList
}
