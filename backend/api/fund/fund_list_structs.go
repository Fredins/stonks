package fund

type fundListRequestModel struct {
	StartIndex                     int64
	IndexFund                      bool
	SustainabilityProfile          bool
	LowCo2                         bool
	NoFossilFuelInvolvement        bool
	RegionFilter                   []string
	CountryFilter                  []string
	AlignmentFilter                []string
	IndustryFilter                 []string
	FundTypeFilter                 []string
	InterestTypeFilter             []string
	SortField                      string
	SortDirection                  string
	Name                           string
	RecommendedHoldingPeriodFilter []string
	CompanyFilter                  []string
	ProductInvolvementsFilter      []string
	RatingFilter                   []string
	SustainabilityRatingFilter     []string
	EnvironmentalRatingFilter      []string
	SocialRatingFilter             []string
	GovernanceRatingFilter         []string
}

type FundListResponseModel struct {
	FundListViews []FundListViews
	TotalNoFunds  int
	FilterCounts  map[string][]Counts
}

type Counts struct {
	Title  string
	Count  int
	Type   string
	Active bool
}

type FundListViews struct {
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
	Taglist                               []TagList
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
	ProductInvolvements                   []string
	ProductInvolvementViews               []ProductInvolvementViews
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
}

type TagList struct {
	Title           string
	FundTagCategory string
}

type ProductInvolvementViews struct {
	Product            string
	ProductDescription string
	Value              float32
}
