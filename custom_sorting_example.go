package main

import (
	"fmt"
	"sort"
)

type byCustomSort []string

func (s byCustomSort) Len() int {
	return len(s)
}

func (s byCustomSort) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byCustomSort) Less(i, j int) bool {
	if (len(s[i]) % 3 < len(s[j]) % 3) {
		return true
	} else {
		if (len(s[i]) % 3 == len(s[j]) % 3) {
			if s[i] < s[j] {
				return false
			} else {
				return true
			}
		} else {
			return false
		}
	}
}

func main() {
	a := []string{"a","ba","ab","z","cab","bac","acb","defgxyz"}

	fmt.Println("Given an array of string, sort the array by 1) descending modulus 3 of array element length, then ascending alphabetical order")
	fmt.Println(a)
	sort.Sort(sort.Reverse(byCustomSort(a)))
	fmt.Println(a)
}
