package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"time"
)

func main() {
	trace.Start(os.Stderr)
	defer trace.Stop()

	table := make(chan int)
	go player(table, "ping")
	go player(table, "pong")

	table <- 0
	time.Sleep(time.Second)
	ball := <-table
	close(table)
	fmt.Printf("played %d turns\n", ball)
}

func player(table chan int, name string) {
	for ball := range table {
		fmt.Printf("%d\t%s\n", ball, name)
		table <- ball + 1
	}
}
