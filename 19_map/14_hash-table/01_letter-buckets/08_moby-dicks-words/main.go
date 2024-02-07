package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	// NewScanner takes a reader and res.Body implements the reader interface (so it is a reader)
	scanner := bufio.NewScanner(res.Body)

	// 사용했으니 잘 닫아줘야 메모리 이슈 발생하지 않는다.
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Loop over the words
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
