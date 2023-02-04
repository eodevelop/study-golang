package main

import (
	"fmt"
	"math/cmplx"
)

/** GO의 기본 타입
bool

string

int  int8  int16  int32  int64
uint uint8 uint16 uint32 uint64 uintptr

byte // uint8의 별칭

rune // int32의 별칭
	// 유니코드에서 code point를 의미합니다.

float32 float64

complex64 complex128\
*/

/** 초기 값이 없이 생성되는 경우 해당 형태의 Zero Value가 할당

숫자 type에는 0
boolean type에는 false
string에는 "" (빈 문자열)

*/

var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}
