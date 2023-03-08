package main

/*


В неориентированном графе требуется найти длину минимального пути между двумя вершинами.
Формат ввода

Первым на вход поступает число N – количество вершин в графе (1 ≤ N ≤ 100). Затем записана матрица смежности (0 обозначает отсутствие ребра, 1 – наличие ребра). Далее задаются номера двух вершин – начальной и конечной.

Формат вывода

Выведите L – длину кратчайшего пути (количество ребер, которые нужно пройти).

Если пути нет, нужно вывести -1.
*/

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)

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

/*
	 func dfs(graf [][]int, Used []bool, v, parent int) {
		if flag {
			return
		}
		Used[v] = true //При входе в вершину окрашиваем ее в серый цвет
		st.Push(v)
		for i := 0; i < len(graf[v]); i++ {
			if graf[v][i] == parent{	//  граф указвает на тот из которого пришли
				continue
			}

			if !Used[graf[v][i]] {
				dfs(graf, Used, graf[v][i], v)
				if flag {
					return
				}
			} else if 	*(st.front()) == graf[v][i] {
				flag = true
				return
			}
		}
		st.Pop()
	}
*/
func PrintGraf(Graf [][]int) {
	for ind, v := range Graf {
		fmt.Println(ind, v)
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
		return -1
	} else {
		path := []int{}
		for v := endnode; v != -1; v = p[v] {
			path = append(path, v)
		}
		/*  for i:= len(path)-1;i>=0; i--{
			fmt.Print( path[i]," " )
			 }
		 fmt.Print( "\n" ) */

		return len(path) - 1
	}
}

func main() {
	//var command string
	var N int // размер матрицы
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
	ns := make([][]int, N+1, N+1)

	for i := 1; i < len(ns); i++ { //строка
		ns[i] = make([]int, 0, 10)
		for j := 1; j < len(ns); j++ {
			var f int
			_, err = fmt.Fscan(rd, &f) //   command
			if err != nil {
				fmt.Println(err)
				return
			}
			if f != 0 {
				ns[i] = append(ns[i], j)
			}
		}
	}
	var startnode, endnode int
	_, err = fmt.Fscan(rd, &startnode, &endnode) //   command
	if err != nil {
		fmt.Println(err)
		return
	}

	//PrintGraf(ns)
	fmt.Println(BFS(ns, startnode, endnode))

}
