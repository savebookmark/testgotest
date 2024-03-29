package baseball

import (
	"fmt"
	"math/rand"
	// "time"
)

// / Result 구조체
type Result struct {
	strikes int
	balls   int
}

func Bsmain() {
	// rand.Seed(time.Now().UnixNano())
	// 무작위 숫자 3개를 만든다.
	numbers := MakeNumbers()

	cnt := 0
	for {
		cnt++
		// 사용자의 입력을 받는다.
		inputNumbers := InputNumbers()

		// 결과를 비교한다.
		result := CompareNumbers(numbers, inputNumbers)

		PrintResult(result)

		// 3S 인가?
		if IsGameEnd(result) {
			// 게임끝
			break
		}
	}

	// 게임끝. 몇번만에 맞췄는지 출력.
	fmt.Printf("%d 번만에 맞췄습니다.\n", cnt)
}

// MakeNumbers 0~9 사이의 겹치지 않는 무작위 숫자 3개를 반환한다.
func MakeNumbers() [3]int {
	var rst [3]int

	for i := 0; i < 3; i++ {
		for {
			n := rand.Intn(10)
			duplicated := false
			for j := 0; j < i; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if !duplicated {
				rst[i] = n
				break
			}
		}
	}

	//fmt.Println(rst)
	return rst
}

// InputNumbers 키보드로부터 0~9 사이의 겹치지 않는 숫자 3개를 입력받아 반환한다.
func InputNumbers() [3]int {
	var rst [3]int

	for {
		fmt.Println("겹치지 않는 0~9 사이의 숫자 3개를 입력하세요.")
		var no int
		_, err := fmt.Scanf("%d\n", &no)
		if err != nil {
			fmt.Println("잘못 입력하셨습니다.")
			continue
		}

		success := true
		idx := 0
		for no > 0 {
			n := no % 10
			no = no / 10

			duplicated := false
			for j := 0; j < idx; j++ {
				if rst[j] == n {
					// 겹친다. 다시 뽑는다.
					duplicated = true
					break
				}
			}
			if duplicated {
				fmt.Println("숫자가 겹치지 않아야 합니다.")
				success = false
				break
			}

			if idx >= 3 {
				fmt.Println("3개보다 많은 숫자를 입력하셨습니다.")
				success = false
				break
			}

			rst[idx] = n
			idx++
		}
		if success && idx < 3 {
			fmt.Println("3개의 숫자를 입력하세요.")
			success = false
		}

		if !success {
			continue
		}
		break
	}
	rst[0], rst[2] = rst[2], rst[0]
	//fmt.Println(rst)
	return rst
}

// CompareNumbers 두개의 숫자 3개를 비교해서 결과를 반환한다.
func CompareNumbers(numbers, inputNumbers [3]int) Result {
	strikes := 0
	balls := 0

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if numbers[i] == inputNumbers[j] {
				if i == j {
					strikes++
				} else {
					balls++
				}
				break
			}
		}
	}

	return Result{strikes, balls}
}

// PrintResult aaa
func PrintResult(result Result) {
	fmt.Printf("%dS%dB\n", result.strikes, result.balls)
}

// IsGameEnd  비교 결과가 3 스트라이크 인지 확인
func IsGameEnd(result Result) bool {
	return result.strikes == 3
}
