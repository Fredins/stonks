package fund

type FundDetailsResponseModel struct {
	OrderbookId                   string // int in list but string in details
	Name                          string
	MinimumThresholdAdditionalBuy float32
	MinimumThresholdBuy           float32
	ProductFee                    float32
	ProspectusLink                string
	BuyStopDateTime               string
	SellStopDateTime              string
	Nav                           float32
	NavDate                       string
	BuyCommission                 float32
	HasCurrencyExchangeFee        bool
	Rating                        int
	Risk                          string
	RiskNumber                    int
	Owners                        int
	Description                   string
	NavChangeThreeMonths          float32
	NavChangeOneYears             float32
	NavChangeThreeYears           float32
	NavChangeFiveYears            float32
	NavChangeTenYears             float32
	ManagementFee                 float32
}
