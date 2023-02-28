package main

/*


На клетчатой плоскости закрашено K клеток. Требуется найти минимальный по площади прямоугольник, со сторонами, параллельными линиям сетки, покрывающий все закрашенные клетки.

Формат ввода

Во входном файле, на первой строке, находится число K (1 ≤ K ≤ 100). На следующих K строках находятся пары чисел Xi и Yi – координаты закрашенных клеток (|Xi|, |Yi| ≤ 109).

Формат вывода

Выведите в выходной файл координаты левого нижнего и правого верхнего углов прямоугольника.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var K int //  число K (1 ≤ K ≤ 100).
	var X []int = make([]int, 0, 100)
	var Y []int = make([]int, 0, 100)

	//и конечного секторов раздела (1 ≤ ai ≤ bi ≤ M).
	//получаем данные
	file, err := os.Open("input.txt") // open file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file) // create reader
	_, err = fmt.Fscan(rd, &K)  //  время запроса клиента
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i < K; i++ {
		var x, y int
		_, err = fmt.Fscan(rd, &x, &y)
		if err != nil {
			fmt.Println(err)
			return
		}
		X = append(X, x)
		Y = append(Y, y)
	}
	var top, bottom = Y[0], Y[0]

	var left, rigth int = X[0], X[0]

	for i := 1; i < K; i++ {
		if X[i] > rigth {
			rigth = X[i]
		}
		if X[i] < left {
			left = X[i]
		}
		if Y[i] > top {
			top = Y[i]
		}
		if Y[i] < bottom {
			bottom = Y[i]
		}

	}
	//fmt.Println( X,"\n",Y)
	fmt.Println(left, bottom, rigth, top)
}
