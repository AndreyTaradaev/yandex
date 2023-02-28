package main

/*


Научитесь пользоваться стандартной структурой данных queue для целых чисел. Напишите программу, содержащую описание очереди и моделирующую работу очереди, реализовав все указанные здесь методы. 

Программа считывает последовательность команд и в зависимости от команды выполняет ту или иную операцию. После выполнения каждой команды программа должна вывести одну строчку.

Возможные команды для программы:

push n
Добавить в очередь число n (значение n задается после команды). Программа должна вывести ok.

pop
Удалить из очереди первый элемент. Программа должна вывести его значение.

front
Программа должна вывести значение первого элемента, не удаляя его из очереди.

size
Программа должна вывести количество элементов в очереди.

clear
Программа должна очистить очередь и вывести ok.

exit
Программа должна вывести bye и завершить работу.

Перед исполнением операций front и pop программа должна проверять, содержится ли в очереди хотя бы один элемент. Если во входных данных встречается операция front или pop, и при этом очередь пуста, то программа должна вместо числового значения вывести строку error.

*/

import (
	"bufio"
	"fmt"
	"os"
)

// очередь  FIFO
type Queue struct {
	values [] int
}

// добавление в очередь
func (q *Queue) Enqueue(element int) {
	q.values = append(q.values, element) // Simply append to enqueue.
}

// значение пераого в очереди
func (q *Queue) Peek() (*int, error) {

	if len(q.values) == 0 {
		return nil, fmt.Errorf("Queue is emply!")
	}
	return &q.values[0], nil
}
func (q *Queue) peek() (*int) {

	if len(q.values) == 0 {
		return nil
	}
	return &q.values[0]
}


// удаление из очереди
func (q *Queue) Dequeue() (*int, error) {

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

func (q Queue) Len() int {
	return len(q.values)
}

func (q *Queue) Cls()  {
	  q.values = []int{}
}


func main() {

	var command string
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}

	defer file.Close()
	rd := bufio.NewReader(file) // create reader

	var st Queue

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
			st.Enqueue(inp)
			fmt.Println("ok")
		case "pop":
			res, err := st.Dequeue()
			if err != nil {
				fmt.Println("error")
			} else {
				fmt.Println(*res)
			}
		case "front":
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
