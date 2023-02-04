package main

import (
	"fmt"
	"math"
)

func main() {
	// 형 변환 시 아래와 같이 T(v)의 형태로 명시적인 형 변환 필요
	var x, y int = 3, 4
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)
	fmt.Println(x, y, z)
}
