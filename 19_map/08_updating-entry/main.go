package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"Tim":   "Good morning!",
		"Jenny": "Bonjour!",
	}

	// 값을 추가 하고
	myGreeting["Harleen"] = "Howdy"
	fmt.Println(myGreeting)

	// 다시 수정하고
	myGreeting["Harleen"] = "Gidday"
	fmt.Println(myGreeting)
}
