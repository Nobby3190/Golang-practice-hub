package main

import "fmt"

/*
func if_else() int {
	var x int = 1
	if x > 0 {
		return 1
	}
	return 0
}


func if_else_two() {
	input := 10
	if remain := input % 2; remain == 1 {
		fmt.Printf("%d 為奇數\n", input)
	} else {
		fmt.Printf("%d 為偶數\n", input)
	}
}


func if_else_three() {
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


func switch_case() {
	var level rune
	score := 71

	switch score / 10 {
	case 10, 9:
		level = 'A'
	case 8:
		level = 'B'
	case 7:
		level = 'C'
	case 6:
		level = 'D'
	default:
		level = 'E'
	}
	fmt.Printf("得分等級：%c\n", level)
}


func switch_case_two() {
	var level rune

	switch score := 100; score / 10 {
	case 10:
		fmt.Println("滿分")
		fallthrough
	case 9:
		level = 'A'
	case 8:
		level = 'B'
	case 7:
		level = 'C'
	case 6:
		level = 'D'
	default:
		level = 'E'
	}
	fmt.Printf("得分等級：%c\n", level)
}
*/

func switch_case_three(score int32) {
	var level rune

	switch {
	case score == 100:
		fmt.Println("滿分")
		fallthrough
	case score >= 90:
		level = 'A'
	case score >= 80 && score < 90:
		level = 'B'
	case score >= 70 && score < 80:
		level = 'C'
	case score >= 60 && score < 70:
		level = 'D'
	default:
		fmt.Println("不及格")
		level = 'E'
	}
	fmt.Printf("得分等級：%c\n", level)
}

func main() {
	switch_case_three(100)
}
