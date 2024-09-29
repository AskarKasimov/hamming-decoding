package decoder

import (
	"errors"
	"fmt"
	"math"
	"slices"
)

func minimumControlBits(length float64) (int, error) {
	// using k >= log2(n + 1) from the 2**m >= n + 1
	for k := 2.0; k < length; k++ {
		if k >= math.Log2(length) {
			return int(k), nil
		}
	}
	return 0, errors.New("Not valid length of input")
}

func checkPowerOfTwo(n int) int {
	if n == 0 {
		return 1
	}
	return n & (n - 1)
}

func getSliceWithStep(slice []int, step int) []int {
	var newSlice []int
	offset := 0
	for index := step - 1; index < len(slice); index += step * 2 {
		// preventing "index out of range" error
		if index+step >= len(slice) {
			offset = index + step - len(slice)
		} else {
			offset = 0
		}
		// previous block of code gives (index+step-offset)
		for subIndex := index; subIndex < index+step-offset; subIndex++ {
			// ignoring control bits
			if checkPowerOfTwo(subIndex+1) == 0 {
				continue
			}
			newSlice = append(newSlice, slice[subIndex])
		}
	}
	return newSlice
}

func removeContolBits(slice []int, rBits []int) []int {
	for rIndex := len(rBits) - 1; rIndex >= 0; rIndex-- {
		slice = slices.Delete[[]int, int](slice, rBits[rIndex]-1, rBits[rIndex])
	}
	return slice
}

func generateControlBits(rlen int) []int {
	// making slice with control bits indexes
	rPositions := make([]int, rlen)
	for index := range len(rPositions) {
		// making powers of 2 for control bit indexes
		rPositions[index] = int(math.Pow(2.0, float64(index)))
	}
	return rPositions
}

func Decode(encoded []int, rLen int) ([]int, error) {
	var err error
	if rLen == 0 {
		// if number of control bits is unknown, set the minimum
		rLen, err = minimumControlBits(float64(len(encoded)))
		if err != nil {
			return nil, err
		}
	}

	rPositions := generateControlBits(rLen)

	errorBit := 0
	for rBitIndex, rBit := range rPositions {
		var dueValue int
		rSlice := getSliceWithStep(encoded, rBit)
		for _, iBit := range rSlice {
			// xoring to calculate valid control bit
			dueValue ^= iBit
		}
		if dueValue != encoded[rBit-1] {
			// summaring indexes to find invalid bit
			errorBit += rBitIndex + 1
			fmt.Printf("Not valid control bit: %d (value: %d), must be %d\n", rBitIndex+1, encoded[rBit-1], dueValue)
		}
	}
	if errorBit > 0 {
		// doing xor to make 0 from 1 and 1 from 0 :)
		encoded[errorBit-1] ^= encoded[errorBit-1]
		fmt.Printf("Bit №%d is incorrect\n", errorBit)
	}
	return removeContolBits(encoded, rPositions), nil
}
