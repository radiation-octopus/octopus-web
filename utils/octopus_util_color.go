package utils

import (
	"fmt"
)

func GetColorStr(color string, isHighLighted bool, str string) string {
	var colorInt int
	switch color {
	case "black":
		colorInt = 30
	case "red":
		colorInt = 31
	case "green":
		colorInt = 32
	case "yellow":
		colorInt = 33
	case "blue":
		colorInt = 34
	case "purple":
		colorInt = 35
	case "cyan":
		colorInt = 36
	case "white":
		colorInt = 37
	default:
		colorInt = 30
	}
	if isHighLighted {
		return fmt.Sprintf("\x1b[1;%dm%s\x1b[0m", colorInt, str)
	} else {
		return fmt.Sprintf("\x1b[0;%dm%s\x1b[0m", colorInt, str)
	}

}
