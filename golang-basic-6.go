package main

import "fmt"

// 반환은 갯수에 상관없이 가능하다.
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
