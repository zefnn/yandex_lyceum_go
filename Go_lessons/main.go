package main

import (
	"fmt"
	"time"
)

func NextWorkday(start time.Time) time.Time {
	next := start.AddDate(0, 0, 1)

	switch next.Weekday() {
	case time.Saturday:
		return next.AddDate(0, 0, 2)
	case time.Sunday:
		return next.AddDate(0, 0, 1)
	default:
		return next
	}

}

func main() {
	loc, err := time.LoadLocation("UTC")
	if err != nil {
		fmt.Println("Ошибка")
		return
	}

	currentTime := time.Now().In(loc)
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	fmt.Println("Текущее время (UTC):", formattedTime)
}
