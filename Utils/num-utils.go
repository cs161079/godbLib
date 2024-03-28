package utils

import (
	"strconv"
	"strings"
)

func StrToInt8(input interface{}) int8 {
	return int8(stringToNumberInternal(input.(string), 8))
}

func StrToInt64(input interface{}) int64 {
	return stringToNumberInternal(input.(string), 64)
}

func StrToInt32(input interface{}) int32 {
	return int32(stringToNumberInternal(input.(string), 32))
}

func StrToInt16(input interface{}) int16 {
	return int16(stringToNumberInternal(input.(string), 16))
}

func stringToNumberInternal(input string, bitSize int) int64 {
	sourceNumVal, error := strconv.ParseInt(strings.Trim(input, " "), 10, bitSize)
	if error != nil {
		panic("Δεν ήταν δυνατή η μετατροπή της συμβολοσειράς σε αριθμό για το πεδίο " + input)
	}
	return sourceNumVal
}

func stringToFloatInternal(input string, bitSize int) float64 {
	sourceNumVal, error := strconv.ParseFloat(strings.Trim(input, " "), bitSize)
	if error != nil {
		panic("Δεν ήταν δυνατή η μετατροπή της συμβολοσειράς σε αριθμό για το πεδίο " + input)
	}
	return sourceNumVal
}

func StrToFloat(input interface{}) float64 {
	return stringToFloatInternal(input.(string), 64)
}
func StrToFloat32(input interface{}) float32 {
	return float32(stringToFloatInternal(input.(string), 32))
}
