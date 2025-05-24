package util

import (
	"encoding/json"
	"fmt"
)

func PrintJSON(v interface{}) {
	ctx, _ := json.MarshalIndent(v, "", "\t")
	fmt.Println("json:", string(ctx))
}
