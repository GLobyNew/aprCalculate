package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"time"
)

const (
	floatBase  = 64
	hoursInDay = 24
)

func calculateBasedOnDayNumber() error {
	initVal, err := strconv.ParseFloat(os.Args[1], floatBase)
	if err != nil {
		log.Fatalln("can't convert initial value to float")
	}
	resVal, err := strconv.ParseFloat(os.Args[2], floatBase)
	if err != nil {
		log.Fatalln("can't convert res value to float")
	}
	daysInv, err := strconv.ParseFloat(os.Args[3], floatBase)
	if err != nil {
		log.Fatalln("can't convert days invested to float")
	}

	if daysInv == 0 {
		log.Fatalln("invested days can't be 0")
	}

	var apr float64

	if daysInv <= 365 {
		apr = ((resVal - initVal) / initVal) * (365 / daysInv) * 100
	} else {
		apr = (math.Pow((resVal/initVal), 365/daysInv) - 1) * 100
	}

	fmt.Printf("APR = %.2f%%\n", apr)

	return nil
}

func calculateBasedOnDates() error {
	initVal, err := strconv.ParseFloat(os.Args[1], floatBase)
	if err != nil {
		log.Fatalln("can't convert initial value to float")
	}
	resVal, err := strconv.ParseFloat(os.Args[2], floatBase)
	if err != nil {
		log.Fatalln("can't convert res value to float")
	}
	startDate, err := time.Parse(time.DateTime, os.Args[3])
	if err != nil {
		log.Fatalln("can't parse start date")
	}
	endDate, err := time.Parse(time.DateTime, os.Args[4])
	if err != nil {
		log.Fatalln("can't parse end date")
	}

	dur := endDate.Sub(startDate)

	if dur.Hours() <= 0 {
		log.Fatalln("invested hours can't be zero")
	}

	daysInv := dur.Hours() / hoursInDay

	if daysInv <= 0 {
		log.Fatalln("invested days can't be 0 or less")
	}

	var apr float64

	if daysInv <= 365 {
		apr = ((resVal - initVal) / initVal) * (365 / daysInv) * 100
	} else {
		apr = (math.Pow((resVal/initVal), 365/daysInv) - 1) * 100
	}

	fmt.Printf("APR = %.2f%%\n", apr)

	return nil
}

func main() {
	args := os.Args
	if len(args) != 4 && len(args) != 5 {
		fmt.Printf("Usage:\t%v <initial value> <result value> <days invested>\n", os.Args[0])
		fmt.Printf("Usage:\t%v <initial value> <result value> <start date> <end date>\n", os.Args[0])
		return
	}

	if len(args) == 4 {
		calculateBasedOnDayNumber()
	} else if len(args) == 5 {
		calculateBasedOnDates()
	} else {
		log.Fatalln("invalid number of arguments")
	}
}
