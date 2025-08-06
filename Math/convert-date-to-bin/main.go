package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertDateToBinary(date string) string {
	// Convert date string to int
	dateSplit := strings.Split(date, "-")
	// var binArr []string
	for k, val := range dateSplit {
		num, _ := strconv.Atoi(val)
		dateBin := strconv.FormatInt(int64(num), 2)
		dateSplit[k] = dateBin
	}
	// fmt.Println(dateSplit)
	return strings.Join(dateSplit, "-")
}

func main() {
	fmt.Print("Result: ", convertDateToBinary("2022-05-01"))
}
