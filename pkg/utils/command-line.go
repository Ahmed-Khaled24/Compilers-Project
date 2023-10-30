package utils

import (
	"os"
)

// Function to get values passed as argument while executing the app.
// Param:
//
//	arg_value: argument that you want its value ex. --path
//	all: return list of all argv
//
// Return:
//
//	return the value of the desired arg , and if the all is setted will return all argv
func GetValueOfArg(arg_value string, all bool) (string, []string) {
	argv := os.Args[1:]
	if all == true {
		return "", argv
	}
	for index, value := range argv {
		if value == arg_value {
			if index+1 < len(argv) {
				return argv[index+1], nil
			} else {
				panic("OutOfIndex")
			}
		}
	}
	return "", nil
}
