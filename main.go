package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result struct {
	strikes int
	balls   int
}

func main() {
	fmt.Printf("[숫자야구게임]\n컴퓨터가 3자리의 숫자를 임의로 정했다.\n플레이어는 3자리의 숫자를 불러서 결과를 확인한다.\n그리고 그 결과를 토대로 상대가 적은 숫자를 예상한 뒤 맞힌다.\n숫자는 포함되어 있지만 위치가 틀렸을 때는 볼(B)\n숫자와 위치가 전부 맞으면 스트라이크(S)\n숫자의 중복은 없음.\n")

	//랜덤을 만드는 로직을 위해 Seed라는 것을 만들어주면 변하는 Seed값을 기반으로 숫자들을 뽑아줌
	rand.Seed(time.Now().UnixNano())
	//게임 진행을 위해 무작위 숫자 3개 생성
	numbers := MakeNumbers()
	cnt := 0

	for {
		cnt++
		//사용자 입력 받고
		inputNumbers := InputNumbers()

		//결과 비교
		result := CompareNumbers(numbers, inputNumbers)

		
		PrintResult(result)

		// 3S조건 만족하면 게임끝
		if IsGameEnd(result) {
			break
		}
	}
	fmt.Printf("[%d 번만에 맞췄습니다.]\n[10초 후 종료됩니다.]", cnt)
	time.Sleep(10 * time.Second)
}

// 1. 0~9사이의 중복되지 않는 무작위 숫자 3개 반환
func MakeNumbers() ([3]int) {
    var rst [3]int
		
		// 1-1. 3개의 숫자를 생성하기 위해 루프 3바퀴
		for i := 0; i < 3; i++ {
			// 1-2. 숫자가 겹치는 경우 다시뽑기 위한 for문
      for {  
        n := rand.Intn(10)
        // 1-3. 겹침여부 flag
        duplicated := false      

        // 1-4. 숫자 비교 위한 for문
        for j := 0; j < i; j++ { 
          if rst[j] == n {
            // 1-5. 겹쳐서 다시 뽑아줌
            duplicated = true
            break
          }
        }
				// 1-6. 겹침여부 확인 후 반복문을 빠져나가 결과값에 뽑은 숫자들을 넣음
        if !duplicated {
          rst[i] = n
          break
        }
      }
    }
		return rst
}

// 2. 0~9사이의 중복되지 않는 무작위 숫자 3개 키보드로부터 반환
func InputNumbers() ([3]int) {
    var rst [3]int

		for {
			fmt.Println("0~9사이의 숫자 3개를 입력하세요.")
			var no int
			// 키보드로 부터 입력값을 받음, return값이 입력받은 갯수, 실패시 에러 두개
			// 입력받은 갯수는 중요치 않으므로 공란 '_'
			_, err := fmt.Scanf("%d\n", &no)
			if err != nil {
				fmt.Println("입력오류")
				continue
			}

			success := true
      idx := 0

			for no > 0 {
				n := no%10
				no = no/10
				/* ex) 숫자 루프
				1. no = 123
				    n = no%10 = 3
						no = no/10 = 12
				2. no = 12
				    n = no%10 = 2
						no = no/10 = 1
				3. no = 1
				    n = no%10 = 1
						no = no/10 = 0			
				*/
        
				// 겹침여부 확인
				duplicated := false

				for j := 0; j < idx; j++{
					if rst[j] == n {
						//겹치게 되면 다시 뽑음
						duplicated = true
						break
					}
				}

				if duplicated {
					fmt.Println("숫자가 겹치지 않아야 합니다")
					success = false
					break
				}

				// 입력값이 3개를 넘어 인덱스의 범위를 벗어날 경우
				if idx >= 3 {
          fmt.Println("3개 보다 많은 숫자를 입력하였습니다")
          success = false
          break
        }

				rst[idx] = n
				idx++
			}
			
			// 입력값이 2개여서 인덱스에 0이 임의로 들어가는 경우
			if success && idx < 3 {
        fmt.Println("3개의 숫자를 입력하세요")
        success = false
      }

			if !success {
				continue
			}
			break
		}
		//입력값이 앞뒤로 바뀌여서 다시 앞뒤로 바꿈
		rst[0], rst[2] = rst[2], rst[0]
		fmt.Println(rst)
		return rst
}

// 3. 두 개의 숫자 3개를 비교해서 결과를 반환
func CompareNumbers(numbers, inputNumbers ([3]int)) (Result) {
	strikes := 0
	balls := 0
  //생성된 숫자와 입력된 숫자 비교하는 for문
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

// 4. 출력함수
func PrintResult(result Result) {
	// 스트라이크와 볼이 몇 개인지 출력
    fmt.Printf("%dS %dB\n", result.strikes, result.balls)
}

// 5. 3S조건 확인
func IsGameEnd(result Result) (bool) {
	return result.strikes == 3
}
