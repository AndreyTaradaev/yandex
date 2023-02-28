package main

/*


Научитесь пользоваться стандартной структурой данных stack для целых чисел. Напишите программу, содержащую описание стека и моделирующую работу стека, реализовав все указанные здесь методы. Программа считывает последовательность команд и в зависимости от команды выполняет ту или иную операцию. После выполнения каждой команды программа должна вывести одну строчку. Возможные команды для программы:

push n
Добавить в стек число n (значение n задается после команды). Программа должна вывести ok.

pop
Удалить из стека последний элемент. Программа должна вывести его значение.

back
Программа должна вывести значение последнего элемента, не удаляя его из стека.

size
Программа должна вывести количество элементов в стеке.

clear
Программа должна очистить стек и вывести ok.

exit
Программа должна вывести bye и завершить работу.

Перед исполнением операций back и pop программа должна проверять, содержится ли в стеке хотя бы один элемент. Если во входных данных встречается операция back или pop, и при этом стек пуст, то программа должна вместо числового значения вывести строку error.
Формат ввода

Вводятся команды управления стеком, по одной на строке
Формат вывода

Программа должна вывести протокол работы стека, по одному сообщению на строке
*/

import (
	"bufio"
	"fmt"
	"os"
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

func (q *Stack) Get() []int {
	return q.values
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

	var command string
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}

	defer file.Close()
	rd := bufio.NewReader(file) // create reader

	var st Stack

exit:
	for {
		_, err = fmt.Fscan(rd, &command) //   command
		if err != nil {
			fmt.Println(err)
			return
		}

		switch command {
		case "push":
			var inp int
			_, err = fmt.Fscan(rd, &inp) //   command
			if err != nil {
				fmt.Println(err)
				return
			}
			st.Push(inp)
			fmt.Println("ok")
		case "pop":
			res, err := st.Pop()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(*res)
			}
		case "back":
			res, err := st.Peek()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(*res)
			}
		case "size":
			lenst := st.Len()
			fmt.Println(lenst)
		case "clear":
			st.Cls()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			break exit
		default:
			fmt.Println("unknow")
		}

	}
}
