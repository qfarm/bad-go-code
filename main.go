package main

import "github.com/qfarm/bad-go-code/logger"

func main() {
	l, _ := logger.NewLogger("/tmp/test")
	l.Log("Test log")
}

func dead() string {
	return deadFunc()
}

func deadFunc() string {
	return "I am dead"
}
