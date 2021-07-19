package database

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Fredins/stonks/api/fund"
)

type Entry struct {
	Isin                                  string
	Name                                  string
	Rating                                int
	Risk                                  int
	DevelopmentOneDay                     float32
	DevelopmentOneWeek                    float32
	DevelopmentOneMonth                   float32
	DevelopmentThreeMonths                float32
	DevelopmentOneYear                    float32
	DevelopmentThisYear                   float32
	DevelopmentThreeYears                 float32
	DevelopmentFiveYears                  float32
	DevelopmentTenYears                   float32
	LowCarbon                             float32
	SharpeRatio                           float32
	StandardDeviation                     float32
	Capital                               float64
	FossilFuelInvolvement                 float32
	CarbonRiskScore                       int
	PrimaryBenchmark                      string
	RecommendedHoldingPeriod              string
	EsgScore                              float32
	EnvironmentalScore                    float32
	SocialScore                           float32
	GovernanceScore                       float32
	ManagementFee                         float32
	TotalFee                              float32
	TransactionFee                        float32
	OngoingFee                            float32
	OtherFee                              float32
	MinimumBuy                            float32
	MinimumBuyMonthlySaving               float32
	HasCurrencyExchangeFee                bool
	NrOfOwners                            int
	OrderbookId                           int
	Taglist                               []byte
	Category                              string
	IndexFund                             bool
	StartDate                             string
	SuperloanOrderbook                    bool
	FundType                              string
	CompanyName                           string
	MatchesSustainabilityProfile          bool
	SustainabilityLevel                   string
	SustainabilityRating                  int
	SustainabilityRatingCategoryName      string
	ProductInvolvements                   []byte
	ProductInvolvementViews               []byte
	SustainabilityProfileIcon             string
	CarbonSolutionsInvolvement            float32
	AumCoveredCarbon                      float32
	ThermalCoalInvolvement                float32
	OilSandsExtractionInvolvement         float32
	ArcticOilAndGasExplorationInvolvement float32
	EnvironmentalRating                   int
	SocialRating                          int
	GovernanceRating                      int
	ProspectusLink                        string
	CurrencyCode                          string
	RecommendedHoldingPeriodValue         string
	InsertDate                            string
	Description                           string
	NavChangeThreeMonths                  float32
	NavChangeOneYears                     float32
	NavChangeThreeYears                   float32
	NavChangeFiveYears                    float32
	NavChangeTenYears                     float32
	MinimumThresholdAdditionalBuy         float32
	MinimumThresholdBuy                   float32
	ProductFee                            float32
	BuyStopDateTime                       string
	SellStopDateTime                      string
	Nav                                   float32
	NavDate                               string
	BuyCommission                         float32
}

type List struct {
	Isin                                  string
	Name                                  string
	Rating                                int
	Risk                                  int
	DevelopmentOneDay                     float32
	DevelopmentOneWeek                    float32
	DevelopmentOneMonth                   float32
	DevelopmentThreeMonths                float32
	DevelopmentOneYear                    float32
	DevelopmentThisYear                   float32
	DevelopmentThreeYears                 float32
	DevelopmentFiveYears                  float32
	DevelopmentTenYears                   float32
	LowCarbon                             float32
	SharpeRatio                           float32
	StandardDeviation                     float32
	Capital                               float64
	FossilFuelInvolvement                 float32
	CarbonRiskScore                       int
	PrimaryBenchmark                      string
	RecommendedHoldingPeriod              string
	EsgScore                              float32
	EnvironmentalScore                    float32
	SocialScore                           float32
	GovernanceScore                       float32
	ManagementFee                         float32
	TotalFee                              float32
	TransactionFee                        float32
	OngoingFee                            float32
	OtherFee                              float32
	MinimumBuy                            float32
	MinimumBuyMonthlySaving               float32
	HasCurrencyExchangeFee                bool
	NrOfOwners                            int
	OrderbookId                           int
	Taglist                               []byte
	Category                              string
	IndexFund                             bool
	StartDate                             string
	SuperloanOrderbook                    bool
	FundType                              string
	CompanyName                           string
	MatchesSustainabilityProfile          bool
	SustainabilityLevel                   string
	SustainabilityRating                  int
	SustainabilityRatingCategoryName      string
	ProductInvolvements                   []byte
	ProductInvolvementViews               []byte
	SustainabilityProfileIcon             string
	CarbonSolutionsInvolvement            float32
	AumCoveredCarbon                      float32
	ThermalCoalInvolvement                float32
	OilSandsExtractionInvolvement         float32
	ArcticOilAndGasExplorationInvolvement float32
	EnvironmentalRating                   int
	SocialRating                          int
	GovernanceRating                      int
	ProspectusLink                        string
	CurrencyCode                          string
	RecommendedHoldingPeriodValue         string
	InsertDate                            string
}

var (
	db *sql.DB
)

func Data() []byte {
	var l fund.FundListViews
	r := db.QueryRow("SELECT * FROM List")
	var taglist []byte
	var productInvolvements []byte
	var productInvolvementViews []byte
	var ti string
	r.Scan(
		&l.Isin, &l.Name, &l.Rating, &l.Risk, &l.DevelopmentOneDay, &l.DevelopmentOneWeek, &l.DevelopmentOneMonth, &l.DevelopmentThreeMonths, &l.DevelopmentOneYear,
		&l.DevelopmentThisYear, &l.DevelopmentThreeYears, &l.DevelopmentFiveYears, &l.DevelopmentTenYears, &l.LowCarbon, &l.SharpeRatio, &l.StandardDeviation, &l.Capital,
		&l.FossilFuelInvolvement, &l.CarbonRiskScore, &l.PrimaryBenchmark, &l.RecommendedHoldingPeriod, &l.EsgScore, &l.EnvironmentalScore, &l.SocialScore, &l.GovernanceScore,
		&l.ManagementFee, &l.TotalFee, &l.TransactionFee, &l.OngoingFee, &l.OtherFee, &l.MinimumBuy, &l.MinimumBuyMonthlySaving, &l.HasCurrencyExchangeFee, &l.NrOfOwners,
		&l.OrderbookId, &taglist, &l.Category, &l.IndexFund, &l.StartDate, &l.SuperloanOrderbook, &l.FundType, &l.CompanyName, &l.MatchesSustainabilityProfile, &l.SustainabilityLevel,
		&l.SustainabilityRating, &l.SustainabilityRatingCategoryName, &productInvolvements, &productInvolvementViews, &l.SustainabilityProfileIcon, &l.CarbonSolutionsInvolvement, &l.AumCoveredCarbon,
		&l.ThermalCoalInvolvement, &l.OilSandsExtractionInvolvement, &l.ArcticOilAndGasExplorationInvolvement, &l.EnvironmentalRating, &l.SocialRating, &l.GovernanceRating,
		&l.ProspectusLink, &l.CurrencyCode, &l.RecommendedHoldingPeriod, &ti)
	j, err := json.Marshal(l)
	if err != nil {
		log.Fatal(err)
	}
	return j
}

func Ping() {
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("db connected")
	}
}

func Connect() {
	// Create the database handle, confirm driver is present
	var err error
	db, err = sql.Open("mysql", "fm:marmar@/stonks")
	if err != nil {
		log.Fatal(err)
	}
}

func CreateEntry(l fund.FundListViews, d fund.FundDetails) Entry {
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

	var r List
	b, err := json.Marshal(l)
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(b, &r)

	r.Taglist = taglist
	r.ProductInvolvements = productInvolvements
	r.ProductInvolvementViews = productInvolvementViews
	r.InsertDate = time.Now().Format(time.UnixDate)
	return r
}

func Insert(e Entry) {
	_, err := db.Exec(`INSERT INTO List (
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
        InsertDate,
		Description,
		NavChangeThreeMonths,
		NavChangeOneYears,
		NavChangeThreeYears,
		NavChangeFiveYears,
		NavChangeTenYears,
    	MinimumThresholdAdditionalBuy,
	    MinimumThresholdBuy,
	    ProductFee,
	    BuyStopDateTime,
	    SellStopDateTime,
	    Nav,
	    NavDate,
	    BuyCommission,
	    ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`,
		e.Isin, e.Name, e.Rating, e.Risk, e.DevelopmentOneDay, e.DevelopmentOneWeek, e.DevelopmentOneMonth, e.DevelopmentThreeMonths, e.DevelopmentOneYear,
		e.DevelopmentThisYear, e.DevelopmentThreeYears, e.DevelopmentFiveYears, e.DevelopmentTenYears, e.LowCarbon, e.SharpeRatio, e.StandardDeviation, e.Capital,
		e.FossilFuelInvolvement, e.CarbonRiskScore, e.PrimaryBenchmark, e.RecommendedHoldingPeriod, e.EsgScore, e.EnvironmentalScore, e.SocialScore, e.GovernanceScore,
		e.ManagementFee, e.TotalFee, e.TransactionFee, e.OngoingFee, e.OtherFee, e.MinimumBuy, e.MinimumBuyMonthlySaving, e.HasCurrencyExchangeFee, e.NrOfOwners,
		e.OrderbookId, e.Taglist, e.Category, e.IndexFund, e.StartDate, e.SuperloanOrderbook, e.FundType, e.CompanyName, e.MatchesSustainabilityProfile, e.SustainabilityLevel,
		e.SustainabilityRating, e.SustainabilityRatingCategoryName, e.ProductInvolvements, e.ProductInvolvementViews, e.SustainabilityProfileIcon, e.CarbonSolutionsInvolvement, e.AumCoveredCarbon,
		e.ThermalCoalInvolvement, e.OilSandsExtractionInvolvement, e.ArcticOilAndGasExplorationInvolvement, e.EnvironmentalRating, e.SocialRating, e.GovernanceRating,
		e.ProspectusLink, e.CurrencyCode, e.RecommendedHoldingPeriod, e.InsertDate, e.Description, e.NavChangeThreeMonths, e.NavChangeOneYears, e.NavChangeThreeMonths, e.NavChangeFiveYears,
		e.NavChangeTenYears, e.MinimumThresholdAdditionalBuy, e.MinimumBuy, e.ProductFee, e.BuyCommission, e.SellStopDateTime, e.Nav, e.NavDate, e.BuyCommission)
	if err != nil {
		log.Fatal(err)
	}
}

func InsertList(l fund.FundListViews) {
	r := listRow(l)

	_, err := db.Exec(`INSERT INTO List (
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
		r.Isin, r.Name, r.Rating, r.Risk, r.DevelopmentOneDay, r.DevelopmentOneWeek, r.DevelopmentOneMonth, r.DevelopmentThreeMonths, r.DevelopmentOneYear,
		r.DevelopmentThisYear, r.DevelopmentThreeYears, r.DevelopmentFiveYears, r.DevelopmentTenYears, r.LowCarbon, r.SharpeRatio, r.StandardDeviation, r.Capital,
		r.FossilFuelInvolvement, r.CarbonRiskScore, r.PrimaryBenchmark, r.RecommendedHoldingPeriod, r.EsgScore, r.EnvironmentalScore, r.SocialScore, r.GovernanceScore,
		r.ManagementFee, r.TotalFee, r.TransactionFee, r.OngoingFee, r.OtherFee, r.MinimumBuy, r.MinimumBuyMonthlySaving, r.HasCurrencyExchangeFee, r.NrOfOwners,
		r.OrderbookId, r.Taglist, r.Category, r.IndexFund, r.StartDate, r.SuperloanOrderbook, r.FundType, r.CompanyName, r.MatchesSustainabilityProfile, r.SustainabilityLevel,
		r.SustainabilityRating, r.SustainabilityRatingCategoryName, r.ProductInvolvements, r.ProductInvolvementViews, r.SustainabilityProfileIcon, r.CarbonSolutionsInvolvement, r.AumCoveredCarbon,
		r.ThermalCoalInvolvement, r.OilSandsExtractionInvolvement, r.ArcticOilAndGasExplorationInvolvement, r.EnvironmentalRating, r.SocialRating, r.GovernanceRating,
		r.ProspectusLink, r.CurrencyCode, r.RecommendedHoldingPeriod, r.InsertDate)
	if err != nil {
		log.Fatal(err)
	}
	// n, _ := res.RowsAffected()
	// fmt.Printf("affected %d rows\n", n)
}

func InsertDetails(d fund.FundDetails) {
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
		d.NavChangeThreeMonths, d.NavChangeOneYears, d.NavChangeThreeYears, d.NavChangeFiveYears, d.NavChangeTenYears, d.ManagementFee, "time me please")

	if err != nil {
		log.Fatal(err)
	}
	n, _ := res.RowsAffected()
	fmt.Printf("affected %d rows\n", n)
}
