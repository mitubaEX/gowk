package utils

import (
	"strconv"
	"fmt"
)

func StringToInt(str string) int{
	val, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return val
}

func IntToString(num int) string {
	val := strconv.Itoa(num)
	return val
}

func StringToFloat(str string) (float64, error){
	val, err := strconv.ParseFloat(str, 64)
	return val, err
}

func FloatToString(num float64) string {
	return fmt.Sprint(num)
}

func IsFloat(str string) bool {
	_, err := StringToFloat(str)
	if err != nil {
		return false
	}
	return true
}
