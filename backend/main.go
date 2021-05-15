package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/Fredins/stonks/api/fund"
	_ "github.com/go-sql-driver/mysql"
)

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

func maria() {
	// Create the database handle, confirm driver is present
	db, _ := sql.Open("mysql", "fm:marmar@/stonks")
	defer db.Close()

	// Connect and check the server version
	var version string
	db.QueryRow("SELECT VERSION()").Scan(&version)
	fmt.Println("Connected to:", version)
}

func main() {
	maria()
	// fundList := fund.QuerryFundList()
	// structPrint(fundList)
	// fundListJson := fund.QuerryFundListJson()
	// jsonPrint(fundListJson)
	// fundDetails := fund.QuerryFundDetails(fundList.FundListViews[0].OrderbookId)
	// structPrint(fundDetails)
	// fundDetailsJson := fund.QuerryFundDetailsJson(fundList.FundListViews[0].OrderbookId)
	// jsonPrint(fundDetailsJson)

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("in handler")
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")

		fmt.Fprint(w, string(fund.QuerryFundListJson()))

	})

	http.ListenAndServe(":8080", nil)
}
