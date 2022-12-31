package test

import (
	"fmt"
	"strconv"
)

func CheckResult(number int) {
	res := Result(number)
	if res == true{
		fmt.Println("true")
	}else{
		fmt.Println("false")
	}
	return
}

func Result(number int) bool{
	tmp := number
	digits := len(strconv.Itoa(number))
	sum := 0
	div := 0
	for {
		if tmp <= 0 {
			break;
		}
		div = tmp % 10
		temp := 1
		for i := 0; i < digits; i++ {
			temp = temp * div
		}
		sum += temp
		tmp = tmp / 10
	}
	return number == sum
}