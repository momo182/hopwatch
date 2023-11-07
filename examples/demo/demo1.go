package main

import (
	"os"
	"time"

	"github.com/momo182/hopwatch"
)

func main() {
	haveABreak()
}

func haveABreak() {
	hopwatch.Break()
	printNow()
}

func printNow() {
	hopwatch.Printf("time is: %v", time.Now())
	dumpArgs()
}

func dumpArgs() {
	hopwatch.Dump(os.Args).Break()
	waitHere()
}
