package main

import "fmt"

func sum() int {
	var sum int = 3 * 10
	sum += 22
	return sum
}

func main(){
	var result int = sum()
	fmt.Println(result)
}