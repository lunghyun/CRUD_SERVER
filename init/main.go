package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", helloWorld) // HandleFunc가 뭐지?
	// 서버가 안꺼지게 하는 기능이 없음 -> 시작하자마자 꺼짐

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func helloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hello World")
}
