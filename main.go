package main

import (
	"fmt"
	"time"
)

type Transaction struct {
	Value     int
	Timestamp int
}

func graphicFunc(graphic []*Transaction, x string) {
	var filteredGraphic []Transaction

	for _, v := range graphic {
		timestamp := int64(v.Timestamp)
		t := time.Unix(timestamp, 0)

		if x == "day" {
			dayOfMonth := t.Day()
			if dayOfMonth == 18 {
				filteredGraphic = append(filteredGraphic, Transaction{Value: v.Value, Timestamp: v.Timestamp})
			}
		} else if x == "month" {
			month := t.Month()
			monthString := month.String()
			if monthString == "March" {
				filteredGraphic = append(filteredGraphic, Transaction{Value: v.Value, Timestamp: v.Timestamp})
			}
		} else if x == "hour" {
			hour := t.Hour()
			if hour == 10 {
				filteredGraphic = append(filteredGraphic, Transaction{Value: v.Value, Timestamp: v.Timestamp})
			}
		}
	}
	fmt.Println(filteredGraphic)
}

func main() {
	graphic := []*Transaction{
		{
			4456,
			1616026248,
		},
		{
			4231,
			1616022648,
		},
		{
			5212,
			1616019048,
		},
		{
			4321,
			1615889448,
		},
		{
			4567,
			1615871448,
		},
	}
	//interval1 := "day"
	interval1 := "month"
	//interval1 := "hour"
	graphicFunc(graphic, interval1)
}
