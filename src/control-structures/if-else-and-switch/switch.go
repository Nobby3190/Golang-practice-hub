package ifelseandswitch

import "fmt"

func SwitchCase() {
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

func SwitchCaseTwo() {
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

func SwitchCaseThree(score int32) {
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
