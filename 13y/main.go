package main

/*
 В постфиксной записи (или обратной польской записи) операция записывается после двух операндов. Например, сумма двух чисел A и B записывается как A B +. Запись B C + D * обозначает привычное нам (B + C) * D, а запись A B C + D * + означает A + (B + C) * D. Достоинство постфиксной записи в том, что она не требует скобок и дополнительных соглашений о приоритете операторов для своего чтения.
Формат ввода
В единственной строке записано выражение в постфиксной записи, содержащее цифры и операции +, -, *. Цифры и операции разделяются пробелами. В конце строки может быть произвольное количество пробелов.
Формат вывода
Необходимо вывести значение записанного выражения.
*/

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
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

func add(i, j int) int {
	return i + j
}
func sub(i, j int) int {
	return i - j
}

func mul(i, j int) int {
	return i * j
}

func main() {

	var command string
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}

	defer file.Close()
	rd := bufio.NewReader(file) // create reader
	var st Stack
	for {
		_, err = fmt.Fscan(rd, &command) //   command
		if err == io.EOF {
			res, _ := st.Pop()
			fmt.Println(*res)
			return
		}
		if err != nil {
			fmt.Println(err)
			return
		}

		if s, err := strconv.Atoi(command); err == nil {
			st.Push(s)
		} else {
			var f func(i, j int) int
			switch command {
			case "+":
				f = add
			case "-":
				f = sub
			case "*":
				f = mul
			}
			j, _ := st.Pop()
			i, _ := st.Pop()
			st.Push(f(*i, *j))
		}
	}

}
