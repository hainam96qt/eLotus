package main

import (
	"fmt"
	"math"
)

type GrayParticiple struct {
	IsCheck bool
	ArrBits []int
}

var input = 3

func main() {
	fmt.Println(grayCode(input))
}

func grayCode(n int) []int {
	var mapGrayInit = make(map[int]GrayParticiple)
	for i := 0; i < int(math.Pow(float64(2), float64(n))); i++ {
		mapGrayInit[i] = GrayParticiple{
			ArrBits: breakNumberToBits(i, n),
			IsCheck: false,
		}
	}

	var result []int
	for _, v := range sortCode(mapGrayInit) {
		result = append(result, bitToInt(v))
	}

	return result
}

func bitToInt(arr []int) int {
	result := 0
	pow := 0
	for i := len(arr) - 1; i > -1; i-- {
		if arr[i] == 1 {
			result += int(math.Pow(float64(2), float64(pow)))
		}
		pow++
	}
	return result
}

func sortCode(mapGrayInit map[int]GrayParticiple) (result [][]int) {
	if len(result) == 0 && len(mapGrayInit) > 0 {
		result = append(result, mapGrayInit[0].ArrBits)

		var s = mapGrayInit[0]
		s.IsCheck = true
		mapGrayInit[0] = s
	}

	for {
		previous := result[len(result)-1]
		isHaveNext := false
		// get the next
		for i := 0; i < len(mapGrayInit); i++ {
			v := mapGrayInit[i]
			if !v.IsCheck && checkDiffOne(previous, v.ArrBits) {
				var s = mapGrayInit[i]
				s.IsCheck = true
				mapGrayInit[i] = s
				result = append(result, v.ArrBits)

				isHaveNext = true

				break
			}
		}
		if !isHaveNext {
			break
		}
	}
	return result
}

func checkDiffOne(arr1 []int, arr2 []int) bool {
	count := 0
	for k, v := range arr1 {
		if v != arr2[k] {
			count++
		}
	}
	return count == 1
}

func breakNumberToBits(arg int, binaryLength int) (result []int) {
	var arrBits []int
	for i := 0; i < binaryLength; i++ {
		if arg%2 == 1 {
			arrBits = append(arrBits, 1)
		} else {
			arrBits = append(arrBits, 0)
		}
		arg = arg / 2
	}
	for i := len(arrBits) - 1; i > -1; i-- {
		result = append(result, arrBits[i])
	}
	return result
}
