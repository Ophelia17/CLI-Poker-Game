package main

import (
	"fmt"
)

func GetInput() [5]string {
	Arr := [5]string{}
	for i := range 5 {
		fmt.Scan(&Arr[i])
	}
	return Arr
}

func GenerateTemp(Arr [5]string) []int {
	temp := []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}

	for i := range 5 {
		switch t := Arr[i][0] - 50; t {
		case 34:
			temp[8]++
		case 24:
			temp[9]++
		case 31:
			temp[10]++
		case 25:
			temp[11]++
		case 15:
			temp[12]++
		default:
			temp[t]++
		}
	}

	return temp
}

func GenerateValues(Arr [5]string) []int {
	v := [5]int{}

	for i := range 5 {
		switch t := Arr[i][0] - 50; t {
		case 34:
			v[i] = 8
		case 24:
			v[i] = 9
		case 31:
			v[i] = 10
		case 25:
			v[i] = 11
		case 15:
			v[i] = 12
		case 0:
			v[i] = 0
		case 1:
			v[i] = 1
		case 2:
			v[i] = 2
		case 3:
			v[i] = 3
		case 4:
			v[i] = 4
		case 5:
			v[i] = 5
		case 6:
			v[i] = 6
		case 7:
			v[i] = 7
		default:
			v[i] = -1
		}

	}
	v_temp := []int{v[0], v[1], v[2], v[3], v[4]}
	return v_temp
}

func SetRank(Arr [5]string, temp []int) int {
	if RoyalFlush(Arr, temp) {
		return 10
	} else if StraightFlush(Arr, temp) {
		return 9
	} else if FourOfAKind(temp) {
		return 8
	} else if FullHouse(temp) {
		return 7
	} else if Flush(Arr) {
		return 6
	} else if Straight(temp) {
		return 5
	} else if ThreeOfAKind(temp) {
		return 4
	} else if TwoPairs(temp) {
		return 3
	} else if OnePair(temp) {
		return 2
	} else {
		return 1
	}
}
