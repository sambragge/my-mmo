package util

import "log"

func HandleErr(err error, filename, line string){
	if err != nil {
		log.Fatalf("-- %v.go line %v -- %v", filename, line, err.Error())
	}
}
