package ifelseandswitch

import "fmt"

func IfElse() int {
	var x int = 1
	if x > 0 {
		return 1
	}
	return 0
}

func IfElseTwo() {
	input := 10
	if remain := input % 2; remain == 1 {
		fmt.Printf("%d 為奇數\n", input)
	} else {
		fmt.Printf("%d 為偶數\n", input)
	}
}

func IfElseThree() {
	var level rune
	if score := 88; score > 90 {
		level = 'A'
	} else if score >= 80 && score < 90 {
		level = 'B'
	} else if score >= 70 && score < 80 {
		level = 'C'
	} else if score >= 60 && score < 70 {
		level = 'D'
	} else {
		level = 'E'
	}
	fmt.Printf("得分等級：%c\n", level)
}
