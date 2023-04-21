// 猜数字游戏
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(maxNum)
	fmt.Println("The secret number is ", secretNumber)

	for {
		fmt.Println("Please input your guess")
		var guess int
		fmt.Scan(&guess)

		if guess < 1 || guess > 100 {
			fmt.Println("wrong scanf")
			continue
		}

		if guess > secretNumber {
			fmt.Println("bigger")
		} else if guess < secretNumber {
			fmt.Println("smaller")
		} else {
			fmt.Println("you are righ,baby!")
			break
		}
	}
}
