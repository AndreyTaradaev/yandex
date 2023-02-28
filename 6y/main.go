package main

/*

Васин жесткий диск состоит из M секторов. Вася последовательно устанавливал
на него различные операционные системы следующим методом: он создавал новый раздел
диска из последовательных секторов, начиная с сектора номер ai и
до сектора bi включительно, и устанавливал на него очередную систему.
При этом, если очередной раздел хотя бы по одному сектору пересекается
с каким-то ранее созданным разделом, то ранее созданный раздел «затирается»,
и операционная система, которая на него была установлена, больше не может быть загружена.

Напишите программу, которая по информации о том, какие разделы на диске создавал Вася,
определит, сколько в итоге работоспособных операционных систем установлено и работает
в настоящий момент на Васином компьютере.
Формат ввода

Сначала вводятся натуральное число M — количество секторов на жестком диске (1 ≤ M ≤ 109)
и целое число N — количество разделов, которое последовательно создавал Вася (0 ≤ N ≤ 1000).

Далее идут N пар чисел ai и bi, задающих номера начального и конечного секторов раздела (1 ≤ ai ≤ bi ≤ M).
Формат вывода

Выведите одно число — количество работающих операционных систем на Васином компьютере.
*/

import (
	"bufio"
	"fmt"
	"os"
	//"sort"
)

type Pair struct {	
	begin int
	end   int
}

type Pairs []Pair

func intersection(first, second Pair) bool {
	if first.begin <= second.begin && first.end >= second.begin {
		return true
	}
	if first.begin <= second.end && first.end >= second.begin {
		return true
	}
	return false
}

func main() {
	var M int                           //  количество секторов на жестком диске (1 ≤ M ≤ 109)
	var N int                           // количество разделов, которое последовательно создавал Вася (0 ≤ N ≤ 1000)
	var sa Pairs = make([]Pair, 0, 30) //Далее идут N пар чисел ai и bi, задающих номера начального
	//и конечного секторов раздела (1 ≤ ai ≤ bi ≤ M).
	//получаем данные
	file, err := os.Open("input.txt") // open file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file) // create reader
	_, err = fmt.Fscan(rd, &M)  // scaN  count  sector
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Fscan(rd, &N) // scaN  count  partition
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < N; i++ {
		var beg, end int
		_, err = fmt.Fscan(rd, &beg, &end) // scaN  count  partition
		if err != nil {
			fmt.Println(err)
			return
		}
		sa = append(sa, Pair{ beg, end})
	}
	 for i := 0; i < len(sa); i++ {
		for j := i+1 ; j < len(sa); j++ {
			if(intersection(sa[i],sa[j] )){
				sa[i].begin = -1
				sa[i].end = -1 
				N--
			}

		}

	} 

fmt.Println( N)
}

/* 	for i := 0; i < N; i++ {
		var k int
		_, err = fmt.Fscan(rd, &k)
		if err != nil {
			fmt.Println(err)
			return
		}
		sa = append(sa, k)
	}
	lengood := 0

	for i := 1; i < len(sa); i++ {
		lengood += min(sa[i-1], sa[i])
	}

	fmt.Println(lengood)


) */
