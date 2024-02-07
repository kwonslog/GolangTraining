package main

import "fmt"

func main() {

	var myGreeting map[string]string
	fmt.Println(myGreeting)

	// 초기화 여부를 확인하기 위해 nil 비교 처리는 필요하다.
	fmt.Println(myGreeting == nil)
}

// add these lines:
/*
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."
*/
// and you will get this:
// panic: assignment to entry in nil map
