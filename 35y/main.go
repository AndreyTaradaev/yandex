package main

/*
Дан неориентированный граф. Требуется определить, есть ли в нем цикл, и, если есть, вывести его.

Формат ввода

В первой строке дано одно число n — количество вершин в графе ( 1 ≤ n ≤ 500 ).
Далее в n строках задан сам граф матрицей смежности.

Формат вывода

Если в иcходном графе нет цикла, то выведите «NO». Иначе, в первой строке выведите «YES»,
во второй строке выведите число k — количество вершин в цикле, а в третьей строке выведите
k различных чисел — номера вершин, которые принадлежат циклу в порядке обхода (обход можно начинать с
	 любой вершины цикла). Если циклов несколько, то выведите любой. */

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)

type Stack struct {
	values []int
}

// добавление в очередь
func (q *Stack) Push(element int) {

	q.values = append(q.values, element) // Simply append to enqueue.
}

// значение пераого в очереди
func (q *Stack) Peek() (*int, error) {
	s := len(q.values)
	if s == 0 {
		return nil, fmt.Errorf("error")
	}
	return &q.values[s-1], nil
}

// удаление из очереди
func (q *Stack) Pop() (*int, error) {
	s := len(q.values)
	if s == 0 {
		return nil, fmt.Errorf("error")
	}
	element := q.values[s-1]
	if len(q.values) == 1 {
		q.values = []int{}
		return &element, nil
	}
	q.values = q.values[:s-1]
	return &element, nil
}

func (q *Stack) pop() *int {
	pval, err := q.Pop()
	if err != nil {
		return nil
	}
	return pval
}

func (q Stack) Len() int {
	return len(q.values)
}

var flag bool = false
var st Stack

func dfs(graf [][]int, Used []bool, v, parent int) {
	if flag {
		return
	}
	Used[v] = true //При входе в вершину окрашиваем ее в серый цвет
	st.Push(v)
	for i := 0; i < len(graf[v]); i++ {

		if !Used[graf[v][i]] {
			dfs(graf, Used, graf[v][i], v)
			if flag {
				return
			}
		} else if graf[v][i] != parent {
			flag = true
			return
		}
	}
	st.Pop()
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

	//fmt.Println(ns)

	for i := 1; i < len(ns); i++ {
		used := make([]bool, N+1, N+1)
		//if !used[i] {
		dfs(ns, used, i, -1)
		//}
		if flag {
			break
		}
	}
	if flag {
		fmt.Println("YES")
		fmt.Println(st.Len())
		for {
			pval := st.pop()
			if pval == nil {
				break
			}
			fmt.Print(*pval, " ")
		}
		return
	}

	fmt.Println("NO")

	/* for {
		pval := st.pop()
		if pval == nil {
			break
		}
		fmt.Print(*pval, " ")
	}
	*/
}
