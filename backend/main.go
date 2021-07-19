package main

import (
	"fmt"
	"net/http"

	"github.com/Fredins/stonks/api/fund"
	"github.com/Fredins/stonks/database"
	_ "github.com/go-sql-driver/mysql"
)

func handleDb(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	res := database.Data()
	fmt.Fprint(w, string(res))
}

func handleApi(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
	res := fund.QuerryFundListJson()
	fmt.Fprint(w, string(res))
}

func fillDb() {
	database.Connect()
	ll := fund.QuerryFundList().FundListViews
	fmt.Println(len(ll))
	for _, l := range ll {
		d := fund.QuerryFundDetails(l.OrderbookId)
		e := database.CreateEntry(l, d)
		database.Insert(e)

	}
}

// TODO fix struct for database row

func main() {
	fillDb()
	// d := database.Data()
	// fmt.Println("data: " + string(d))
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")
		fmt.Fprint(w, "hello world")
	})
	http.HandleFunc("/api", handleApi)
	http.HandleFunc("/db", handleDb)
	http.ListenAndServe(":8080", nil)
}
