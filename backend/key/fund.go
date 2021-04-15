package key

import "fmt"

func GetFundListUrl() string {
	return "https://www.avanza.se/_api/fund-guide/list"
}

func GetFundDetailsUrl(orderbookId int) string {
	return fmt.Sprintf("https://www.avanza.se/_api/fund-guide/fund-orderbook/details/%d", (orderbookId))
}
