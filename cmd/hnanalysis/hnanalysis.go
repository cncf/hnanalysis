package main

import (
	"fmt"
	lib "hnanalysis"
	"time"
)

func main() {
	dtStart := time.Now()
	_ = lib.Dummy()
	dtEnd := time.Now()
	fmt.Printf("Time: %v\n", dtEnd.Sub(dtStart))
}
