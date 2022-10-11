package utils

import (
	"fmt"
	"strings"
)

var traceLevel int = 0
var tracingEnabled bool = false

const traceIdentPlaceholder string = "\t"

func identLevel() string {
	return strings.Repeat(traceIdentPlaceholder, traceLevel-1)
}

func tracePrint(str string) {
	if tracingEnabled {
		fmt.Printf("%s%s\n", identLevel(), str)
	}
}

func incIdent() { traceLevel = traceLevel + 1 }
func decIdent() { traceLevel = traceLevel - 1 }

func Trace(msg string) string {
	incIdent()
	tracePrint("BEGIN " + msg)
	return msg
}

func UnTrace(msg string) {
	tracePrint("END " + msg)
	decIdent()
}
