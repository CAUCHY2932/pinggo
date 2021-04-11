package main

import "fmt"

func main() {

	fmt.Println("hello")
	p := PingTask{}
	p.currentSeq = 8

	fmt.Println("p.currentSeq = ", p.currentSeq)
}
