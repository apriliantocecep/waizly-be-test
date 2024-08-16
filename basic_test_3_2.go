package main

import (
	"fmt"
	"time"
)

func main() {
	// validate time format
	//inputTimeString := "12:05:45AM"
	inputTimeString := "07:05:45PM"

	layoutFormat := "15:04:05PM"
	t, err := time.Parse(layoutFormat, inputTimeString)

	if err != nil {
		panic(err)
	}

	// get time only
	timeOnly := t.Format(layoutFormat)

	fmt.Println(timeOnly)
}
