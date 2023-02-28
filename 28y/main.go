package main

/*


Дана прямоугольная доска N × M (N строк и M столбцов). В левом верхнем углу находится шахматный конь,
 которого необходимо переместить в правый нижний угол доски. В данной задаче конь может перемещаться на две
  клетки вниз и одну клетку вправо или на одну клетку вниз и две клетки вправо.

Необходимо определить, сколько существует различных маршрутов, ведущих из левого верхнего в правый нижний угол.

Формат ввода

Входной файл содержит два натуральных числа N и M , .
Формат вывода

В выходной файл выведите единственное число — количество способов добраться конём до правого нижнего угла доски.
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
	a := make([][]int, N+1, N+1)
	_, err = fmt.Fscan(rd, &M) //   command
	if err != nil {
		fmt.Println(err)
		return
	}

	for i := 0; i < N+1; i++ {
		a[i] = make([]int, M+1, M+1)
	}
	a[1][1] = 1
	for i := 2; i < N+1; i++ {
		for j := 2; j < M+1; j++ {
			a[i][j] = a[i-1][j-2] + a[i-2][j-1]

		}
	}

		fmt.Println(a[N][M])


}
