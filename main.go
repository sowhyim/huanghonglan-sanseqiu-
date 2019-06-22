package main

import "fmt"

func main() {
	Boll(boll)
	fmt.Println(boll)
	Boll(boll1)
	fmt.Println(boll1)
	Boll(boll2)
	fmt.Println(boll2)
	Boll(boll3)
	fmt.Println(boll3)
	Boll(boll4)
	fmt.Println(boll4)
}

var boll = []int{1, 2, 3, 2, 1, 3, 1, 2, 1, 3, 2, 1, 2, 1, 2}
var boll1 = []int{}
var boll2 = []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
var boll3 = []int{2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2, 2}
var boll4 = []int{3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3, 3}

func Boll(boll []int) {
	if len(boll) == 0 {
		return
	}
	var Lflag, Rflag = 0, len(boll)
	var i, j = 0, len(boll)-1
	for {
		if  boll[i] != 1 {
			Lflag = i
			break
		}
			i++
		if i == len(boll)-1 {
			return
		}
	}

	for {
		if j >= 0 && boll[j] != 3 {
			Rflag = j
			break
		}
		j--
		if j == 0 {
			return
		}
	}


	for i, j = Lflag+1, Rflag-1; i < j; {
		/*if boll[i] == 1 {
			boll[i], boll[Lflag] = boll[Lflag], boll[i]
			Lflag++
		}
		if boll[i] == 3 {
			boll[i], boll[Rflag] = boll[Rflag], boll[i]
			Rflag--
		}

		if boll[j] == 1 {
			boll[j], boll[Lflag] = boll[Lflag], boll[j]
			Lflag++
		}
		if boll[j] == 3 {
			boll[j], boll[Rflag] = boll[Rflag], boll[j]
			Rflag--
		}*/
		switch boll[i] {
		case 1:
			boll[i], boll[Lflag] = boll[Lflag], boll[i]
			Lflag++
		case 3:
			boll[i], boll[Rflag] = boll[Rflag], boll[i]
			Rflag--
		}
		switch boll[j] {
		case 1:
			boll[j], boll[Lflag] = boll[Lflag], boll[j]
			Lflag++
		case 3:
			boll[j], boll[Rflag] = boll[Rflag], boll[j]
			Rflag--
		}
		i++
		j--
	}
}
