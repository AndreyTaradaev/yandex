package main

/*




Во время контрольной работы профессор Флойд заметил, что некоторые студенты обмениваются записками. 
Сначала он хотел поставить им всем двойки, но в тот день профессор был добрым, а потому
 решил разделить студентов на две группы: списывающих и дающих списывать, и поставить двойки только первым.

У профессора записаны все пары студентов, обменявшихся записками. Требуется определить, сможет
 ли он разделить студентов на две группы так, чтобы любой обмен записками осуществлялся от студента одной 
 группы студенту другой группы.
Формат ввода

В первой строке находятся два числа N и M — количество студентов и количество пар студентов, обменивающихся 
записками (1 ≤ N ≤ 102, 0 ≤ M ≤ N(N−1)/2).

Далее в M строках расположены описания пар студентов: два числа, соответствующие номерам студентов, 
обменивающихся записками (нумерация студентов идёт с 1). Каждая пара студентов перечислена не более 
одного раза.

Формат вывода

Необходимо вывести ответ на задачу профессора Флойда. Если возможно разделить студентов на две группы 
- выведите YES; иначе выведите NO. 
*/

import (
	"bufio"
	"fmt"
	"os"
	//	"sort"
)
var Error int = 0

func dfs (graf [][]int,Used [] int , v,color int )  {
	if(Error !=0 ){
		return
	}
	Used[v] = color
	for i := 0; i < len(graf[v]) ; i++ {
		if(Used[graf[v][i]] == 0 )	{
			dfs(graf,Used, graf[v][i],3-color )
		}
		if	(Used[v] == Used[graf[v][i]]) {
			Error =1
			return
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
		if( f!=s){
			ns[s] =  append(ns[s], f )
		}
	}
	sl :=make([]int,N+1,N+1)
 //fmt.Println( ns)
 for i := 1; i < len(ns); i++ {
	if (sl[i]==0){
	dfs(ns,sl,i,1)
	}
	if (Error==1){
		fmt.Println("NO")
		return
	 }	
	
 }
 
 
 fmt.Println("YES")

}
