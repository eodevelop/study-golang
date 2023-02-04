package main

import "fmt"

var i, j int = 1, 2

func main() {
	var c, python, java = true, false, "no!" // 초기값 세팅 시 타입 생략 가능
	fmt.Println(i, j, c, python, java)
}
