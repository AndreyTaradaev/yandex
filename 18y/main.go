package main

/*
Научитесь пользоваться стандартной структурой данных deque для целых чисел.  Напишите программу,
содержащую описание дека и моделирующую работу дека, реализовав все указанные здесь методы.
Программа считывает последовательность команд и в зависимости от команды выполняет ту
или иную операцию. После выполнения каждой команды программа должна вывести одну строчку.
Возможные команды для программы:
push_front n
Добавить (положить) в начало дека новый элемент. Программа должна вывести ok.

push_back n
Добавить (положить) в конец дека новый элемент. Программа должна вывести ok.

pop_front
Извлечь из дека первый элемент. Программа должна вывести его значение.

pop_back
Извлечь из дека последний элемент. Программа должна вывести его значение.

front
Узнать значение первого элемента (не удаляя его). Программа должна вывести его значение.

back
Узнать значение последнего элемента (не удаляя его). Программа должна вывести его значение.

size
Вывести количество элементов в деке.

clear
Очистить дек (удалить из него все элементы) и вывести ok.

exit
Программа должна вывести bye и завершить работу.

Гарантируется, что количество элементов в деке в любой момент не превосходит 100.
 Перед исполнением операций pop_front, pop_back, front, back программа должна проверять,
 содержится ли в деке хотя бы один элемент. Если во входных данных встречается
 операция pop_front, pop_back, front, back, и при этом дек пуст,
 то программа должна вместо числового значения вывести строку error.
Формат ввода

Вводятся команды управления деком, по одной на строке.
Формат вывода

Требуется вывести протокол работы дека, по одному сообщению на строке
*/

import (
	"bufio"
	"fmt"
	"os"
)

// очередь  Deque
type Deque struct {
	values []int
}

// добавление в начало  очередь 
func (q *Deque) Push_Front(element int) {
	q.values = append(q.values, element) // Simply append to enqueue.
}

// добавление в конец  очередь 
func (q *Deque) Push_Back(element int) {
	tmp := make ([]int, 0,len(q.values)+1)
	tmp = append(tmp,element)
	tmp = append(tmp, q.values...)
	q.values = tmp // Simply append to enqueue.
}

// значение пераого в очереди
func (q *Deque) Front() (*int, error) {

	if len(q.values) == 0 {
		return nil, fmt.Errorf("Queue is emply!")
	}

	return &q.values[len(q.values)-1 ], nil
}

// значение пераого в очереди
func (q *Deque) Back() (*int, error) {

	if len(q.values) == 0 {
		return nil, fmt.Errorf("Queue is emply!")
	}
	return &q.values[0], nil
}

func (q *Deque) back() *int {

	if len(q.values) == 0 {
		return nil
	}
	return &q.values[0]
}

func (q *Deque) front() *int {

	if len(q.values) == 0 {
		return nil
	}
	return &q.values[len(q.values)-1]
}

func (q Deque) Len() int {
	return len(q.values)
}

func (q *Deque) Cls() {
	q.values = []int{}
}
 // удаление из начала очереди
func (q *Deque) Pop_Front() (*int, error) {

	if len(q.values) == 0 {
		return nil, fmt.Errorf("Queue is emply!")
	}
	element := q.values[len(q.values)-1]
	if len(q.values) == 1 {
		q.values = []int{}
		return &element, nil
	}
	q.values = q.values[0:len(q.values)-1]
	return &element, nil
}

// удаление из конца очереди
func (q *Deque) Pop_Back() (*int, error) {

	if len(q.values) == 0 {
		return nil, fmt.Errorf("Queue is emply!")
	}
	element := q.values[0]
	if len(q.values) == 1 {
		q.values = []int{}
		return &element, nil
	}
	q.values = q.values[1:]
	return &element, nil
}





func main() {
	var dq Deque 


 var command string
	file, err := os.Open("input1.txt")
	if err != nil {
		return
	}

	defer file.Close()
	rd := bufio.NewReader(file) // create reader

	

exit:
	for {
		_, err = fmt.Fscan(rd, &command) //   command
		if err != nil {
			fmt.Println(err)
			return
		}

		switch command {
		case "push_back":
			var inp int
			_, err = fmt.Fscan(rd, &inp) //   command
			if err != nil {
				fmt.Println(err)
				return
			}
			dq.Push_Back(inp)
			fmt.Println("ok")
		case "push_front":
			var inp int
			_, err = fmt.Fscan(rd, &inp) //   command
			if err != nil {
				fmt.Println(err)
				return
			}
			dq.Push_Front(inp)
			fmt.Println("ok")	
		case "pop_front":
			res, err := dq.Pop_Front()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(*res)
			}
		case "pop_back":
			res, err := dq.Pop_Back()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(*res)
			}	
		case "front":
			res, err := dq.Front()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(*res)
			}
		case "back":
			res, err := dq.Back()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(*res)
			}	
		case "size":
			
			fmt.Println(dq.Len())
		case "clear":
			dq.Cls()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			break exit
		default:
			fmt.Println("unknow")
		}

	} 
}
