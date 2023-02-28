package main 
//import "fmt"

func main(){
var a string
a = "абц"
println(len(a)) // 6

for i:=0 ;i<len(a);i++{
println(a[i])
}


// выведет значение второго байта 99
// строки в Go неизменяемы, нельзя написать a[2] = 10

}