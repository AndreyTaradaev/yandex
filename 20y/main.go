package main

/*


Отсортируйте данный массив. Используйте пирамидальную сортировку.

Формат ввода

Первая строка входных данных содержит количество элементов в массиве N, N ≤ 105. Далее задаются N целых чисел, 
не превосходящих по абсолютной величине 109.

Формат вывода

Выведите эти числа в порядке неубывания. 
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
	for pos > 0 && h.values[pos] <= h.values[(pos-1)/2] {
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
		if h.values[pos*2+2] <= h.values[min_son_index] {
			min_son_index = pos*2 + 2
		}
		if h.values[pos] >= h.values[min_son_index] {
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
		hp.Add(f) 	
	}
for   {
	if  f:= hp.Pop(); f !=nil{
		fmt.Print(*f," ")
	}else {
		break
	}
}
fmt.Print("\n")
}
