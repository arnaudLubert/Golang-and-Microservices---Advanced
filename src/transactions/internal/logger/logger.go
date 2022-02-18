package logger

import (
	"log"
	"runtime"
	"strings"
)

func Log(args ...interface{}) {
	av := args
	pc, file, line, success := runtime.Caller(1)

	if success {
		index := strings.LastIndexByte(file, '/')

		if index != -1 {
			file = file[index+1:]
		}
		log.Printf("%s() %s:%d %v", runtime.FuncForPC(pc).Name()[5:], file, line, av)
	} else {
		log.Print(av)
	}
}
