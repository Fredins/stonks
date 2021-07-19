package key

func GetFundListUrl() string {
	return "https://www.avanza.se/_api/fund-guide/list"
}

func GetFundDetailsUrl(orderbookId string) string {
	return "https://www.avanza.se/_api/fund-guide/fund-orderbook/details/" + orderbookId
}
