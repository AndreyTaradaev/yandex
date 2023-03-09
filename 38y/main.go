package main

/*
На клеточном поле, размером NxM (2 ≤ N, M ≤ 250) сидит Q (0 ≤ Q ≤ 10000) блох в различных клетках. "Прием пищи" блохами возможен только в кормушке - одна из клеток поля, заранее известная. Блохи перемещаются по полю странным образом, а именно, прыжками, совпадающими с ходом обыкновенного шахматного коня. Длина пути каждой блохи до кормушки определяется как количество прыжков. Определить минимальное значение суммы длин путей блох до кормушки или, если собраться блохам у кормушки невозможно, то сообщить об этом. Сбор невозможен, если хотя бы одна из блох не может попасть к кормушке.

Формат ввода

В первой строке входного файла находится 5 чисел, разделенных пробелом: N, M, S, T, Q. N, M - размеры доски (отсчет начинается с 1); S, T - координаты клетки - кормушки (номер строки и столбца соответственно), Q - количество блох на доске. И далее Q строк по два числа - координаты каждой блохи.

Формат вывода

Содержит одно число - минимальное значение суммы длин путей или -1, если сбор невозможен.
*/

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)

func getXY(i int) (x1, y1 int) {
	X := [...]int{-1, 1, 2, 2, 1, -1, -2, -2}
	Y := [...]int{-2, -2, -1, 1, 2, 2, 1, -1}
	return X[i], Y[i]
}

const ()

type Queue struct {
	values []int
}

// добавление в очередь
func (q *Queue) Enqueue(element int) {

	q.values = append(q.values, element) // Simply append to enqueue.
}

// значение пераого в очереди
func (q *Queue) Peek() (int, error) {

	if len(q.values) == 0 {
		return 0, fmt.Errorf("Queue is emply!")
	}
	return q.values[0], nil
}

// удаление из очереди
func (q *Queue) Dequeue() (int, error) {

	if len(q.values) == 0 {
		return 0, fmt.Errorf("Queue is emply!")
	}
	element := q.values[0]
	if len(q.values) == 1 {
		q.values = []int{}
		return element, nil
	}
	q.values = q.values[1:]
	return element, nil
}

func (q Queue) Len() int {
	return len(q.values)
}

func (q Queue) Get() []int {
	return q.values
}

func (q *Queue) String() string {
	var ret string
	ar := q.values
	for _, val := range ar {
		ret = ret + fmt.Sprintf("%d ", val)
	}
	return ret
}

var flag bool = false
var queue Queue

func PrintGraf(Graf [][]int) {
	for ind, v := range Graf {
		fmt.Println(ind,"\t", v)
	}
}
func PrintGraf2(Graf [][]int) {
	for i:=2; i<len(Graf)-2;i++ {
		fmt.Print(i,"\t")	
		for j := 2; j < len(Graf[j])-2; j++ {
			fmt.Print(Graf[i][j]," ")	
		}
		fmt.Print("\n")	
		
	}
}

func BFS(graf [][]int, startnode, endnode int) int {
	if startnode < 0 || startnode >= len(graf) {
		return 0
	}

	if endnode < 0 || endnode >= len(graf) {
		return 0
	}
	// очередь
	var q Queue //   очередь вершин
	q.Enqueue(startnode)
	size := len(graf)
	used := make([]bool, size, size) // признак  посещености вершины
	d := make([]int, size, size)     // расстояние
	p := make([]int, size, size)     // предки

	used[startnode] = true
	p[startnode] = -1
	for q.Len() != 0 {
		v, _ := q.Peek() //Извлекаем из головы очереди вершину
		q.Dequeue()
		size = len(graf[v]) //получаем количество связей у вершины
		for i := 0; i < size; i++ {
			to := graf[v][i]
			if !used[to] { //Если вершина не посещена
				used[to] = true  //посещаем ее
				q.Enqueue(to)    //и добавляем к концу очереди
				d[to] = d[v] + 1 //Считаем расстояние до вершины
				p[to] = v        //Запоминаем предка
			}
		}
	}

	if !used[endnode] {
		fmt.Println("-1")
		return -1
	} else {
		path := []int{}
		for v := endnode; v != -1; v = p[v] {
			path = append(path, v)
		}
		fmt.Println(len(path) - 1)
		if len(path)-1 == 0 {
			return len(path) - 1
		}
		for i := len(path) - 1; i >= 0; i-- {
			fmt.Print(path[i], " ")
		}
		fmt.Print("\n")

		return len(path) - 1
	}
}

func main() {
	//var command string
	var N, M, S, T, Q int // N, M - размеры доски (отсчет начинается с 1);
	//S, T - координаты клетки - кормушки (номер строки и столбца соответственно), Q - количество блох на доске.
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file)                // create reader
	_, err = fmt.Fscan(rd, &N, &M, &S, &T, &Q) //   command
	if err != nil {
		fmt.Println(err)
		return
	}

	//var  X, Y int
	rn, rm, rs, rt := N+4, M+4, S+1, T+1

	ns := make([][]int, rn, rn)
	for i := 0; i < len(ns); i++ {
		ns[i] = make([]int, rm, rm)
	}

	for i := 0; i < 8; i++ {
		x, y := getXY(i)
		ns[rs+x][rt+y] = 1
	}

	/* fmt.Println(N, M, S, T, Q)*/
//	fmt.Println(rn, rm, rs, rt)
//	PrintGraf2(ns) 
	step := 1
	waschanges := true
	for waschanges {
		waschanges = false
		for i := 2; i < len(ns)-2; i++ {
			for j := 2; j < len(ns[i])-2; j++ {
				if ns[i][j] == step {
					for z := 0; z < 8; z++ {
						x, y := getXY(z)
						if ns[i+x][j+y] == 0 {
							ns[i+x][j+y] = step + 1
							waschanges = true
						}
					}

				 	/* fmt.Println("\n")
					fmt.Println(i, j, step)
					PrintGraf2(ns)  */
				}

			}
		}
		step++
	}
	ns[rs][rt] = 0
	/* fmt.Println("\n")
	fmt.Println(rn, rm, rs, rt)
	PrintGraf(ns) */
	answer := 0
	nosolution := false
//	fmt.Println(Q,"\n")
	for i := 0; i < Q; i++ {
		var xb, yb int
		_, err = fmt.Fscan(rd, &xb, &yb) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		if  ns[xb+1][yb+1] == 0 {
			if  (xb+1 != rs || yb+1 != rt){
			nosolution = true
			break
			}
		}
		//fmt.Print(    ns[xb+1][yb+1],"\n")
		answer += ns[xb+1][yb+1]

	}
	if (nosolution){
		fmt.Println("-1")
		return
	}
	fmt.Println( answer)


}
