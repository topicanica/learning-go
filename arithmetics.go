package main

import "fmt"

func multiplication() int {
	var result int = 3 * 10
	result += 22
	return result
}

func subtraction() int {
	var result int = 15 - 5
	result += 10
	return result 
}

func main(){
	//fmt.Printf("%v %v", multiplication(), subtraction())

	var mresult, sresult int = multiplication(), subtraction()
	fmt.Printf("%v %v", mresult, sresult)
}