package main

/*
 В левом верхнем углу прямоугольной таблицы размером N×M находится черепашка. В каждой клетке таблицы записано некоторое число.
  Черепашка может перемещаться вправо или вниз, при этом маршрут черепашки
 заканчивается в правом нижнем углу таблицы.

Подсчитаем сумму чисел, записанных в клетках, через которую проползла черепашка (включая начальную и конечную клетку).
Найдите наибольшее возможное значение этой суммы и маршрут, на котором достигается эта сумма.
Формат ввода
В первой строке входных данных записаны два натуральных числа N и M, не превосходящих 100 — размеры таблицы. Далее идет N строк,
 каждая из которых содержит M чисел, разделенных пробелами — описание таблицы. Все числа в клетках таблицы целые и могут принимать
  значения от 0 до 100.
Формат вывода
Первая строка выходных данных содержит максимальную возможную сумму, вторая — маршрут, на котором достигается эта сумма.
 Маршрут выводится в виде последовательности, которая должна содержать N-1 букву D, означающую передвижение вниз и M-1 букву R,
  означающую передвижение направо. Если таких последовательностей несколько, необходимо вывести ровно одну (любую) из них.
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

	for i := 1; i < N+1; i++ {
		for j := 1; j < M+1; j++ {
			var g int
			_, err = fmt.Fscan(rd, &g) //   command
			if err != nil {
				fmt.Println(err)
				return
			}
			a[i][j] = g
		}
	}
 
	cost := make([][]int, N+1, N+1)
	for i := 0; i < N+1; i++ {
		cost[i] = make([]int, M+1, M+1)
	}
	//cost[0][0] = a[1]0]
	/* for i := 1; i < M; i++ {
		cost[0][i] = a[0][i] + cost[0][i-1]
	}

	for i := 1; i < N; i++ {
		cost[i][0] = a[i][0] + cost[i-1][0]
	} */
	road:= make([]	byte,0,M+N)
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}

	for i := 1; i < N+1; i++ {
		for j := 1; j < M+1; j++ {
			cost[i][j] = max(cost[i-1][j], cost[i][j-1]) + a[i][j]
		}
	} 

	// маршрут
		//maxcost := cost[N][M]
		  x,y := N,M
		for  x!=1 || y!=1 {
				
				//ищем  максимум из верхней и левой
				if cost[x-1][y] >=cost[x][y-1]{
					road =append(road, byte ('D'))
					x--					
			}else{
				road =append(road, byte ('R'))
				y--					
			}

		}

	
		fmt.Println(cost[N][M])	

	for i :=len(road)-1;i>=0;i-- {
		fmt.Print(string(road[i])," ")
	}
}
