package main

import (
	"fmt"
	"strconv"
	"strings"
)

func timeConversion(s string) string {
	parts := strings.Split(s, ":")

	hour, _ := strconv.Atoi(parts[0])
	minute, _ := strconv.Atoi(parts[1])
	second := parts[2][:2]
	meridian := parts[2][2:]

	if meridian == "PM" {
		if hour != 12 {
			hour += 12
		}
	} else {
		if hour == 12 {
			hour = 0
		}
	}

	return fmt.Sprintf("%02d:%02d:%s", hour, minute, second)
}

func main() {
	fmt.Println(timeConversion("12:01:00PM"))
	fmt.Println(timeConversion("12:01:00AM"))
	fmt.Println(timeConversion("07:05:45PM"))
	fmt.Println(timeConversion("01:05:45PM"))
}
