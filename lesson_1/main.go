package main

import (
	"fmt"
)

func main() {
	fmt.Println("Pomodoro Timer v0.1")

	for i := 1; i < 31; i++ {
		fmt.Printf("number %d is %s\n", i, fizzBuzz(i))
	}

	fmt.Println(divide(1, 0))
	fmt.Println(mibutesToPomodoros(60))
}

func fizzBuzz(n int) string {
	if n%15 == 0 {
		return "FizzBuzz"
	} else if n%3 == 0 {
		return "Fizz"
	} else if n%5 == 0 {
		return "Buzz"
	}

	return ""
}

func divide(a, b float64) (float64, error)  {
	if (b == 0) {
		return 0, fmt.Errorf("0 division error")
	}

	return a / b, nil
}

func mibutesToPomodoros(minutes int) (sessions int, breackTime int) {
	const workDuration = 25
	const breackDuration = 5

	totalSessions := workDuration + breackDuration
	sessions = minutes / totalSessions
	breackTime = minutes % totalSessions

	return
}