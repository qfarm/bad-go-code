package calculator

import "github.com/qfarm/bad-go-code/logger"

func Add(a, b int) int {
	return a + b



}

func Multiple(a, b int) 	int {
	l, _ := logger.NewLogger("/tmp/test1")

	l.UnusedExt = "test7"

	return a * b
}

func Divide(a, b int) int {
	return a / b
}