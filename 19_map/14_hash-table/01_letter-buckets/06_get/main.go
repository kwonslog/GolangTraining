package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {

		// 오류 메세지를 출력하고 프로그램은 종료 된다.
		// 그리고 os.Exit(1) 을 호출하기 때문에 즉시 종료 되며
		// **defer 로 등록한 함수들이 실행되지 않는다.**
		log.Fatal(err)
	}
	bs, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", bs)
}
