package Helpers

import "log"

func ValidateErr(err error) {
	if err != nil {
		log.Fatalf("Error happened. Err: %s", err)
		panic("Fatal Error")
	}
}
