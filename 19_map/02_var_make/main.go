package main

import "fmt"

func main() {

	var myGreeting = make(map[string]string)

	// make 를 통해 생성 및 초기화가 되었기 때문에 아래 처럼 대입이 가능함.
	myGreeting["Tim"] = "Good morning."
	myGreeting["Jenny"] = "Bonjour."

	fmt.Println(myGreeting)
}
