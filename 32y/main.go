package main

/*


Дан неориентированный невзвешенный граф. Необходимо посчитать количество его компонент связности и
вывести их.

Формат ввода

Во входном файле записано два числа N и M (0 < N ≤ 100000, 0 ≤ M ≤ 100000). В следующих M строках
записаны по два числа i и j (1 ≤ i, j ≤ N), которые означают, что вершины i и j соединены ребром.

Формат вывода

В первой строчке выходного файла выведите количество компонент связности. Далее выведите сами компоненты
связности в следующем формате: в первой строке количество вершин в компоненте, во второй - сами вершины в
произвольном порядке. .
*/

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)

type comp struct {
	vector int
	c      int
}

func dfs(graf [][]int,  v, comp int) {
	visited[v] = true
	for i := 0; i < len(graf[v]); i++ {
		if visited[graf[v][i].vector] == false {
			graf[v][i].c = comp
			dfs(graf, visited, graf[v][i].vector, comp)
		}
	}
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
	rd := bufio.NewReader(file)    // create reader
	_, err = fmt.Fscan(rd, &N, &M) //   command
	if err != nil {
		fmt.Println(err)
		return
	}
	ns := make([][]comp, N+1, N+1)
	for i := 0; i < len(ns); i++ {
		ns[i] = make([]comp, 0, 0)
	}

	for i := 0; i < M; i++ {
		var f, s int
		_, err = fmt.Fscan(rd, &f, &s) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		ns[f] = append(ns[f], comp{s, 0})
		if f != s {
			ns[s] = append(ns[s], comp{f, 0})
		}
	}
	//fmt.Println( ns)
	vis := make([]bool, N+1, N+1)
comp:=1
	for  i:=1 ; i<len(ns) {	
		dfs(ns,vis,1, comp)
		comp++
	}
	//dfs(ns, vis, 1)

	var ff string
	count := 0
	for i := 1; i < len(vis); i++ {

		if vis[i] {
			ff += fmt.Sprintf("%d ", i)
			count++
		}
	}
	fmt.Println(count)
	fmt.Println(ff)
}
