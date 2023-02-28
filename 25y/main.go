package main

/*


В дощечке в один ряд вбиты гвоздики. Любые два гвоздика можно соединить ниточкой. Требуется соединить некоторые пары гвоздиков ниточками так, чтобы к каждому гвоздику была привязана хотя бы одна ниточка, а суммарная длина всех ниточек была минимальна.

Формат ввода

В первой строке входных данных записано число N — количество гвоздиков (2 ≤ N ≤ 100). В следующей строке заданы N чисел — координаты всех гвоздиков (неотрицательные целые числа, не превосходящие 10000).

Формат вывода

Выведите единственное число — минимальную суммарную длину всех ниточек.
*/

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func min(x, y int) int {

	if x > y {
		return y
	}
	return x
}

func main() {
	//var command string
	var N int // количество гвоздиков
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

	bye := make([]int, N+1, N+1)

	// заполняем виртуальных людей
	for j := 1; j < len(bye); j++ {
		var t int
		_, err = fmt.Fscan(rd, &t) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		bye[j] = t
	}
	sort.Ints(bye)
	
	sl := make([][2]int, N+3, N+3)
	sl[1][0] = 1e9
	sl[1][1] = 0

	for i := 2; i <= N; i++ {
		sl[i][0] = min(sl[i-1][0], sl[i-1][1])+(bye[i]-bye[i-1])
		sl[i][1] = sl[i-1][0]

	}
	fmt.Println(sl[N][0])
}
