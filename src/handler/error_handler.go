package handler

import "fmt"

func ErrorCheck(err error, message string) {
	if err != nil {
		fmt.Errorf(message+"%v", err)
		panic(err)
	}
}
