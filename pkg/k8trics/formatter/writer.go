package formatter

import (
	"encoding/json"
	"fmt"

	"github.com/fatih/color"
	"github.com/sagacious-labs/kcli/pkg/utils"
	"golang.org/x/term"
)

func Write(data []byte, skips ...string) {
	mp := map[string]interface{}{}

	if err := json.Unmarshal(data, &mp); err != nil {
		return
	}

	dataMap, ok := mp["data"].(map[string]interface{})
	if !ok {
		return
	}

	printFullLine()
	prettyPrint(dataMap, "", false, skips...)
}

func prettyPrint(data interface{}, prefix string, adjustFirst bool, skips ...string) {
	switch c := data.(type) {
	case map[string]interface{}:
		fmt.Println()
		prettyPrintMap(c, prefix, adjustFirst, skips...)
	case []interface{}:
		prettyPrintSlice(c, prefix, adjustFirst, skips...)
	default:
		fmt.Printf(" %v\n", c)
	}
}

func prettyPrintMap(data map[string]interface{}, prefix string, adjustFirst bool, skips ...string) {
	for k, v := range data {
		// Skip the keys mentioned in the skips
		if utils.ContainsString(skips, k) {
			continue
		}

		// Print the key with given prefix
		color.New(color.FgHiGreen).Printf("%s%s:", prefix, k)

		prettyPrint(v, generateWhitespacePrefix(len(prefix)+2), adjustFirst, skips...)
	}
}

func prettyPrintSlice(data []interface{}, prefix string, adjustFirst bool, skips ...string) {
	for _, val := range data {
		fmt.Printf("\n%s-", prefix)
		prettyPrint(val, generateWhitespacePrefix(len(prefix)+2), true, skips...)
	}
}

func printFullLine() {
	_, w, err := term.GetSize(0)
	if err != nil {
		w = 52
	}

	line := ""
	for i := 0; i < w; i++ {
		line += "="
	}

	color.New(color.FgMagenta).Println(line)
}

func generateWhitespacePrefix(num int) string {
	prefix := ""

	for i := 0; i < num; i++ {
		prefix += " "
	}

	return prefix
}
