package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func JsonPrint(uglyJson []byte) {
	var prettyJson bytes.Buffer
	json.Indent(&prettyJson, uglyJson, "", "   ")
	prettyJson.WriteTo(os.Stdout)
}
func StructPrint(v interface{}) {
	prettyJson, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(prettyJson))

}
