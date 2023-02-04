package main

import "fmt"

// pakage 단에서도 선언 가능
var c, python, java bool

func main() {
	var i int //타입이 가장 마지막에 온다.
	fmt.Println(i, c, python, java)
}
