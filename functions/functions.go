package functions

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// this func just to read the banner file
func ReadFile(filename string) []string {
	data, err := os.ReadFile(filename)
	// handle err
	if err != nil {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}
	// handling if the banner file was writenf by windows
	stringdata := string(data)
	stringdata = strings.ReplaceAll(stringdata, "\r\n", "\n")

	result := strings.Split(stringdata, "\n")

	return result
}

// the result file name
func Extract_Result_File_Name(resultFile string) string {
	if strings.HasPrefix(resultFile, "--output=") && strings.HasSuffix(resultFile, ".txt") || resultFile == "" {
		re := regexp.MustCompile(`--output=`)
		resultFile = re.ReplaceAllString(resultFile, "")
	} else {
		fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]\n\nExample: go run . --output=<fileName.txt> something standard")
		os.Exit(0)
	}

	return resultFile
}

// this is the the traitment functions
func TraitmentData(text []byte, arg string, status bool) string {
	// cheeck if the char is in range or not if
	for _, char := range arg {
		if char < 32 || char > 126 {
			fmt.Println("Ereur : one of this carcters is not in range")
			return ""
		}
	}
	result := ""

	arrData := strings.Split(string(text), "\n")

	words := strings.Split(arg, `\n`)

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
		result = Final_result(arrData, words)
	}
	if status {
		//os.WriteFile(resultFile, []byte(result), 0o777)

	} else {
		return result
	}
	return result
}

// traitment the data if it have charachters
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
				if words[k][j] < 32 || words[k][j] > 126 {
					fmt.Println(" error : one of this charachter not in range ")
					os.Exit(0)
				} else {
					result += arrData[start]
				}

			}
			result += "\n"
		}
	}

	return result
}

func TrueOutput(resultFile string) bool {
	if strings.HasPrefix(resultFile, "--output=") && strings.HasSuffix(resultFile, ".txt") {
		return true
	}
	return false
}
