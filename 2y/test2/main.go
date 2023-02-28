package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {

	var count int
	line := make([]byte, 0, 2000)
	file, err := os.Open("input.txt")
	if err != nil {
		return
	}

	defer file.Close()
	rd := bufio.NewReader(file)

	buf, _, err := rd.ReadLine() // читаем число замен
	if err != nil {
		return
	}
	count, err = strconv.Atoi(string(buf))

	isPrefix := true // read fine line
	for isPrefix {
		buf, isPrefix, err = rd.ReadLine()
		if err != nil {
			return
		}
		line = append(line, buf...)
	}
	mb := make(map[byte]int)
	for _, val := range line {
		if _, ok := mb[val]; ok {
			mb[val] = mb[val] + 1
		} else {
			mb[val] = 1
		}
	}
	max := 0                 // длина красивой строки
	for key, _ := range mb { //  проход  по словарю
		//	key := byte(97)
		t, counts := count, 0 //  кол-во замен, длина красивой строки, текущая длина красивой строки
		var left int
		rigth := 0                               // левый и правый указатель
		for left = 0; left < len(line); left++ { //  проход по левому вказателю
			for rigth < len(line) && (t > 0 || key == line[rigth]) { //  проверяем правый указатель меньше длины строки и кол замен больше 0
				if key == line[rigth] { // если правый указатель  равен букве из словаря
					rigth++  // прав. указатель сдвигаем
					counts++ // увеличиваем тек длину красив строки
				} else { // // если правый указатель  не равен букве из словаря
					t-- // замена
					counts++
					rigth++
				}
			}
			if max < counts { // когда количество замен стало =0  проверяем длину кр. строки
				max = counts
				//fmt.Println(string(key) )
				//fmt.Println( string(key) ,max)
			}
			counts--
			if line[left] != key {
				t++
			}
		}
	}
	fmt.Println(max)
}
