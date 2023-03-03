package main

/*
Дан ориентированный граф. Необходимо построить топологическую сортировку.
Формат ввода

В первой строке входного файла два натуральных числа N и M (1 ≤ N, M ≤ 100 000) — количество вершин и рёбер
 в графе соответственно. Далее в M строках перечислены рёбра графа. Каждое ребро задаётся парой чисел — 
 номерами начальной и конечной вершин соответственно.
Формат вывода

Выведите любую топологическую сортировку графа в виде последовательности номеров вершин 
(перестановка чисел от 1 до N). Если топологическую сортировку графа построить невозможно, выведите -1. 
*/

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
	pval,err := q.Pop()
	if (err!=nil){
		return nil
	}
	return pval
}

func (q Stack) Len() int {
	return len(q.values)
}

var Error bool =false
var st Stack 

func dfs (graf [][]int,Used [] int , v int )  {
	if Error{
		return
	}

	Used[v] = 1 //При входе в вершину окрашиваем ее в серый цвет
	for i := 0; i < len(graf[v]) ; i++ {

		if(Used[graf[v][i]] == 0 )	{
			dfs(graf,Used, graf[v][i])
		} else  if (Used[graf[v][i]] == 1){
			Error = true
			return
		}
	}
	st.Push(v)	
	Used[v] = 2
}



func main() {
	//var command string
	var N, M int // вершины  ребра
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file) // create reader
	_, err = fmt.Fscan(rd, &N,&M)  //   command
	if err != nil {
		fmt.Println(err)
		return
	}
	ns := make([] []int, N+1, N+1)
	for i := 0; i < len(ns); i++ {
		ns[i] = make([]int,0,10)
	}

	for i := 0; i < M; i++ {
		var f,s int
		_, err = fmt.Fscan(rd, &f,&s) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		ns[f] =  append(ns[f], s )
		/* if( f!=s){
			ns[s] =  append(ns[s], f )
		} */
	}
	used :=make([] int ,N+1,N+1)

	for i := 1; i < len(ns); i++ {
		if (used[i]== 0){
			dfs(ns,used,i)
		}		
		if (Error){
			fmt.Println("-1")
			return
		}
	}

	for {
		pval := st.pop()
		if(pval==nil){
			break
		}
		fmt.Print( *pval," ")
	}

}
