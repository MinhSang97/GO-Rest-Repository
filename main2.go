package main

import (
	"fmt"
	"time"
)

func main() {
	StartTime := "01-01-2024 12:00:00"
	fmt.Println("Start Date before:", StartTime)

	EndTime := "01-01-2024 12:00:00"
	fmt.Println("End Datebefore:", EndTime)
	Period := "1D"

	startDate, err := time.Parse("02-01-2006 15:04:05", StartTime)
	if err != nil {
		fmt.Println("Error parsing start date:", err)
		return
	}

	endDate, err := time.Parse("02-01-2006 15:04:05", EndTime)
	if err != nil {
		fmt.Println("Error parsing end date:", err)
		return
	}

	switch Period {
	case "30M":
		endDate = startDate.Add(30 * time.Minute)
	case "1H":
		endDate = startDate.Add(1 * time.Hour)
	case "1D":
		endDate = startDate.Add(24 * time.Hour)
	default:
		fmt.Println("Invalid period")
		return
	}

	StartTime = fmt.Sprintf("%d", startDate.UnixNano()/int64(time.Millisecond))
	EndTime = fmt.Sprintf("%d", endDate.UnixNano()/int64(time.Millisecond))

	fmt.Println("Start Time:", StartTime)
	fmt.Println("End Time:", EndTime)
}
