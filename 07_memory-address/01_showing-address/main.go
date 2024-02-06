package main

import "fmt"

func main() {

	a := 43

	fmt.Println("a - ", a)
	fmt.Println("a's memory address - ", &a)
	// &a 의 주소값(16진수)을 10진수로 변환해서 출력 된다.
	fmt.Printf("%d \n", &a)
}
