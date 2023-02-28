package main

/*


В игре в пьяницу карточная колода раздается поровну двум игрокам. Далее они вскрывают по одной верхней карте,
и тот, чья карта старше,
забирает себе обе вскрытые карты, которые кладутся под низ его колоды. Тот, кто остается без карт – проигрывает.
Для простоты будем считать, что все карты различны по номиналу, а также, что самая младшая карта побеждает
самую старшую карту ("шестерка берет туза"). Игрок, который забирает себе карты,
сначала кладет под низ своей колоды карту первого игрока,
затем карту второго игрока (то есть карта второго игрока оказывается внизу колоды).
Напишите программу, которая моделирует игру в пьяницу и определяет, кто выигрывает.
В игре участвует 10 карт, имеющих значения от 0 до 9,
большая карта побеждает меньшую, карта со значением 0 побеждает карту 9.
Формат ввода

Программа получает на вход две строки: первая строка содержит 5 чисел,
разделенных пробелами — номера карт первого игрока, вторая – аналогично
5 карт второго игрока. Карты перечислены сверху вниз, то есть каждая строка начинается с той карты,
которая будет открыта первой.
Формат вывода

Программа должна определить, кто выигрывает при данной раздаче, и вывести слово first или second,
после чего вывести количество ходов, сделанных до выигрыша.
Если на протяжении 106 ходов игра не заканчивается, программа должна вывести слово botva.
*/

import (
	"bufio"
	"fmt"
	"os"
)

// очередь  FIFO
type Queue struct {
	values []int
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
func (q *Queue) peek() *int {

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

func (q *Queue) Cls() {
	q.values = []int{}
}

func main() {

	file, err := os.Open("input.txt")
	if err != nil {
		return
	}

	defer file.Close()
	rd := bufio.NewReader(file) // create reader
	var first Queue
	var second Queue
	// fill first player

	for i := 0; i < 5; i++ {
		var N int
		_, err = fmt.Fscan(rd, &N) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		first.Enqueue(N)
	}

	for i := 0; i < 5; i++ {
		var N int
		_, err = fmt.Fscan(rd, &N) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		second.Enqueue(N)
	}

	// ходим
	var i int
	for i = 0; i < 1000000; i++ {
		card1, err := first.Dequeue()
		if err != nil {
			break
		}
		card2, err := second.Dequeue()
		if err != nil {
			break
		}
		switch {
		case (*card1 == 0 && *card2 == 9):
			first.Enqueue(*card1)
			first.Enqueue(*card2)
		case (*card2 == 0 && *card1 == 9):
			second.Enqueue(*card1)
			second.Enqueue(*card2)
		case *card1 > *card2: 
			first.Enqueue(*card1)
			first.Enqueue(*card2)
		case (*card2 >*card1 ):
			second.Enqueue(*card1)
			second.Enqueue(*card2)	
		}
	}
	
	 if  first.Len() == 0 {
		fmt.Println("second", i)
	 } else if  second.Len() == 0 {
		fmt.Println("first", i)
	 } else {
		fmt.Println("botva")	
	 }
	

}
