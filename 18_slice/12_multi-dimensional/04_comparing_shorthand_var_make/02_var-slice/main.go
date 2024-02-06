package main

import (
	"fmt"
)

func main() {
	var student []string
	var students [][]string
	student[0] = "Todd"
	// student = append(student, "Todd")

	// 출력결과는 fmt.Println 에서 [] 이렇게 처리한다.
	// 하지만 실제 값은 nil 이다.
	fmt.Println(student)
	fmt.Println(students)
}
