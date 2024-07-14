package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func debugWritef(msg string, args ...interface{}) {
	fmt.Printf(msg, args...)
}

func recoverPanicToFile() {
	if x := recover(); x != nil {
		fo, err := os.Create("crash_report.log")
		if err != nil {
			panic(err)
		}
		fo.Write([]byte(fmt.Sprintf("Panic: %v \nTrace:\n%s", x, string(debug.Stack()))))

		if err := fo.Close(); err != nil {
			panic(err)
		}

		// Panic again for a crash
		panic(x)
	}
}
