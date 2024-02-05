package main

import "fmt"

func main() {
	//	fmt.Printf("%d - %b - %x \n", 42, 42, 42)
	//	fmt.Printf("%d - %b - %#x \n", 42, 42, 42)
	//	fmt.Printf("%d - %b - %#X \n", 42, 42, 42)
	
	// 값 42를 10진수, 2진수, 16진수(대문자)로 출력한다.
	fmt.Printf("%d \t %b \t %#X \n", 42, 42, 42)
}
