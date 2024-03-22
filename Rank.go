package main

import (
	"sort"
)

var StraightBit int = -1

func OnePair(temp []int) bool {
	temp2 := make([]int, len(temp))
	_ = copy(temp2, temp)
	sort.Ints(temp2)

	return (temp2[12] == 2)
}

func TwoPairs(temp []int) bool {
	temp2 := make([]int, len(temp))
	_ = copy(temp2, temp)
	sort.Ints(temp2)

	return (temp2[12] == 2 && temp2[11] == 2)

}

func ThreeOfAKind(temp []int) bool {
	temp2 := make([]int, len(temp))
	_ = copy(temp2, temp)
	sort.Ints(temp2)

	return (temp2[12] == 3)

}

func Straight(temp []int) bool {
	x := false
	cnt := 0

	for i := range 13 {
		if x {
			if temp[i] == 1 {
				cnt++
				StraightBit = i
			} else {
				break
			}
		} else {
			if temp[i] == 1 {
				x = true
				cnt++
			}
		}
	}

	return (cnt == 5)
}

func Flush(Arr [5]string) bool {
	v := []int{0, 0, 0, 0}
	for i := range 5 {
		switch t := Arr[i][1]; t {
		case 'H':
			v[0]++

		case 'D':
			v[1]++

		case 'S':
			v[2]++

		case 'C':
			v[3]++
		}
	}
	sort.Ints(v)
	return (v[3] == 5)
}

func FullHouse(temp []int) bool {
	temp2 := make([]int, len(temp))
	_ = copy(temp2, temp)
	sort.Ints(temp2)

	return (temp2[12] == 3 && temp2[11] == 2)

}

func FourOfAKind(temp []int) bool {
	temp2 := make([]int, len(temp))
	_ = copy(temp2, temp)
	sort.Ints(temp2)

	return (temp2[12] == 4)
}

func StraightFlush(Arr [5]string, temp []int) bool {
	return (Flush(Arr) && Straight(temp))
}

func RoyalFlush(Arr [5]string, temp []int) bool {
	var x bool = Flush(Arr) && Straight(temp)
	return (x && (StraightBit == 12))
}

func ResolveTie(r int, a []int, b []int) int {
	sort.Ints(a)
	sort.Ints(b)

	if r == 10 {
		return -1
	} else if r == 9 || r == 5 {
		sort.Ints(a)
		sort.Ints(b)

		if a[4] > b[4] {
			return 1
		} else if a[4] < b[4] {
			return 0
		} else {
			return -1
		}
	} else if r == 8 {
		a1 := -1
		a2 := -1
		b1 := -1
		b2 := -1

		if a[4] == a[3] {
			a1 = a[0]
			a2 = a[4]
		} else {
			a1 = a[4]
			a2 = a[0]
		}

		if b[4] == b[3] {
			b1 = b[0]
			b2 = b[4]
		} else {
			b1 = b[4]
			b2 = b[0]
		}

		if a2 > b2 {
			return 1
		} else if b2 > a2 {
			return 0
		} else {
			if a1 > b1 {
				return 1
			} else if b1 > a1 {
				return 0
			} else {
				return -1
			}
		}
	} else if r == 7 {
		a1 := -1
		a2 := -1
		b1 := -1
		b2 := -1

		if a[3] == a[2] {
			a1 = a[0]
			a2 = a[4]
		} else {
			a1 = a[4]
			a2 = a[0]
		}

		if b[3] == b[2] {
			b1 = b[0]
			b2 = b[4]
		} else {
			b1 = b[4]
			b2 = b[0]
		}

		if a2 > b2 {
			return 1
		} else if b2 > a2 {
			return 0
		} else {
			if a1 > b1 {
				return 1
			} else if b1 > a1 {
				return 0
			} else {
				return -1
			}
		}

	} else if r == 6 || r == 1 {
		for i := range 5 {
			if a[4-i] > b[4-i] {
				return 1
			} else if a[4-i] < b[4-i] {
				return 0
			}
		}
		return -1
	} else if r == 4 {
		a1 := -1
		a2 := -1
		a3 := -1
		b1 := -1
		b2 := -1
		b3 := -1

		if a[0] == a[1] {
			a1 = a[0]
			a2 = a[3]
			a3 = a[4]
		} else if a[3] == a[4] {
			a1 = a[4]
			a2 = a[0]
			a3 = a[1]
		} else {
			a1 = a[1]
			a2 = a[0]
			a3 = a[4]
		}
		if b[0] == b[1] {
			b1 = b[0]
			b2 = b[3]
			b3 = b[4]
		} else if b[3] == b[4] {
			b1 = b[4]
			b2 = b[0]
			b3 = b[1]
		} else {
			b1 = b[1]
			b2 = b[0]
			b3 = b[4]
		}

		if a1 > b1 {
			return 1
		} else if b1 > a1 {
			return 0
		} else {
			if a3 > b3 {
				return 1
			} else if b3 > a3 {
				return 0
			} else {
				if a2 > b2 {
					return 1
				} else if b2 > a2 {
					return 0
				} else {
					return -1
				}
			}
		}
	} else if r == 3 {
		a1 := -1
		a2 := -1
		a3 := -1
		b1 := -1
		b2 := -1
		b3 := -1

		if a[0] == a[1] && a[2] == a[3] {
			a1 = a[0]
			a2 = a[2]
			a3 = a[4]
		} else if a[0] == a[1] && a[3] == a[4] {
			a1 = a[0]
			a2 = a[4]
			a3 = a[2]
		} else {
			a3 = a[0]
			a1 = a[1]
			a2 = a[4]
		}

		if b[0] == b[1] && b[2] == b[3] {
			b1 = b[0]
			b2 = b[2]
			b3 = b[4]
		} else if b[0] == b[1] && b[3] == b[4] {
			b1 = b[0]
			b2 = b[4]
			b3 = b[2]
		} else {
			b3 = b[0]
			b1 = b[1]
			b2 = b[4]
		}

		if a2 > b2 {
			return 1
		} else if b2 > a2 {
			return 0
		} else {
			if a1 > b1 {
				return 1
			} else if b1 > a1 {
				return 0
			} else {
				if a3 > b3 {
					return 1
				} else if b3 > a3 {
					return 0
				} else {
					return -1
				}
			}
		}
	} else if r == 2 {
		a1 := -1
		a2 := -1
		a3 := -1
		a4 := -1
		b1 := -1
		b2 := -1
		b3 := -1
		b4 := -1

		if a[0] == a[1] {
			a1 = a[0]
			a2 = a[2]
			a3 = a[3]
			a3 = a[4]
		} else if a[1] == a[2] {
			a1 = a[1]
			a2 = a[0]
			a3 = a[3]
			a4 = a[4]
		} else if a[2] == a[3] {
			a1 = a[2]
			a2 = a[0]
			a3 = a[1]
			a4 = a[4]
		} else {
			a1 = a[3]
			a2 = a[0]
			a3 = a[1]
			a4 = a[2]
		}

		if b[0] == b[1] {
			b1 = b[0]
			b2 = b[2]
			b3 = b[3]
			b3 = b[4]
		} else if b[1] == b[2] {
			b1 = b[1]
			b2 = b[0]
			b3 = b[3]
			b4 = b[4]
		} else if b[2] == b[3] {
			b1 = b[2]
			b2 = b[0]
			b3 = b[1]
			b4 = b[4]
		} else {
			b1 = b[3]
			b2 = b[0]
			b3 = b[1]
			b4 = b[2]
		}

		if a1 > b1 {
			return 1
		} else if b1 > a1 {
			return 0
		} else {
			if a4 > b4 {
				return 1
			} else if b4 > a4 {
				return 0
			} else {
				if a3 > b3 {
					return 1
				} else if b3 > a3 {
					return 0
				} else {
					if a2 > b2 {
						return 1
					} else if b2 > a2 {
						return 0
					} else {
						return -1
					}
				}
			}
		}

	} else {
		return -1
	}

}
