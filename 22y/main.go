package main

/*

У одного из студентов в комнате живёт кузнечик, который очень любит прыгать по клетчатой одномерной доске. Длина доски — N клеток. К его сожалению, он умеет прыгать только на 1, 2, …, k клеток вперёд.

Однажды студентам стало интересно, сколькими способами кузнечик может допрыгать из первой клетки до последней. Помогите им ответить на этот вопрос.
Формат ввода

В первой и единственной строке входного файла записано два целых числа — N и k .
Формат вывода

Выведите одно число — количество способов, которыми кузнечик может допрыгать из первой клетки до последней.
*/

import (
	"bufio"
	"fmt"
	"os"
)

func f(n int) int {
	switch n {
	case 1:
		return 2
	case 2:
		return 4
	case 3:
		return 7
	}
	return f(n-1) + f(n-2) + f(n-3)
}

func main() {
	//var command string
	var N,k int
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
	_, err = fmt.Fscan(rd, &k)  //   command
	if err != nil {
		fmt.Println(err)
		return
	}

	dp := make([]int, N+1, N+1)
dp[1] = 1

for  i := 2; i <= N; i++ {
	for  j := i-1; j >= i-k && j >= 1; j-- {
			dp[i] += dp[j];
	}
}
	fmt.Println(dp[N] ,dp)
}
