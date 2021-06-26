package sampleTut

import (
	"fmt"
	"log"
	"reflect"
	"strconv"
	"time"
)

func DateConverter(timeVal, dateVal int64) *time.Time {

	str := strconv.Itoa(int(dateVal)) + " " + strconv.Itoa(int(timeVal))
	fmt.Println(str)

	// layoutStr := "2006-01-02T03:04:05:0006Z"
	// myDateString := "2012-07-21T11:12:32.3434Z"
	// fmt.Printf("This is the date string %T \t %v \n", myDateString, myDateString)

	min := time.Minute
	fmt.Printf("Minute is %v \n", min)

	layoutStr := "20060102 150304"
	duration, err := time.Parse(layoutStr, str)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Duration is %v \t type: %v \n", duration, reflect.TypeOf(duration))

	return &duration
}

// This method creates a formatted output of current time
func GetCurrentTimeString() *string {
	currentTime := time.Now().Format("2006-01-02 15:04:05.0006")

	// get current time in format yyy-mm-dd
	// 01- corresponds to day of month, 02- month, 03- hour, 04- min, 05- sec, 06 - MilliSeconds, 15 fr hour in 24 hr format
	fmt.Printf("VV: %T  \t %v \n", currentTime, currentTime)

	return &currentTime

}
