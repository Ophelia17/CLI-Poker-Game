package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)

	for i := 0; i < t; i++ {
		Arr1 := GetInput()
		Arr2 := GetInput()

		temp1 := GenerateTemp(Arr1)
		temp2 := GenerateTemp(Arr2)

		r1 := SetRank(Arr1, temp1)
		r2 := SetRank(Arr2, temp2)

		fmt.Printf("For Game %v : ", i+1)

		if r1 > r2 {
			fmt.Println("A wins")
		} else if r2 > r1 {
			fmt.Println("A wins")
		} else {
			if ResolveTie(r1, GenerateValues(Arr1), GenerateValues(Arr2)) == 1 {
				fmt.Println("A wins")
			} else if ResolveTie(r1, GenerateValues(Arr1), GenerateValues(Arr2)) == 0 {
				fmt.Println("B wins")
			} else {
				fmt.Println("It's a Tie")
			}
		}
	}
}
