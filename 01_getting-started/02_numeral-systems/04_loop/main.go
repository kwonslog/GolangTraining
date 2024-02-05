package main

import "fmt"

func main() {
	// 1000000 부터 1000100 까지 10진수, 2진수, 16진수(소문자)를 출력한다.
	for i := 1000000; i < 1000100; i++ {
		fmt.Printf("%d \t %b \t %x \n", i, i, i)
	}
}
