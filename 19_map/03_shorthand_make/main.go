package main

import "fmt"

func main() {

	// 타입을 명시하지 않고 사용가능. 타입은 컴파일러가 유추한다.
	myGreeting := make(map[string]string)
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
