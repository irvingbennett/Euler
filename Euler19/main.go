package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Counting Sundays")
	panama, err := time.LoadLocation("America/Panama")
	if err != nil {
		fmt.Println(err)
	}
	year := 1901
	month := time.January
	count := 0
	for {
		d := time.Date(year, month, 1, 1, 1, 1, 1, panama)
		if d.Weekday() == time.Sunday {
			count++
		}

		fmt.Println(month)
		if month != time.December {
			month++
		} else {
			month = time.January
			year++
			if year > 2000 {
				break
			}
		}
	}
	fmt.Println(count)
}
