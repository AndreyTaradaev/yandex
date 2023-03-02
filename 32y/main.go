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
произвольном порядке.
*/

import (
	"bufio"
	"fmt"
	"os"
	//"time"
	//	"sort"
)

func dfs(graf [][]int, visited []int, v, group int) {
	visited[v] = group
	for i := 0; i < len(graf[v]); i++ {
		if visited[graf[v][i]] != group {
			dfs(graf, visited, graf[v][i], group)
		}
	}
}

func main() {
	//var command string
	//start := time.Now()
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
	ns := make([][]int, N+1, N+1)
	for i := 0; i < len(ns); i++ {
		ns[i] = make([]int, 0, 0)
	}

	for i := 0; i < M; i++ {
		var f, s int
		_, err = fmt.Fscan(rd, &f, &s) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		ns[f] = append(ns[f], s)
		if f != s {
			ns[s] = append(ns[s], f)
		}
	}
	//fmt.Println( ns)
	vis := make([]int, N+1, N+1)
	for i := 1; i < len(vis); i++ {
		vis[i] = -1
	}
	comp := 1
	for i := 1; i < len(ns); i++ {
		if vis[i] == -1 {
			dfs(ns, vis, i, comp)
			comp++
		}

	}
	/* fmt.Println( ns)
	   fmt.Println(vis)
	*/

	m := make(map[int][]int)
	for i := 1; i < comp; i++ {
		m[i] = make([]int, 0, 2)
	}

	for i := 1; i < len(vis); i++ {
		m[vis[i]] = append(m[vis[i]], i)
	}
	//finish :=  time.Now().Sub(start)

	fmt.Println(comp - 1)
	for i := 1; i < comp; i++ {
		l := m[i]
		fmt.Println(len(l))
		for _, v := range l {
			fmt.Print(v, " ")
		}
		fmt.Println()
	}

	/* fmt.Println("finish",finish)
	   fmt.Println("finish print",time.Now().Sub(start)) */

}
