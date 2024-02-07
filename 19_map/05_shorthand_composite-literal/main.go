package main

import "fmt"

func main() {

	// 리터럴 방식으로 선언과 초기화를 동시에 한다.
	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	fmt.Println(myGreeting["Jenny"])
}
