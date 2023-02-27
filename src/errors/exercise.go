package error_exercise

import (
	"fmt"
	"strconv"
	"strings"
)

type Time struct {
	hour, minute, second int
}

type TimeParseError struct {
	msg   string
	input string
}

func (t *TimeParseError) Error() string {
	return fmt.Sprintf("%v: %v", t.msg, t.input)
}

func ParseTime(input string) (Time, error) {
	components := strings.Split(input, ":")
	if len(components) != 3 {
		return Time{}, &TimeParseError{"Invalid number of time components", input}
	}

	hour, err := strconv.Atoi(components[0])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing hour: %v", err), input}
	}

	minute, err := strconv.Atoi(components[1])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing minute: %v", err), input}
	}

	second, err := strconv.Atoi(components[2])
	if err != nil {
		return Time{}, &TimeParseError{fmt.Sprintf("Error parsing second: %v", err), input}
	}

	if hour > 23 || hour < 0 {
		return Time{}, &TimeParseError{"Hour out of range <0,24)", fmt.Sprintf("%v", hour)}
	}

	if minute > 59 || minute < 0 {
		return Time{}, &TimeParseError{"Minute out of range <0,60)", fmt.Sprintf("%v", minute)}
	}

	if second > 59 || second < 0 {
		return Time{}, &TimeParseError{"Second out of range <0,60)", fmt.Sprintf("%v", second)}
	}

	return Time{
		hour:   hour,
		minute: minute,
		second: second,
	}, nil
}
