package tools

import (
	"bytes"
	"encoding/json"
	"fmt"
	"os"
)

func jsonPrint(uglyJson []byte) {
	var prettyJson bytes.Buffer
	json.Indent(&prettyJson, uglyJson, "", "   ")
	prettyJson.WriteTo(os.Stdout)
}
func structPrint(v interface{}) {
	prettyJson, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(prettyJson))

}
