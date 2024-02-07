package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	// 맵에 값을 추가한다.
	myGreeting["Harleen"] = "Howdy"

	fmt.Println(myGreeting)
}
