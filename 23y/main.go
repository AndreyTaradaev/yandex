package main

/*


Имеется калькулятор, который выполняет следующие операции:

    умножить число X на 2;
    умножить число X на 3;
    прибавить к числу X единицу.

Определите, какое наименьшее количество операций требуется, чтобы получить из числа 1 число N.
Формат ввода

Во входном файле написано натуральное число N, не превосходящее 106.
Формат вывода

В первой строке выходного файла выведите минимальное количество операций.
 Во второй строке выведите числа, последовательно получающиеся при выполнении операций. 
Первое из них должно быть равно 1, а последнее N. Если решений несколько, выведите любое. 
*/

import (
	"bufio"
	"fmt"
	"os"
)

func min(i,j int) int {
	if i>j {
		return j
	}
	return i
}

func f(n int) int {
	if n==1{
		return 0
	}
	v:= f(n-1)
	if n%3==0 {
		v = min (v,f(n/3))
	}
	if n%2==0 {
		v = min (v,f(n/2))
	}
	return v+1
}

func main() {
	//var command string
	var N int
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

//	fmt.Println(f(N))
	sl := make ([]int,N+1,N+1 )
	for i := 2; i < N+1; i++ {
		v:= sl[i-1]
		if (i%2==0){
			v = min(v,sl[i/2])			
		}
		if (i%3==0){
			v = min(v,sl[i/3])
		}
		sl[i]=v+1		
	}
fmt.Println(sl[N])	
sl1 := make ([]int,1,N+1)
sl1 [0] =N
for i := N ; i >1 ; {
	if(i%3==0 && sl[i]==sl[i/3]+1 ){		
		i =i/3
		sl1 =append(sl1,i)
		continue
	}

	if( i%2==0 && sl[i]==sl[i/2]+1 ){
		i= i/2
		sl1 =append(sl1,i)
		continue
	}



	i--
	sl1 =append(sl1,i)
}
for i := len( sl1); i>0; i-- {
fmt.Print(sl1[i-1]," ")		
}

}
