package main

/*
По данному числу N определите количество последовательностей из нулей и единиц длины N, в которых никакие три единицы не стоят рядом.
Формат ввода

Во входном файле написано натуральное число N, не превосходящее 35.
Формат вывода

Выведите количество искомых последовательностей. Гарантируется, что ответ не превосходит 231-1.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func f(n int) int {
	switch n {
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 7
	}
	return f(n-1) + f(n-2) + f(n-3)
}

func main() {
	//var command string
	var N int
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
	sl := make([]int, 35, 35)
	sl[0] = 2
	sl[1] = 4
	sl[2] = 7
	for i := 3; i < N; i++ {
		sl[i] = sl[i-1] + sl[i-2] + sl[i-3]
	}
	fmt.Println(sl[N-1])
	//fmt.Println(f(N))

}
