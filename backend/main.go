package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

func getRequest() []byte {
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/heatran")
	if err != nil {
		fmt.Println(err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return b
}

func jsonPrint(uglyJson []byte) {
	var prettyJson bytes.Buffer
	json.Indent(&prettyJson, uglyJson, "", "   ")
	prettyJson.WriteTo(os.Stdout)
}
func structPrint(v interface{}) {
	prettyJson, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(prettyJson))

}

func main() {
	// fundList := fund.QuerryFundList()
	// structPrint(fundList)
	// fundListJson := fund.QuerryFundListJson()
	// jsonPrint(fundListJson)
	// fundDetails := fund.QuerryFundDetails(fundList.FundListViews[0].OrderbookId)
	// structPrint(fundDetails)
	// fundDetailsJson := fund.QuerryFundDetailsJson(fundList.FundListViews[0].OrderbookId)
	// jsonPrint(fundDetailsJson)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")

	})

	http.ListenAndServe(":8080", nil)
}
