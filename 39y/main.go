package main

/*
Пещера представлена кубом, разбитым на N частей по каждому измерению (то есть на N3 кубических клеток). Каждая клетка может быть или пустой, или полностью заполненной камнем. Исходя из положения спелеолога в пещере, требуется найти, какое минимальное количество перемещений по клеткам ему требуется, чтобы выбраться на поверхность. Переходить из клетки в клетку можно, только если они обе свободны и имеют общую грань.
Формат ввода

В первой строке содержится число N (1 ≤ N ≤ 30). Далее следует N блоков. Блок состоит из пустой строки и N строк по N символов: # - обозначает клетку, заполненную камнями, точка - свободную клетку. Начальное положение спелеолога обозначено заглавной буквой S. Первый блок представляет верхний уровень пещеры, достижение любой свободной его клетки означает выход на поверхность. Выход на поверхность всегда возможен.

Формат вывода

Вывести одно число - длину пути до поверхности.
*/

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)

type Point struct {
	x, y, z int
}
type Queue struct {
	values []Point
}

// добавление в очередь
func (q *Queue) Enqueue(element Point) {

	q.values = append(q.values, element) // Simply append to enqueue.
}

// значение пераого в очереди
func (q *Queue) Peek() (*Point, error) {

	if len(q.values) == 0 {
		return nil, fmt.Errorf("Queue is emply!")
	}
	return &q.values[0], nil
}

// удаление из очереди
func (q *Queue) Dequeue() (*Point, error) {

	if len(q.values) == 0 {
		return nil, fmt.Errorf("Queue is emply!")
	}
	element := q.values[0]
	if len(q.values) == 1 {
		q.values = []Point{}
		return &element, nil
	}
	q.values = q.values[1:]
	return &element, nil
}

func (q Queue) Len() int {
	return len(q.values)
}

func (q Queue) Get() []Point {
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

var queue Queue

func BFS(graf [][][]int, start Point) int {

	var step = -1
	// очередь
	var q Queue //   очередь вершин
	q.Enqueue(start)

	for q.Len() != 0 {
		v, _ := q.Peek()
		q.Dequeue()
		if v.z == 0 {
			step = graf[v.z][v.x][v.y]
			break
		}
		for i := 0; i < 6; i++ {
			x, y, z := GetMove(i)
			x += v.x
			y += v.y
			z += v.z
			if correct(x, y, z, len(graf)) && graf[z][x][y] == 0 {
				graf[z][x][y] = graf[v.z][v.x][v.y] + 1
				q.Enqueue(Point{x, y, z})
			}
		}

		
	}
	return step - graf[start.z][start.x][start.y]
}
func PrintGraf(Graf [][][]int) {
	for ind, v := range Graf {
		fmt.Println(ind, "\t", v)
	}
}

func correct(x, y, z, size int) bool {
	if x < 0 || y < 0 || z < 0 {
		return false
	}
	if x >= size || y >= size || z >= size {
		return false
	}
	return true
}

func GetMove(i int) (int, int, int) {
	X := [...]int{0, 0, 0, 0, -1, 1}
	Y := [...]int{0, 0, -1, 1, 0, 0}
	Z := [...]int{-1, 1, 0, 0, 0, 0}
	return X[i], Y[i], Z[i]
}

func main() {

	var N int // В первой строке содержится число N (1 ≤ N ≤ 30)

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
	spl := make([][][]int, N, N)

	for i := 0; i < N; i++ {
		spl[i] = make([][]int, N, N)
		for j := 0; j < N; j++ {
			spl[i][j] = make([]int, N, N)
		}
	}
	//PrintGraf(spl)
	var base Point
	// read block
	for z := 0; z < N; z++ {

		//fmt.Println()

		for j := 0; j < N; j++ {
			var in string
			_, err = fmt.Fscan(rd, &in) //   command
			if err != nil {
				fmt.Println(err)
				return
			}

			for i := 0; i < N; i++ {

				r := in[i]
				//	fmt.Println("i=",i,"j=",j,"z=",z,"   " ,  r, string(r))
				if r == '#' {
					spl[z][i][j] = 1
				} else if r == 'S' {
					spl[z][i][j] = 2
					base = Point{i, j, z}
					//fmt.Println(base)
				}

			}
		}

	}

	//PrintGraf(spl)
	

	fmt.Println(BFS(spl, base))
}
