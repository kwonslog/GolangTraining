package main

import "fmt"

func main() {
	// 60부터 122까지 10진수, 2진수, 16진수(소문자), 문자를 출력한다.
	for i := 60; i < 122; i++ {
		fmt.Printf("%d \t %b \t %x \t %q \n", i, i, i, i)
	}
}
