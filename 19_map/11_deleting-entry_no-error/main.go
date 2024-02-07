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

	// 없는 key 를 제거하려고 시도해도 오류가 발생하지 않는다.

	// 이유는 go의 설계 철학과 관련이 있다.
	// 간결성 : 삭제를 위해 사전에 존재 여부를 체크하지 않아도 된다.
	// 효율성 : 사전에 체크하지 않기 때문에 불필요한 오버헤드가 없다.
	// 안정성 : 삭제 대상이 없어도 오류가 발생하지 않는다.
	delete(myGreeting, 7)

	fmt.Println(myGreeting)
}
