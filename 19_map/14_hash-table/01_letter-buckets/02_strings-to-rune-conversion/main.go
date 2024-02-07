package main

import "fmt"

func main() {
	// "A" 문자열의 0 번째 값을 가져온다.
	// rune 을 사용해서 문자 타입으로 캐스팅한다.
	letter := rune("A"[0])
	fmt.Println(letter)

	// rune 로 타입 캐스팅 전
	fmt.Printf("%T \n", "A"[0]) //uint8
	// rune 로 타입 캐스팅 후
	fmt.Printf("%T \n", letter) //int32

}
