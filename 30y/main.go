package main

/*
Даны две последовательности, требуется найти и вывести их наибольшую общую подпоследовательность.

Формат ввода

В первой строке входных данных содержится число N – длина первой последовательности (1 ≤ N ≤ 1000). Во второй строке заданы члены первой последовательности
(через пробел) – целые числа, не превосходящие 10000 по модулю.

В третьей строке записано число M – длина второй последовательности (1 ≤ M ≤ 1000). В четвертой строке задаются члены второй последовательности
(через пробел) – целые числа, не превосходящие 10000 по модулю.
Формат вывода

Требуется вывести наибольшую общую подпоследовательность данных последовательностей, через пробел.
*/

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)

func min(x, y int) int {

	if x > y {
		return y
	}
	return x
}

func main() {
	//var command string
	var N, M int // количество гвоздиков
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file) // create reader
	_, err = fmt.Fscan(rd, &N)  //   command
	if err != nil {
		fmt.Println(err)
		return
	}
	ns := make([]int, N, N)
	for i := 0; i < N; i++ {
		var f int
		_, err = fmt.Fscan(rd, &f) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		ns[i] = f
	}

	_, err = fmt.Fscan(rd, &M) //   command
	if err != nil {
		fmt.Println(err)
		return
	}

	ms := make([]int, M, M)
	for i := 0; i < M; i++ {
		var f int
		_, err = fmt.Fscan(rd, &f) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		ms[i] = f
	}
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	dp := make([][]int, N+1, N+1)
	for i := 0; len(dp) > i; i++ {
		dp[i] = make([]int, M+1, M+1)
	}
	for i := 1; i < N+1; i++ {
		for j := 1; j < M+1; j++ {
			if ns[i-1] == ms[j-1] {
				dp[i][j] = dp[i-1][j-1] + 1
			} else {
				dp[i][j] = max(dp[i-1][j], dp[i][j-1])
			}
		}

	}
	i, j := N, M
	sl := make([]int, 0, 0)
	for i != 0 && j != 0 {
		if dp[i][j] == dp[i-1][j] {
			i--
		} else if dp[i][j] == dp[i][j-1] {
			j--
		} else {
			sl = append(sl, ns[i-1])
			i--
			j--
		}

	}
/* 	for _, v := range dp {
		fmt.Println(v)
	}
	fmt.Println(ns)
	fmt.Println(ms) */
	for ii := len (sl)-1; ii >=0; ii-- {
		fmt.Print( sl[ii]," ")	
	}
	
}
