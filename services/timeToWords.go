package services

import (
	"fmt"
	"strings"
)

//This function mainly convert numeric time to human friendly text
func TimeToWords(hour int, min int) string {

	if hour > 23 || hour < 0 || min > 59 || min < 0 {
		return "Invalid input"
	}

	nums := [...]string{"twelve", "one", "two", "three", "four",
		"five", "six", "seven", "eight", "nine",
		"ten", "eleven", "twelve", "thirteen",
		"fourteen", "fifteen", "sixteen", "seventeen",
		"eighteen", "nineteen", "twenty", "twenty one",
		"twenty two", "twenty three", "twenty four",
		"twenty five", "twenty six", "twenty seven",
		"twenty eight", "twenty nine",
	}
	result := ""
	switch {
	case min == 0:
		result = fmt.Sprintf("%s o'clock", FirstUpper(nums[hour%12]))
	case min == 1:
		result = fmt.Sprintf("One past %s", nums[hour%12])
	case min == 59:
		result = fmt.Sprintf("One to %s", nums[(hour%12)+1])
	case min == 15:
		result = fmt.Sprintf("Quarter past %s", nums[hour%12])
	case min == 30:
		result = fmt.Sprintf("Half past %s", nums[hour%12])
	case min == 45:
		result = fmt.Sprintf("Quarter to %s", nums[(hour%12)+1])
	case min < 30:
		result = fmt.Sprintf("%s past %s", FirstUpper(nums[min]), nums[hour%12])
	case min > 30:
		result = fmt.Sprintf("%s to %s", FirstUpper(nums[60-min]), nums[(hour%12)+1])
	}
	return result
}

//This function Upper first letter
func FirstUpper(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(s[:1]) + s[1:]
}
