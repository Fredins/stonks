package pokemon

import (
	"fmt"
	"io"
	"net/http"
)

func GetRequest() []byte {
	res, err := http.Get("https://pokeapi.co/api/v2/pokemon/heatran")
	if err != nil {
		fmt.Println(err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
	}
	return b
}
