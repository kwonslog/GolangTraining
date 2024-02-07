package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	myGreeting["Harleen"] = "Howdy"

	// 맵에 추가된 데이터의 개수 출력
	fmt.Println(len(myGreeting))
}
