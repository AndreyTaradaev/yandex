package main

import "fmt"

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	var rating = []int{1, 1, 3, 3, 4, 4, 6, 10, 11}

	bestsum := 15
	min := len(rating)
	for i := 0; i < len(rating); i++ {
		s := 0
		for j := i; j < len(rating); j++ {
			s = s + rating[j]
			if s > bestsum {
				if j-i+1 < min {

					min = j - i + 1
					sl := rating[i : j+1]
					fmt.Println(sl)
					fmt.Println(j, i, min, "\n")

				}
				break
			}
		}
	}
	fmt.Println(min, "O^2 end")
	min = len(rating)
	s := 0
	j:= 0
	for i := 0; i < len(rating); i++ {
		for  j<len(rating) &&  s<=bestsum {
			s = s + rating[j]
			 j++  
		}
			if s > bestsum {
				if j-i < min {
					min = j - i 
					sl := rating[i : j]
					fmt.Println(sl)
					fmt.Println(j-1, i, min, "\n")					

				}
			} else{
				break
				}
			s = s - rating[i]
		}		
	

	/*        ; // в переменной s подсчитываем текущую сумму
	        if s>k then begin
	            // writeln(j); // этот вывод пригодится дальше
	            if j-i+1>min then min:=j-i+1;
	            break;
	        end;
	    end;
	end; */
	/* 	fmt.Println("len ",len(rating))
	   	bestsum, nowsum, last := 0, 0, 0
	   	for ind:=0; ind <len(rating)-1;ind++ {
	   		fmt.Println("external step:", ind , ",last: ",last, ",nowsum: ", nowsum, ",bestSum: ",bestsum  )
	   		fmt.Println("external step:","first", rating[ind] , ",next" ,rating[ind+1],",last",rating[last]   )
	   		for (last < len(rating)) && (last == ind || rating[ind]+rating[ind+1] >= rating[last]) {
	   			nowsum += rating[last]
	   			last++
	   			fmt.Println("internal step: last",last ,",nowsum", nowsum)
	   		}
	   		bestsum = max(bestsum, nowsum)
	   		nowsum -= rating[ind]

	   		fmt.Println("external step:", "nowsum: ", nowsum, ",bestSum: ",bestsum ,"\n\n" )
	   	}
	   	fmt.Println("result:", "nowsum: ", nowsum, ",bestSum: ",bestsum  ) */
}
