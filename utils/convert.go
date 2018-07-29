package utils

import "strconv"

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
