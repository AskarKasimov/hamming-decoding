package main

import (
	"fmt"
	"strconv"

	decoder "github.com/askarkasimov/hamming-decoding/pkg"
)

func main() {
	var code string
	fmt.Print("Enter the code to analyze: ")
	fmt.Scan(&code)

	if len(code) < 7 {
		fmt.Println("Provide at least 7 bits!")
		return
	}

	var rLen int = 0
	fmt.Println("Using the mimimum possible number of control bits...")

	// TODO: think about custom control bit number
	// fmt.Print("Enter the number of control bits (leave 0 if possible minimum): ")
	// fmt.Scan(&rLen)

	// making slice for input code
	arrayInput := make([]int, len(code))
	for i, symbol := range code {
		// converting runes to ints
		digit, err := strconv.Atoi(string(symbol))
		if err != nil || (digit != 0 && digit != 1) {
			fmt.Println("Not valid input. Provide only 1 and 0")
			return
		}
		arrayInput[i] = digit
	}

	decoded, err := decoder.Decode(arrayInput, rLen)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("True information:")
	for _, decodedBits := range decoded {
		fmt.Print(decodedBits)
	}
}
