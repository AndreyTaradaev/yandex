package main

/*
Вам необходимо ответить на запросы узнать сумму всех элементов числовой матрицы N×M в прямоугольнике с левым верхним углом (x1, y1) и правым нижним (x2, y2)
Формат ввода

В первой строке находится числа N, M размеры матрицы (1 ≤ N, M ≤ 1000) и K — количество запросов (1 ≤ K ≤ 100000). Каждая из следующих N строк содержит по M чисел`— элементы соответствующей строки матрицы (по модулю не превосходят 1000). Последующие K строк содержат по 4 целых числа, разделенных пробелом x1 y1 x2 y2 — запрос на сумму элементов матрице в прямоугольнике (1 ≤ x1 ≤ x2 ≤ N, 1 ≤ y1 ≤ y2 ≤ M)
Формат вывода

Для каждого запроса на отдельной строке выведите его результат — сумму всех чисел в элементов матрице в прямоугольнике (x1, y1), (x2, y2)
*/

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var N, M, K int //  числа N, M размеры матрицы (1 ≤ N, M ≤ 1000) и K — количество запросов (1 ≤ K ≤ 100000)

	//получаем данные
	file, err := os.Open("input1.txt") // open file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file)        // create reader
	_, err = fmt.Fscan(rd, &N, &M, &K) //   параметры матриц и количество запросов
	if err != nil {
		fmt.Println(err)
		return
	}
	//var X [][]int = make([][]int, N, N)
	var pref [][]int = make([][]int, N+1, N+1)

	for v := range pref {
		//X[v] = make([]int, M, M)
		pref[v] = make([]int, M+1, M+1)
	}
	// заполняем матрицу
	for i := 0; i < N; i++ {
		pref[i][0] = 0

		for j := 0; j < M; j++ {
			pref[0][j] = 0
			var f int
			_, err = fmt.Fscan(rd, &f)
			if err != nil {
				fmt.Println(err)
				return
			}
			pref[i+1][j+1] = f + pref[i][j+1] + pref[i+1][j] - pref[i][j]
			//		X[i][j] = f
			//pref[i][j+1] = pref[i][j]+f
		}
	}
//	fmt.Println(pref)

/* 	var x1 = make([]int, K, K)
	var y1 = make([]int, K, K)
	var x2 = make([]int, K, K)
	var y2 = make([]int, K, K)

	for i := 0; i < K; i++ {
		var x11, y11, x22, y22 int
		_, err = fmt.Fscan(rd, &x11, &y11, &x22, &y22)
		if err != nil {
			fmt.Println(err)
			return
		}
		x1[i] = x11 - 1
		y1[i] = y11 - 1
		x2[i] = x22
		y2[i] = y22
	} */

	// читаем данные
	/* fmt.Println(X)
	fmt.Println(pref) */
	//fmt.Println(x1,"\n",y1,"\n","\n",x2,"\n",y2)
	// проходим по списку запросов

	for i := 0; i < K; i++ {
		var s int = 0
		var x11, y11, x22, y22 int
		_, err = fmt.Fscan(rd, &x11, &y11, &x22, &y22)
		if err != nil {
			fmt.Println(err)
			return
		}
		//x11, y11, x22, y22 := x1[i], y1[i], x2[i], y2[i] // координаты  матрицы сумму которой надо сосчитать
		//fmt.Println("x1 ", x11,"y1" ,y11,"x2",x22,"y2",y22 )
		//for xx := x11; xx < x22; xx++ {
			/* fmt.Println( pref[xx][y11])
			fmt.Println( pref[xx][y22]) */
			s = pref[x22][y22] - pref[x11-1][y22] - pref[x22][y11-1] + pref[x11-1][y11-1]

		//}
		fmt.Println(s)
	}

	/* for i := 0; i < K; i++ {
	var x, y int
	_, err = fmt.Fscan(rd, &x, &y)
	if err != nil {
		fmt.Println(err)
		return
	}
	X = append(X, x)
	Y = append(Y, y)
	*/

}
