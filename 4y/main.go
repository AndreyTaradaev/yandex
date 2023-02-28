package main

/*

Петя и Вася — одноклассники и лучшие друзья, поэтому они во всём помогают друг другу. Завтра у них контрольная по математике, и учитель подготовил целых K вариантов заданий.

В классе стоит один ряд парт, за каждой из них (кроме, возможно, последней) на контрольной будут сидеть ровно два ученика. Ученики знают, что варианты будут раздаваться строго по порядку: правый относительно учителя ученик первой парты получит вариант 1, левый — вариант 2, правый ученик второй парты получит вариант 3 (если число вариантов больше двух) и т.д. Так как K может быть меньше чем число учеников N, то после варианта K снова выдаётся вариант 1. На последней парте в случае нечётного числа учеников используется только место 1.

Петя самым первым вошёл в класс и сел на своё любимое место. Вася вошёл следом и хочет получить такой же вариант, что и Петя, при этом сидя к нему как можно ближе. То есть между ними должно оказаться как можно меньше парт, а при наличии двух таких мест с равным расстоянием от Пети Вася сядет позади Пети, а не перед ним. Напишите программу, которая подскажет Васе, какой ряд и какое место (справа или слева от учителя) ему следует выбрать. Если же один и тот же вариант Вася с Петей писать не смогут, то выдайте одно число  - 1.
Формат ввода

В первой строке входных данных находится количество учеников в классе 2 ≤ N ≤ 109. Во второй строке — количество подготовленных для контрольной вариантов заданий 2 ≤ K ≤ N. В третьей строке — номер ряда, на который уже сел Петя, в четвёртой — цифра 1, если он сел на правое место, и 2, если на левое.

Формат вывода

Если Вася никак не сможет писать тот же вариант, что и Петя, то выведите  - 1. Если решение существует, то выведите два числа — номер ряда, на который следует сесть Васе, и 1, если ему надо сесть на правое место, или 2, если на левое. Разрешается использовать только первые N мест в порядке раздачи вариантов. */

import (
	"bufio"
	"fmt"
	"os"
	//"strconv"
)
func getPart(n int ) int {

return (n+1)/2
}

func getMesto(n int) int{	
return 2 - n%2
}




func main() {
	var N int     //  количество учеников
	var K int     //количечтво вариантов
	var nPart int // номер ряда(парта)
	var left int  // место левое или правое

	//получаем данные
	file, err := os.Open("input.txt") // open file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	rd := bufio.NewReader(file) // create reader

	_, err = fmt.Fscan(rd, &N) // scaN  count  Diego
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fscan(rd, &K)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fscan(rd, &nPart)
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fscan(rd, &left)
	if err != nil {
		fmt.Println(err)
		return
	}

	/* fmt.Println("количесто учеников ", N)
	fmt.Println("количество вариантов", K)
	fmt.Println("номер парты  ", nPart)
	fmt.Println("правый  ", left == 1,left) */
	// интересно какой вариант получит Петя?

	seat_number := nPart*2	- (left%2)	
	/* variant := seat_number%(K+1)
	fmt.Println(seat_number, variant)	
	fmt.Println(getPart(seat_number),getMesto(seat_number)) */
	switch  {
	case (seat_number-K <1 && seat_number+K>N) : fmt.Println("-1") //нет вариантов
	case (seat_number-K <1):      // переднего варианта нет
		fmt.Printf ("%d %d",  getPart(seat_number+K), getMesto(seat_number+K))
	case (seat_number+K>N) :	 // заднего варианта нет
		fmt.Printf ("%d %d", getPart(seat_number-K), getMesto(seat_number-K))
	default :
		if (nPart - getPart(seat_number-K) < getPart(seat_number+K)-nPart){
			fmt.Printf ("%d %d", getPart(seat_number-K), getMesto(seat_number-K))
		}	else{
			fmt.Printf ("%d %d", getPart(seat_number+K), getMesto(seat_number+K))
		}			
	}
}
