package main

import (
	"fmt"
	"net/http"

	"github.com/Fredins/stonks/api/fund"
	"github.com/Fredins/stonks/database"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	database.Connect()

	http.HandleFunc("/api", func(w http.ResponseWriter, r *http.Request) {
		fmt.Print("in handler")
		w.Header().Add("Access-Control-Allow-Origin", "http://localhost:3000")

		fmt.Fprint(w, string(fund.QuerryFundListJson()))

	})

	http.ListenAndServe(":8080", nil)
}
