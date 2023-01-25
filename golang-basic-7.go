package main

import "fmt"

// 반환값에 이름을 지정 후 함수의 맨 위에서 정의된 변수처럼 다룰 수 있다.
// 인자가 없는 return 의 경우 주어진 반환값을 반환하며 이를 naked return 이라고 한다.
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return 
}

func main() {
	fmt.Println(split(17))
}
