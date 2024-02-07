package main

import "fmt"

func main() {

	myGreeting := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	fmt.Println(myGreeting)

	// delete(myGreeting, 2)

	// 맵에서 꺼낸 값은 val, 값이 있는지 여부는 exists 이것을 "comma ok" 패턴이라고 한다.
	if val, exists := myGreeting[2]; exists {
		fmt.Println("That value exists.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val) // val 의 값은 "" 빈문자열이다.
		fmt.Println("exists: ", exists)
	}

	fmt.Println(myGreeting)
}
