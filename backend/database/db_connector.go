package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/Fredins/stonks/api/fund"
)

var (
	db *sql.DB
)

func Ping() {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func Connect() {
	// Create the database handle, confirm driver is present
	var err error
	db, err = sql.Open("mysql", "fm:marmar@/stonks")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	d := fund.QuerryFundDetails(1158131)

	res, err := db.Exec(`INSERT INTO details (
		OrderbookId, 
		Name, 
		MinimumThresholdAdditionalBuy, 
		MinimumThresholdBuy, 
		ProductFee, 
		ProspectusLink, 
		BuyStopDateTime, 
		SellStopDateTime, 
		Nav, 
		NavDate, 
		BuyCommission, 
		HasCurrencyExchangeFee, 
		Rating, 
		Risk, 
		RiskNumber, 
		Owners, 
		Description, 
		NavChangeThreeMonths, 
		NavChangeOneYears, 
		NavChangeThreeYears, 
		NavChangeFiveYears, 
		NavChangeTenYears, 
		ManagementFee
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		d.OrderbookId, d.Name, d.MinimumThresholdAdditionalBuy, d.MinimumThresholdBuy,
		d.ProductFee, d.ProspectusLink, d.BuyStopDateTime, d.SellStopDateTime, d.Nav, d.NavDate,
		d.BuyCommission, d.HasCurrencyExchangeFee, d.Rating, d.Risk, d.RiskNumber, d.Owners, d.Description,
		d.NavChangeThreeMonths, d.NavChangeOneYears, d.NavChangeThreeYears, d.NavChangeFiveYears, d.NavChangeTenYears, d.ManagementFee)

	if err != nil {
		log.Fatal(err)
	}
	n, _ := res.RowsAffected()
	fmt.Printf("affected %d rows\n", n)

}
