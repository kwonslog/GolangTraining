package main

import "fmt"

func main() {

	myGreeting := map[string]string{
		"zero":  "Good morning!",
		"one":   "Bonjour!",
		"two":   "Buenos dias!",
		"three": "Bongiorno!",
	}

	fmt.Println(myGreeting)

	// 키가 "two" 인 항목을 제거한다.
	delete(myGreeting, "two")

	fmt.Println(myGreeting)
}
