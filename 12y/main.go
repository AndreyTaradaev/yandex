package main
/*
Рассмотрим последовательность, состоящую из круглых, квадратных и фигурных скобок. Программа дожна определить, является ли данная скобочная последовательность правильной. Пустая последовательность явлется правильной. Если A – правильная, то последовательности (A), [A], {A} – правильные. Если A и B – правильные последовательности, то последовательность AB – правильная.
Формат ввода

В единственной строке записана скобочная последовательность, содержащая не более 100000 скобок.
Формат вывода

Если данная последовательность правильная, то программа должна вывести строку yes, иначе строку no. 
*/
import (
	"bufio"
	"fmt"
	"os"
)

func checkBrackets(input string) bool {
	stack := make([]rune, 0)
	for _, v := range input {
		switch v {
		case '{', '[', '(':
			stack = append(stack, v)
		case '}', ']', ')':
			if len(stack) == 0 {
				return false
			}
			last := stack[len(stack)-1]
			if (v == '}' && last != '{') || (v == ']' && last != '[') || (v == ')' && last != '(') {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

func main() {

	var line string
	file, err := os.Open("input1.txt")
	if err != nil {
		return
	}

	defer file.Close()
	rd := bufio.NewReader(file)      // create reader
	_, err = fmt.Fscan(rd, &line) //   command
	if err != nil {
		fmt.Println(err)
		return
	}
if(checkBrackets(line)) {
	fmt.Println("yes")
	return
}
fmt.Println("no")
}
