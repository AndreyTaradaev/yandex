package main

/*
 В каждой клетке прямоугольной таблицы N×M записано некоторое число. Изначально игрок находится в левой верхней клетке.
  За один ход ему разрешается перемещаться в соседнюю клетку либо вправо, либо вниз (влево и вверх перемещаться запрещено).
   При проходе через клетку с игрока берут столько килограммов еды, какое число записано в этой клетке (еду берут также за 
	первую и последнюю клетки его пути).

Требуется найти минимальный вес еды в килограммах, отдав которую игрок может попасть в правый нижний угол.
Формат ввода
Вводятся два числа N и M — размеры таблицы (1≤N≤20, 1≤M≤20). Затем идет N строк по M чисел в каждой — размеры штрафов 
в килограммах за прохождение через соответствующие клетки (числа от 0 до 100).
Формат вывода
Выведите минимальный вес еды в килограммах, отдав которую можно попасть в правый нижний угол. 
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
	var N,M int // количество гвоздиков
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
a := make([][]int, N,N) 
	_, err = fmt.Fscan(rd, &M)  //   command
	if err != nil {
		fmt.Println(err)
		return		
	}

for	i:=0; i<N;i++{		
	a[i] = make([]int,M,M)	
}

 

for i := 0; i < N; i++ {
	for j := 0; j < M; j++ {
		var g int
		_, err = fmt.Fscan(rd, &g)  //   command
		if err != nil {
			fmt.Println(err)
			return		
		}
		a[i][j] = g
	}	
}	

cost := make([][]int,N,N)
for i := 0; i < N; i++ {
	cost[i] = make([]int,M,M)	
}
cost[0][0] = a[0][0]
for	i:=1; i<M;i++{			
	cost[0][i] = a[0] [i]+cost[0] [i-1]
} 

for	i:=1; i<N;i++{			
	cost[i][0] = a[i] [0]+cost [i-1][0]
} 

 min:= func(i,j int)int{
if(i>j)	{
	return j
}
return i
}

for i := 1; i < N; i++ {
	for j := 1; j < M; j++ {
	cost [i][j] = min ( cost[i-1][j]	,cost[i][j-1] ) +a[i][j]		
	}
	
}
/**/

	fmt.Println(cost[N-1][M-1])


}
