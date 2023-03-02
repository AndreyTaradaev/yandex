package main

/*


Дан неориентированный граф, возможно, с петлями и кратными ребрами. Необходимо построить
компоненту связности, содержащую первую вершину.
Формат ввода

В первой строке записаны два целых числа N (1 ≤ N ≤ 103) и M (0 ≤ M ≤ 5 * 105) — количество вершин и 
ребер в графе. В последующих M строках перечислены ребра — пары чисел, определяющие номера вершин, которые
 соединяют ребра.
Формат вывода

В первую строку выходного файла выведите число K — количество вершин в компоненте связности.
 Во вторую строку выведите K целых чисел — вершины компоненты связности, перечисленные в порядке
  возрастания номеров. 
*/

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)

func dfs (graf [][]int,visited  [] bool, v int )  {
	visited[v] = true
	for i := 0; i < len(graf[v]) ; i++ {
		if(visited[graf[v][i]] == false )	{
			dfs(graf,visited, graf[v][i] )
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
	rd := bufio.NewReader(file) // create reader
	_, err = fmt.Fscan(rd, &N,&M)  //   command
	if err != nil {
		fmt.Println(err)
		return
	}
	ns := make([] []int, N+1, N+1)
	for i := 0; i < len(ns); i++ {
		ns[i] = make([]int,0,0)
	}

	for i := 0; i < M; i++ {
		var f,s int
		_, err = fmt.Fscan(rd, &f,&s) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		ns[f] =  append(ns[f], s )
		if( f!=s){
			ns[s] =  append(ns[s], f )
		}
	}
 //fmt.Println( ns)
 vis := make([] bool, N+1,N+1) 	
dfs(ns,vis,1)

 var ff string
 count:=0
for i := 1; i < len(vis); i++ {

	if(vis[i]){
		ff += fmt.Sprintf("%d ",i)
		count++
	}	
}	
fmt.Println(count)
fmt.Println(ff)
}
