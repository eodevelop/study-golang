package main

import "fmt"

func main() {
	var i, j int = 1, 2
	k := 3 // := 를 사용하여 var 선언 처럼 사용 가능하다.
	c, python, java := true, false, "no!"

	fmt.Println(i, j, k, c, python, java)
}
