package main

/*
Метрополитен состоит из нескольких линий метро. Все станции метро в городе пронумерованы натуральными
числами от 1 до N. На каждой линии расположено несколько станций. Если одна и та же станция расположена
 сразу на нескольких линиях, то она является станцией пересадки и на этой станции можно пересесть с любой
 линии, которая через нее проходит, на любую другую (опять же проходящую через нее).

Напишите программу, которая по данному вам описанию метрополитена определит, с каким минимальным числом
 пересадок можно добраться со станции A на станцию B. Если данный метрополитен не соединяет все линии в одну
 систему, то может так получиться, что со станции A на станцию B добраться невозможно, в этом случае ваша
 программа должна это определить.
Формат ввода

Сначала вводится число N — количество станций метро в городе (2≤N≤100). Далее следует число M — количество
 линий метро (1≤M≤20). Далее идет описание M линий. Описание каждой линии состоит из числа Pi — количество
 станций на этой линии (2≤Pi≤50) и Pi чисел, задающих номера станций, через которые проходит линия (ни через
	 какую станцию линия не проходит дважды).

Затем вводятся два различных числа: A — номер начальной станции, и B — номер станции, на которую нам нужно
попасть. При этом если через станцию A проходит несколько линий, то мы можем спуститься на любую из них.
 Так же если через станцию B проходит несколько линий, то нам не важно, по какой линии мы приедем.
Формат вывода

Выведите минимальное количество пересадок, которое нам понадобится. Если добраться со станции A на станцию
B невозможно, программа должна вывести одно число –1 (минус один).
*/

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)

type station struct {
	num, line int
}

type Contr interface {
	int | int16 | int32 | station
}

type Queue[T Contr] struct {
	values []T
}

// добавление в очередь
func (q *Queue[T]) Enqueue(element T) {

	q.values = append(q.values, element) // Simply append to enqueue.
}

// значение пераого в очереди
func (q *Queue[T]) Peek() (*T, error) {

	if len(q.values) == 0 {
		return nil, fmt.Errorf("Queue is emply!")
	}

	return &q.values[0], nil
}

// удаление из очереди
func (q *Queue[T]) Dequeue() (*T, error) {

	if len(q.values) == 0 {
		return nil, fmt.Errorf("Queue is emply!")
	}
	element := q.values[0]
	if len(q.values) == 1 {
		q.values = []T{}
		return &element, nil
	}
	q.values = q.values[1:]
	return &element, nil
}

func (q Queue[T]) Len() int {
	return len(q.values)
}

func (q Queue[T]) Get() []T {
	return q.values
}

func PrintGraf(Graf [][]int) {
	for ind, v := range Graf {
		fmt.Println(ind, "\t", v)
	}
}

func PrintGraf2(Graf [][]int) {
	for i := 2; i < len(Graf)-2; i++ {
		fmt.Print(i, "\t")
		for j := 2; j < len(Graf[j])-2; j++ {
			fmt.Print(Graf[i][j], " ")
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
	//var q Queue[station] //   очередь вершин
	/* q.Enqueue(startnode)
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
	} */return 0
}

func main() {
	//var command string
	var N, M, first, end int // N — количество станций метро в городе (2≤N≤100). Далее следует число M — количество линий метро (1≤M≤20).

	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file)    // create reader
	_, err = fmt.Fscan(rd, &N, &M) //   command
	if err != nil {
		fmt.Println(err)
		return
	}

	ns := make([][]int, M, M)
	for i := 0; i < len(ns); i++ {
		ns[i] = make([]int, 0, 1)
	}

	//читаем по линиям

	for i := 0; i < M; i++ {
		var counst int
		_, err = fmt.Fscan(rd, &counst) //   command
		if err != nil {
			fmt.Println(err)
			return
		}

		for j := 0; j < counst; j++ {
			var st int
			_, err = fmt.Fscan(rd, &st) //   command
			if err != nil {
				fmt.Println(err)
				return
			}
			ns[i] = append(ns[i], st)
		}

	}
	_, err = fmt.Fscan(rd, &first, &end) //   command
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(ns)

	mas := make([][]station, N+1, N+1)
	for i := 0; i < len(mas); i++ {
		mas[i] = make([]station, 0, N+1)
	}

	for i := 0; i < len(ns); i++ {
		for j := 0; j < len(ns[i])-1; j++ {
			for z := j + 1; z < len(ns[i]); z++ {
				mas[ns[i][j]] = append(mas[ns[i][j]], station{ns[i][z], i + 1})
				mas[ns[i][z]] = append(mas[ns[i][z]], station{ns[i][j], i + 1})
			}
		}
	}
	fmt.Println(mas)
	/* 	 fmt.Println(N, M, S, T, Q)
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
					PrintGraf2(ns)
				}

			}
		}
		step++
	}
	ns[rs][rt] = 0
	/* fmt.Println("\n")
	fmt.Println(rn, rm, rs, rt)
	PrintGraf(ns)
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
		if ns[xb+1][yb+1] == 0 {
			if xb+1 != rs || yb+1 != rt {
				nosolution = true
				break
			}
		}
		//fmt.Print(    ns[xb+1][yb+1],"\n")
		answer += ns[xb+1][yb+1]

	}
	if nosolution {
		fmt.Println("-1")
		return
	}
	fmt.Println(answer) */

}
