package main

import (
	"bufio"
	"io"
	"fmt"
	"os"
	"sort"
)

func print(mss map[rune]int, sss []string) {
	ss := make([]string, 0, len(sss))
	var str string
	for _, val := range sss {
		str += val
	}
	ss = append(ss, str)

	for {
		str = ""
		isEmply := true
		for _, val := range sss {
			rs := []rune(val)
			if mss[rs[0]] != 0 {
				str += "#"
				mss[rs[0]] = mss[rs[0]] - 1
				isEmply = false
			} else {
				str += " "
			}
		}
		if isEmply {
			break
		}
		ss = append(ss, str)
	}

	for i := len(ss) - 1; i >= 0; i-- {
		fmt.Println(ss[i])
	}

}

func main() {
	
	var buf  = make ([]byte,0,10000)
	file, err := os.Open("input.txt")
	if err != nil {
	  return
	}
	defer file.Close()
	rd := bufio.NewReader(file)

	
	for  {		
		var line  = make ([]byte,0,200)
		line ,_,err = rd.ReadLine()
		if( err == io.EOF ){
			break
		} else 	if(err != nil ){
			return
		}
		buf = append(buf, line...)
	}

	ms := make(map[rune]int) // карта подсчета символов в тексте
	ss := make([]string, 0, 0)
	
	// slice  rune for  input
	rs := []rune(string(buf[:]))

	for i := range rs { // обход среза рун для подсчета количества рун
		if rs[i] == ' ' || rs[i] == '\n'|| rs[i] == '\r' { // если пробел или новая строка пропускаем
			continue
		}
		//проверяем в карте наличие руны
		if _, ok := ms[rs[i]]; ok == false {
			ms[rs[i]] = 1
			// для сортировки
			ss = append(ss, string(rs[i]))
			continue
		}
		ms[rs[i]] = ms[rs[i]] + 1
	}
	sort.Strings(ss)
	print(ms, ss) 


}
