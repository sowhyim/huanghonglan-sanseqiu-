package main

import "fmt"

func main() {
	Boll()
	fmt.Println(boll)
}

var boll = []int{1, 2, 3, 2, 1, 3, 1, 2, 1, 3, 2, 1, 2, 1, 2,}

func Boll() {
	var Lflag, Rflag = 0, len(boll)
	var i, j = 0, len(boll)-1
	for {
		if boll[i]!=1{
			Lflag=i
			break
		}
		i++
	}
	for {
		if boll[j]!=3{
			Rflag=j
			break
		}
		j--
	}

	for i,j=Lflag+1,Rflag-1;i<j;{
		if boll[i]==1{
			boll[i],boll[Lflag]=boll[Lflag],boll[i]
			Lflag++
		}
		if boll[i]==3{
			boll[i],boll[Rflag]=boll[Rflag],boll[i]
			Rflag--
		}

		if boll[j]==1{
			boll[j],boll[Lflag]=boll[Lflag],boll[j]
			Lflag++
		}
		if boll[j]==3{
			boll[j],boll[Rflag]=boll[Rflag],boll[j]
			Rflag--
		}
		i++
		j--
	}
}
