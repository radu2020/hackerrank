package main

import "fmt"

// Complete the sockMerchant function below.
func sockMerchant(ar []int32) int32 {
	myMap := make(map[int32]int32)
	for i:=0; i<len(ar)-1; i++{
		if _, ok := myMap[ar[i]]; ok {
			continue
		} else {
			myMap[ar[i]] = 1
		}
		for j:=i+1; j<len(ar); j++ {
			if ar[i] == ar[j] {
				val := myMap[ar[i]]
				myMap[ar[i]] = val + 1
			}
		}
	}
	var totalPairs int32 = 0
	for _, v := range myMap {
		if v > 1 {
			res := v / 2
			totalPairs += res
		}
	}
	return totalPairs
}

func main() {
	list:= []int32{10, 20, 20, 10, 10, 30, 50, 10, 20}
	ret := sockMerchant(list)
	fmt.Println("Types of socks:", list)
	fmt.Println("We could make [", ret, "] total pairs of socks.")
}
