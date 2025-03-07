package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 4 {
		fmt.Printf("Usage:\t%v <initial value> <result value> <days invested>\n", os.Args[0])
		return
	}
	initVal, err := strconv.ParseFloat(os.Args[1], 64)
	if err != nil {
		log.Fatalln("can't convert initial value to float")
	}
	resVal, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		log.Fatalln("can't convert initial value to float")
	}
	daysInv, err := strconv.ParseFloat(os.Args[3], 64)
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
		apr = (math.Pow((resVal / initVal), 365 / daysInv) - 1) * 100
	}

	fmt.Printf("APR = %.2f%%\n", apr)

}
