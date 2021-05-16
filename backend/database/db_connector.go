package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Fredins/stonks/api/fund"
)

var (
	db *sql.DB
	ti string
)

func Ping() {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func Connect() {
	// Create the database handle, confirm driver is present
	ti = time.Now().Format(time.UnixDate)
	var err error
	db, err = sql.Open("mysql", "fm:marmar@/stonks")
	if err != nil {
		log.Fatal(err)
	}
}

func InsertList(l fund.FundListViews) {
	taglist, err := json.Marshal(l.Taglist)
	if err != nil {
		log.Fatal(err)
	}
	productInvolvements, err := json.Marshal(l.ProductInvolvements)
	if err != nil {
		log.Fatal(err)
	}
	productInvolvementViews, err := json.Marshal(l.ProductInvolvementViews)
	if err != nil {
		log.Fatal(err)
	}

	res, err := db.Exec(`INSERT INTO List (
		Isin, 
		Name, 
		Rating, 
		Risk, 
        DevelopmentOneDay, 
        DevelopmentOneWeek, 
        DevelopmentOneMonth, 
        DevelopmentThreeMonths, 
        DevelopmentOneYears, 
        DevelopmentThisYear, 
        DevelopmentThreeYears, 
        DevelopmentFiveYears, 
        DevelopmentTenYears, 
        LowCarbon, 
        SharpeRatio, 
        StandardDeviation, 
        Capital, 
        FossilFuelInvolvement, 
        CarbonRiskScore, 
        PrimaryBenchmark, 
        RecommendedHoldingPeriod, 
        EsqScore, 
        EnvironmentalScore, 
        SocialScore, 
        GovernanceScore, 
        ManagementFee, 
        TotalFee, 
        TransactionFee, 
        OngoingFee, 
        OtherFee, 
        MinimumBuy, 
        MinimumBuyMonthlySaving, 
        HasCurrencyExchangeFee, 
        NrOfOwners, 
        OrderbookId, 
        Taglist, 
        Category, 
        IndexFund, 
        StartDate, 
        SuperloanOrderbook, 
        FundType, 
        CompanyName, 
        MatchesSustainabilityRating, 
		SustainabilityLevel,
        SustainabilityRating, 
        SustainabilityRatingCategoryName, 
        ProductInvolvements, 
        ProductInvolvementViews, 
        SustainabilityProfileIcon, 
        CarbonSolutionsInvolvement, 
        AumCoveredCarbon, 
        ThermalCoalInvolvement, 
        OilSandsExtractionInvolvement, 
        ArcticOilAndGasExplorationInvolvement, 
        EnvironmentalRating, 
        SocialRating, 
        GovernanceRating, 
        ProspectusLink, 
        CurrencyCode, 
        RecommendedHoldingPeriodValue, 
        InsertDate 
	    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		l.Isin, l.Name, l.Rating, l.Risk, l.DevelopmentOneDay, l.DevelopmentOneWeek, l.DevelopmentOneMonth, l.DevelopmentThreeMonths, l.DevelopmentOneYear,
		l.DevelopmentThisYear, l.DevelopmentThreeYears, l.DevelopmentFiveYears, l.DevelopmentTenYears, l.LowCarbon, l.SharpeRatio, l.StandardDeviation, l.Capital,
		l.FossilFuelInvolvement, l.CarbonRiskScore, l.PrimaryBenchmark, l.RecommendedHoldingPeriod, l.EsgScore, l.EnvironmentalScore, l.SocialScore, l.GovernanceScore,
		l.ManagementFee, l.TotalFee, l.TransactionFee, l.OngoingFee, l.OtherFee, l.MinimumBuy, l.MinimumBuyMonthlySaving, l.HasCurrencyExchangeFee, l.NrOfOwners,
		l.OrderbookId, taglist, l.Category, l.IndexFund, l.StartDate, l.SuperloanOrderbook, l.FundType, l.CompanyName, l.MatchesSustainabilityProfile, l.SustainabilityLevel,
		l.SustainabilityRating, l.SustainabilityRatingCategoryName, productInvolvements, productInvolvementViews, l.SustainabilityProfileIcon, l.CarbonSolutionsInvolvement, l.AumCoveredCarbon,
		l.ThermalCoalInvolvement, l.OilSandsExtractionInvolvement, l.ArcticOilAndGasExplorationInvolvement, l.EnvironmentalRating, l.SocialRating, l.GovernanceRating,
		l.ProspectusLink, l.CurrencyCode, l.RecommendedHoldingPeriod, ti)
	if err != nil {
		log.Fatal(err)
	}
	n, _ := res.RowsAffected()
	fmt.Printf("affected %d rows\n", n)
}

func InsertDetails(d fund.FundDetailsResponseModel) {
	res, err := db.Exec(`INSERT INTO Details (
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
		ManagementFee,
		InsertDate
		) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		d.OrderbookId, d.Name, d.MinimumThresholdAdditionalBuy, d.MinimumThresholdBuy,
		d.ProductFee, d.ProspectusLink, d.BuyStopDateTime, d.SellStopDateTime, d.Nav, d.NavDate,
		d.BuyCommission, d.HasCurrencyExchangeFee, d.Rating, d.Risk, d.RiskNumber, d.Owners, d.Description,
		d.NavChangeThreeMonths, d.NavChangeOneYears, d.NavChangeThreeYears, d.NavChangeFiveYears, d.NavChangeTenYears, d.ManagementFee, ti)

	if err != nil {
		log.Fatal(err)
	}
	n, _ := res.RowsAffected()
	fmt.Printf("affected %d rows\n", n)
}
