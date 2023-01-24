package main

import "fmt"

// 같은 타입이 연속해서 나온다면 x int, y int => x, y int 로 변환 가능
func add(x int, y int) int { // 변수명 먼저, 변수 타입 뒤에
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
