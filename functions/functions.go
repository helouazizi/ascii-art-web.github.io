package functions

import (
	"fmt"
	"os"
	"strings"
)

// this func just to read the banner file
func ReadFile(filename string) []string {
	data, err := os.ReadFile(filename)
	// handle err
	if err != nil {
		fmt.Println("Usage: go run .[STRING] [BANNER]\n\nExample: go run .  something standard")
		os.Exit(0)
	}
	// handling if the banner file was writenf by windows
	stringdata := string(data)
	stringdata = strings.ReplaceAll(stringdata, "\r\n", "\n")

	result := strings.Split(stringdata, "\n")
	return result
}

// this is the the traitment functions
func TraitmentData(text []string, arg string) string {
	// Normalize newlines
	arg = strings.ReplaceAll(arg, "\r\n", "\\n")

	for _, char := range arg {
		if char < 32 || char > 126 {
			fmt.Println(char)
			return " error : one of this charachter not in range "
		}
	}
	result := ""

	words := strings.Split(arg, "\\n")

	/////////////////////////////////
	// this part just for newlines
	count := 0
	for _, test := range words {
		if test == "" {
			count++
		}
	}
	// in case the data is all new line
	if count == (len(arg)/2)+1 {
		for i := 0; i < (len(arg) / 2); i++ {
			result += "\n"
		}
	} else {
		result = Final_result(text, words)
	}

	return result
}

func Final_result(arrData, words []string) string {
	result := ""
	for k := 0; k < len(words); k++ {
		if words[k] == "" {
			result += "\n"
			continue
		}
		for i := 0; i < 8; i++ {
			for j := 0; j < len(words[k]); j++ {
				Ascii := (int(words[k][j] - 32))

				start := Ascii*8 + Ascii + 1 + i

				result += arrData[start]

			}
			result += "\n"
		}
	}

	return result
}
