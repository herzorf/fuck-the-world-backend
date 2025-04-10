package utils

import "log"

func HandleError(text string, err error) {
	if err != nil {
		log.Printf("%s: %v\n", text, err)
	}
}
