package main

import (
	"bufio"
	"fmt"
	"os"
)

/*


К тупику со стороны пути 1 (см. рисунок) подъехал поезд. Разрешается отцепить от поезда один или сразу
несколько первых вагонов и завезти их в тупик (при желании, можно даже завезти в тупик сразу весь поезд).
 После этого часть из этих вагонов вывезти в сторону пути 2. После этого можно завезти в тупик еще несколько
 вагонов и снова часть оказавшихся вагонов вывезти в сторону пути 2. И так далее (так, что каждый вагон может
	 лишь один раз заехать с пути 1 в тупик, а затем один раз выехать из тупика на путь 2). Заезжать в тупик с пути
	 2 или выезжать из тупика на путь 1 запрещается. Нельзя с пути 1 попасть на путь 2, не заезжая в тупик.

Известно, в каком порядке изначально идут вагоны поезда. Требуется с помощью указанных операций сделать так,
 чтобы вагоны поезда шли по порядку (сначала первый, потом второй и т.д., считая от головы поезда, едущего по пути 2
	 в сторону от тупика). Напишите программу, определяющую, можно ли это сделать.
Формат ввода

Вводится число N — количество вагонов в поезде (1 ≤ N ≤ 100). Дальше идут номера вагонов в порядке от головы поезда,
 едущего по пути 1 в сторону тупика. Вагоны пронумерованы натуральными числами от 1 до N, каждое из которых встречается
 ровно один раз.

Формат вывода

Если сделать так, чтобы вагоны шли в порядке от 1 до N, считая от головы поезда, когда поезд поедет по пути 2 из тупика,
 можно, выведите сообщение YES, если это сделать нельзя, выведите NO.
*/
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

	var N int // N — количество вагонов в поезде (1 ≤ N ≤ 100)

	file, err := os.Open("input.txt")
	if err != nil {
		return
	}

	defer file.Close()
	rd := bufio.NewReader(file) // create reader

	_, err = fmt.Fscan(rd, &N) //   command
	if err != nil {
		fmt.Println(err)
		return
	}
	vagons := make([]int, N, N) /* way[next] ... way[N-1] – вагоны на пути 1; когда путь пуст, next=N */

	for i := 0; i < N; i++ {
		var f int
		_, err = fmt.Fscan(rd, &f) //   command
		if err != nil {
			fmt.Println(err)
			return
		}
		vagons[i] = f
	}
	//fmt.Println(vagons)
	var next int = 0     // текущий вагон  с пути 1
	var expected int = 1 /* номер вагона, который ждут следующим на пути 2
	(значение N+1 будет означать, что все вагоны уже успешно выведены) */
	var st Stack
	peek := func() int {
		ret, err := st.Peek()
		if err != nil {
			return 0
		}
		return *ret
	}

	for next != N || (st.Len() != 0 && expected == peek()) {
		if st.Len() != 0 && expected == peek() {
			st.Pop()
			expected++
		} else {
			st.Push(vagons[next])
			next++
		}

	}
	if expected == N+1 {
		fmt.Println("YES")
		return
	}
	fmt.Println("NO")

	//
}
