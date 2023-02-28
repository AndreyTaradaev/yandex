package main

/*
В этой задаче вам необходимо самостоятельно (не используя соответствующие классы и функции
стандартной библиотеки) организовать структуру данных Heap для хранения целых чисел,
над которой определены следующие операции: a) Insert(k) – добавить в Heap число k ; b) Extract достать
из Heap наибольшее число (удалив его при этом).
Формат ввода

В первой строке содержится количество команд N (1 ≤ N ≤ 100000),
 далее следуют N команд, каждая в своей строке. Команда может иметь формат: “0 <число>” или “1”,
  обозначающий, соответственно, операции Insert(<число>) и Extract. Гарантируется,
  что при выполенении команды Extract в структуре находится по крайней мере один элемент.

Формат вывода

Для каждой команды извлечения необходимо отдельной строкой вывести число,
 полученное при выполнении команды Extract.
*/

import (
	"bufio"
	"fmt"
	"os"
)

// очередь  Deque
type Heap struct {
	values []int
}

func (h *Heap) Add(el int) {
	h.values = append(h.values, el)
	pos := len(h.values) - 1
	for pos > 0 && h.values[pos] > h.values[(pos-1)/2] {
		h.values[pos], h.values[(pos-1)/2] = h.values[(pos-1)/2], h.values[pos]
		pos = (pos - 1) / 2
	}
}

func (h *Heap) Pop() *int {
	length := len(h.values)
	if length == 0 {
		return nil
	}
	ans := h.values[0]
	h.values[0] = h.values[length-1]
	pos := 0
	for (pos*2 + 1) < (length - 1) {
		min_son_index := pos*2 + 1
		if h.values[pos*2+2] > h.values[min_son_index] {
			min_son_index = pos*2 + 2
		}
		if h.values[pos] < h.values[min_son_index] {
			h.values[pos], h.values[min_son_index] = h.values[min_son_index], h.values[pos]
			pos = min_son_index
		} else {
			break
		}

	}
	h.values = h.values[:length-1]
	return &ans
}

func (h *Heap) Print() {
	if len(h.values) == 0 {
		return
	}
	level := 0
	fmt.Println(h.values[0])
	for {
		if 2*level+1 < len(h.values) {
			fmt.Print("(", h.values[level], ")", h.values[2*level+1], " ")
			if h.values[level] < h.values[2*level+1] {
				fmt.Print(" error ")
			}
		} else {
			break
		}
		if 2*level+2 < len(h.values) {
			fmt.Print(h.values[2*level+2])
			if h.values[level] < h.values[2*level+2] {
				fmt.Print(" error ")
			}
			fmt.Print("\n")
		} else {
			break
		}
		level++
	}
	fmt.Print("\n")
}

func createHeap() *Heap {
	var t Heap
	t.values = make([]int, 0, 20)
	return &t
}

func main() {
	hp := createHeap()
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

	for i := 0; i < N; i++ {
		var f int
		_, err = fmt.Fscan(rd, &f) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		switch f {
		case 0:
			
			_, err = fmt.Fscan(rd, &f) //   command
			if err != nil {
				fmt.Println(err)
				return
			}
			hp.Add(f)
		case 1:
			fmt.Println(*hp.Pop())
		default:
			break

		}

	}

}
