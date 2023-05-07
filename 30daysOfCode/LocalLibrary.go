package main

import (
	"fmt"
	"time"
)

func main_l() {
	readDates()
}

func readDates() error {
	var returnedDay, returnedMonth, returnedYear int
	var dueDay, dueMonth, dueYear int

	_, err := fmt.Scanf("%d %d %d\n", &returnedDay, &returnedMonth, &returnedYear)
	if err != nil {
		return err
	}
	_, err1 := fmt.Scanf("%d %d %d\n", &dueDay, &dueMonth, &dueYear)
	if err1 != nil {
		return err1
	}

	rDate := time.Date(returnedYear, time.Month(returnedMonth), returnedDay, 0, 0, 0, 0, time.UTC)
	dDate := time.Date(dueYear, time.Month(dueMonth), dueDay, 0, 0, 0, 0, time.UTC)

	fine := calculateFine(rDate, dDate)

	fmt.Println(fine)

	return nil
}

func calculateFine(rDate time.Time, dDate time.Time) int {
	if rDate.After(dDate) {
		if dDate.Month() == rDate.Month() && dDate.Year() == rDate.Year() && dDate.Day() < rDate.Day() {
			return 15 * (rDate.Day() - dDate.Day())
		}
		if dDate.Month() < rDate.Month() && dDate.Year() == rDate.Year() {
			return 500 * (int(rDate.Month()) - int(dDate.Month()))
		}
		return 10000
	}
	return 0
}
