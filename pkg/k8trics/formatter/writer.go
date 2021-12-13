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
	printMap(dataMap, "", skips...)
}

func printMap(data map[string]interface{}, prefix string, skips ...string) {
	for k, v := range data {
		if utils.ContainsString(skips, k) {
			continue
		}

		fmt.Print(prefix)
		color.New(color.FgHiGreen).Printf("%s:", k)

		switch c := v.(type) {
		case map[string]interface{}:
			fmt.Println()
			printMap(c, fmt.Sprintf("%s  ", prefix))
		case []interface{}:
			for _, val := range c {
				fmt.Print("\n=>")
				switch cv := val.(type) {
				case map[string]interface{}:
					printMap(cv, fmt.Sprintf("%s  ", prefix))
				default:
					fmt.Printf(" %v\n", c)
				}
			}
		default:
			fmt.Printf(" %v\n", c)
		}
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
