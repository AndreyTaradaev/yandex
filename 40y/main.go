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

func BFS(graf [][]int, startnode, endnode int) []int {
	if startnode < 0 || startnode >= len(graf) {
		return nil
	}

	if endnode < 0 || endnode >= len(graf) {
		return nil
	}
	// очередь
	var q Queue[int] //   очередь вершин
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
		size = len(graf[*v]) //получаем количество связей у вершины
		for i := 0; i < size; i++ {
			to := graf[*v][i]
			if !used[to] { //Если вершина не посещена
				used[to] = true //посещаем ее
				q.Enqueue(to)   //и добавляем к концу очереди
				//if (to.line!=)
				d[to] = d[*v] + 1 //Считаем расстояние до вершины
				p[to] = *v        //Запоминаем предка
			}
		}
	}

	if !used[endnode] {

		return nil
	} else {
		path := []int{}
		for v := endnode; v != -1; v = p[v] {
			//for v := startnode; v != endnode; v = p[v] {
			path = append(path, v)
		}
		return path
	}
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
	// попытка  сформировать граф по станциям не увенчалась успехом
	// будем по другому, по линиям
	for i, v := range ns {
		fmt.Println(i, v)
	}

	// мап для поиска линии по станции

	m := make(map[int][]int) // ключ  станция , данные массив линий
	for i := 0; i < len(ns); i++ {
		for j := 0; j < len(ns[i]); j++ {

			m[ns[i][j]] = append(m[ns[i][j]], i)
		}
	}
	fmt.Println("\n")
	for key, v := range m {
		fmt.Println(key, v)
	}

	// спиок смежности линий

	line := make([][]int, M, M)

	for i := 0; i < M; i++ {
		line[i] = make([]int, 0, M)
	}
	// заполняем список
	isExist := func(ar []int, val int) bool {
		for i := 0; i < len(ar); i++ {
			if ar[i] == val {
				return true
			}
		}
		return false
	}

	for _, v := range m {
		for i := 1; i < len(v); i++ {
			if !isExist(line[v[i]], v[i-1]) {
				line[v[i]] = append(line[v[i]], v[i-1])
				line[v[i-1]] = append(line[v[i-1]], v[i])
			}
		}
	}

	fmt.Println("\n")
	for i, v := range line {
		fmt.Println(i, v)
	}

	getLine := func(station int) []int { // к какой линии принадлежит  станция
		return m[station]
	}

	firstLine := getLine(first)
	endline := getLine(end)
	min :=int(10e9)

	for i := 0; i < len(firstLine); i++ {
		for j := 0; j < len(endline); j++ {
		path := BFS(line,firstLine[i],endline[j])	
			if( path == nil){
				fmt.Println("-1")
				return
			}
			if len(path)<min{
				min = len(path)

			}
			
		}
		
	}
	fmt.Println(min)
	//path :=  BFS(line, startnode, endnode int) []int {

}

/*
   	mas := make([][]int, N+1, N+1)
   	for i := 0; i < len(mas); i++ {
   		mas[i] = make([]int, 0, N+1)
   	}




   	for i := 0; i < len(ns); i++ {
   		for j := 0; j < len(ns[i])-1; j++ {
   				mas[ns[i][j]] = append(mas[ns[i][j]], ns[i][j+1])
   				mas[ns[i][j+1]] = append(mas[ns[i][j+1]], ns[i][j])
   		}
   	}


   	  way :=  BFS(mas,first, end)
   	  if(way==nil){
   		fmt.Println("-1")
   		return
   	  }
   	  fmt.Println(way)
   	isline := func(station , l int ) bool { //проверяет есть ли у этой линии станция

   			for j := 0; j < len(ns[l]); j++ {
   				if(ns[l][j] == station ) {
   					return true
   				}
   			}
   		return false
   	}
   getLine:= func (first,end int) int{
   	for i := 0; i < len(ns); i++ {
   		var l1, l2 int = -1, -2
   		for j := 0; j < len(ns[i]); j++ {
   			if ns[i][j] == first  {
   				l1 = i
   			}
   			if ns[i][j] == end {
   				l2 = i
   			}
   			if	l1 == l2 {
   				return i+1
   			}
   		}
   	}
   return -1
   }

   //fmt.Println(ns)
   jump := 0
   l :=0

   	for i:=1 ;i<len( way);i++ {
   		if i==1 {
   			l = getLine(way[i-1],way[i] )
   			continue
   		}
   		l3 :=  getLine(way[i-1],way[i] )
   		if ( l != l3){
   			jump++
   			l= l3
   		}
   	}
   fmt.Println(jump) */
