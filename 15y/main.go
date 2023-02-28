package main

/*
 Лайнландия представляет из себя одномерный мир, являющийся прямой, на котором располагаются
 N городов, последовательно пронумерованных от 0 до N - 1 . Направление в сторону от первого города к нулевому названо
  западным, а в обратную — восточным.

Когда в Лайнландии неожиданно начался кризис, все были жители мира стали испытывать глубокое смятение.
 По всей Лайнландии стали ходить слухи, что на востоке живётся лучше, чем на западе.

Так и началось Великое Лайнландское переселение. Обитатели мира целыми городами отправились на восток,
покинув родные улицы, и двигались до тех пор, пока не приходили в город, в котором средняя цена проживания
была меньше, чем в родном.

Формат ввода
В первой строке дано одно число N (2≤N≤105) — количество городов в Лайнландии. Во второй строке дано
N чисел ai (0≤ai≤109) — средняя цена проживания в городах с нулевого по (N - 1)-ый соответственно.
Формат вывода
Для каждого города в порядке с нулевого по (N - 1)-ый выведите номер города,
в который переселятся его изначальные жители. Если жители города не остановятся в каком-либо другом городе,
отправившись в Восточное Бесконечное Ничто, выведите -1 .
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	//"strconv"
)

// очередь LIFO
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

func (q *Stack) peek() *int {
	s := len(q.values)
	if s == 0 {
		return nil
	}
	return &q.values[s-1]
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


func (q Stack) Len() int {
	return len(q.values)
}

func (q *Stack) String() string {
	var ret string
	ar := q.values
	for _, val := range ar {
		ret = ret + fmt.Sprintf("%#v ", val)
	}
	return ret
}

func (q *Stack) Cls() bool {
	q.values = []int{}
	return true
}

func main() {

	var N  int  //количество городов в Лайнландии.
	
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}

	defer file.Close()
	rd := bufio.NewReader(file) // create reader

	_, err = fmt.Fscan(rd, &N) //   command	
	if err != nil {
		fmt.Println(err)
		return
	}
	sl := make([]int,N,N)
	
	for i := 0; i < N; i++ {
	var f int	
		_, err = fmt.Fscan(rd, &f) //   command
	if err == io.EOF {		
		break		
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	sl[i] =f
	}
	//fmt.Println(N,sl)
	var st Stack
	result := make([]int,N,N)

	//st.Push(0)
	for  i:=0 ; i<N;i++ {
		for	st.Len()!=0 &&  sl[i] < sl[*st.peek()]	{
			val,_ := st.Pop()
			result[*val] = i
		}
	st.Push(i)	
	}
for _, v := range result {
var res int
if(v==0)	{
	res = -1
} else{
	res = v
}
fmt.Print(res," ")	
	
}
}
