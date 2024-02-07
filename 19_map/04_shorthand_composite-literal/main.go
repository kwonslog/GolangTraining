package main

import "fmt"

func main() {

	// 리터럴 방식은 {} 안에 초기화 값을 넣을 수 있다.
	// make 방식은 초기화만 가능하다.
	myGreeting := map[string]string{}

	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
