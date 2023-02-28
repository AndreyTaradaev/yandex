package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func SortAndRemoveDuplicates(i []int) []int {
	if len(i) < 1 {
		return i
	}
	// sort
	sort.Sort(sort.Reverse(sort.IntSlice(i)))

	prev := 1
	for curr := 1; curr < len(i); curr++ {
		if i[curr-1] != i[curr] {
			i[prev] = i[curr]
			prev++
		}
	}

	return i[:prev]

}

func main() {

	var n, k int
	var nums []int = make([]int, 0, 200)
	var knums []int = make([]int, 0, 200)

	file, err := os.Open("input4.txt") // open file
	if err != nil {
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file) // create reader
	_, err = fmt.Fscan(rd, &n)  // scaN  count  Diego
	if err != nil {
		fmt.Println(err)
		return
	}	
	for i := 0; i < n; i++ {
		_, err = fmt.Fscan(rd, &k)
		if err != nil {
			return
		}
		nums = append(nums, k)
	}

	_, err = fmt.Fscan(rd, &k)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < k; i++ {
		var r int
		_, err = fmt.Fscan(rd, &r)
		if err != nil {
			return
		}
		knums = append(knums, r)
	}
	//fmt.Println(nums)
	// важно убрать дубликаты
	nums = SortAndRemoveDuplicates(nums)

	//  проход по коллекционерам
	for _, v := range knums {
		fsearch := func(i int) bool { return nums[i] < v }
		ind := sort.Search(len(nums), fsearch)
		fmt.Println( len(nums)-ind)
	}
}
