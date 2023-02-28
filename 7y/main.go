package main

/*


Для того чтобы компьютеры поддерживали актуальное время, они могут обращаться к серверам точного времени SNTP (Simple Network Time Protocol). К сожалению, компьютер не может просто получить время у сервера, потому что информация по сети передаётся не мгновенно: пока сообщение с текущим временем дойдёт до компьютера, оно потеряет свою актуальность. Протокол взаимодействия клиента (компьютера, запрашивающего точное время) и сервера (компьютера, выдающего точное время) выглядит следующим образом:

1. Клиент отправляет запрос на сервер и запоминает время отправления A (по клиентскому времени).

2. Сервер получает запрос в момент времени B (по точному серверному времени) и отправляет клиенту сообщение, содержащее время B.

3. Клиент получает ответ на свой запрос в момент времени C (по клиентскому времени) и запоминает его. Теперь клиент, из предположения, что сетевые задержки при передаче сообщений от клиента серверу и от сервера клиенту одинаковы, может определить и установить себе точное время, используя известные значения A, B, C.

Вам предстоит реализовать алгоритм, с точностью до секунды определяющий точное время для установки на клиенте по известным A, B и C. При необходимости округлите результат до целого числа секунд по правилам арифметики (в меньшую сторону, если дробная часть числа меньше 1/2, иначе в большую сторону).

Возможно, что, пока клиент ожидал ответа, по клиентскому времени успели наступить новые сутки, однако известно, что между отправкой клиентом запроса и получением ответа от сервера прошло менее 24 часов.
Формат ввода

Программа получает на вход три временные метки A, B, C, по одной в каждой строке. Все временные метки представлены в формате «hh:mm:ss», где «hh» – это часы, «mm» – минуты, «ss» – секунды. Часы, минуты и секунды записываются ровно двумя цифрами каждое (возможно, с дополнительными нулями в начале числа).

Формат вывода

Программа должна вывести одну временную метку в формате, описанном во входных данных, – вычисленное точное время для установки на клиенте. В выводе не должно быть пробелов, пустых строк в начале вывода.
*/

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var ClSend, servSend, clientget string // получает на вход три временные метки A, B, C, по одной в каждой строке. Все временные метки представлены в формате «hh:mm:ss», где «hh» – это часы, «mm» – минуты, «ss» – секунды. Часы, минуты и секунды записываются ровно двумя цифрами каждое (возможно, с дополнительными нулями в начале числа).

	//и конечного секторов раздела (1 ≤ ai ≤ bi ≤ M).
	//получаем данные
	file, err := os.Open("input.txt") // open file
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	rd := bufio.NewReader(file)     // create reader
	_, err = fmt.Fscan(rd, &ClSend) //  время запроса клиента
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = fmt.Fscan(rd, &servSend) //  время когда сервер отправил ответ
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = fmt.Fscan(rd, &clientget) //   время когда клиент получил ответ
	if err != nil {
		fmt.Println(err)
		return
	}

	//var time1 time.Time
	t1, _ := time.Parse("15:04:05", ClSend)
	t2, _ := time.Parse("15:04:05", servSend)
	t3, _ := time.Parse("15:04:05", clientget)
	/* 	if(t2.Second() < t1.Second()){
	   		t2 = t2.AddDate(0,0,1)
	   	}
	*/
	if t3.Unix() < t1.Unix() {
		t3 = t3.AddDate(0, 0, 1)
	}
	d1 := t3.Sub(t1).Seconds()

	d1 = d1/2 + 0.5
	d3 := int(t2.Sub(t3).Seconds())
	dur := time.Duration(int(d1) + d3)
	t3 = t3.Add(time.Second * dur)
	fmt.Println(t3.Format("15:04:05"))

}
