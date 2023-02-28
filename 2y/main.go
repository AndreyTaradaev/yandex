package main

import (
	"bufio"
	//"io"
	"fmt"
	"os"
	"strconv"
)

type easysyb struct {
	word  byte
	count int
}
func (e easysyb) String () string{
	return fmt.Sprintf("%s:%d", string( e.word ), e.count)	
}

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
	fmt.Println( count)

	isPrefix := true // read fine line
	for isPrefix {
		buf, isPrefix, err = rd.ReadLine()
		if err != nil {
			return
		}
		line = append(line, buf...)
	}
	// проход для группировки символов
	groupsym := make([]easysyb, 0, len(line))
	lastsym :=  line [0]
	lastpos :=0
	for ind, val := range line {
		if val != lastsym {
			groupsym = append(groupsym, easysyb{lastsym,ind-lastpos})
			lastsym = val
			lastpos= ind
		}
	}
	groupsym = append(groupsym, easysyb{lastsym,len(line)-lastpos})
	fmt.Println(string(line),  line , len (line))
	fmt.Println(groupsym ,len(groupsym))
	

	/* var lenFine int	=0
	//var symbol byte
	for i:=0; i<len(line)-count;i++{
		var ch = line[i]
		counttemp := count // для текущего символа количетво замен
		currentfine := 1  // длина красивой строки для текущего символа
			for j:=i+1; j<len(line);j++{
				if(ch ==line[j]){	// следующие символы  такие же		 замену проводить не надо увеличиваем длину красивой строки
					currentfine++
				} else{ // надо проводить замену счетчик замен уменьшается,  длина красивой строки увеличивается
					currentfine++
					counttemp --
				}
				if(counttemp==0){ // если  количество замен кончилось, сравниваем с текущим значением длины красивой строки и выходим из цикла
					if(currentfine> lenFine) {
						lenFine = currentfine
					//	symbol = line[i]
					}
					break
				}

			}

	}
	fmt.Println( lenFine) */

	//	line= append(line,buf... )

	/* if err != nil {

	} */

	//	isPrefix :=true
	//line := make([]byte,0,200000)

	//buf := make([]byte,0,200)

}
