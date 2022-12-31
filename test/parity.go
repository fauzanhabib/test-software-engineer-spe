package test

import "fmt"

func ResultOutlier(number int){
	if FindOutlier(number){
		fmt.Println("(the only odd number)",number)
	}else{
		fmt.Println("(the only even number)",number)
	}
	return
}
func FindOutlier(num int) bool{
	result := false
	for num != 0{
		if num & 1 != 0 {
				result = !result
			}
			num = num >> 1
	}
	return result
}
