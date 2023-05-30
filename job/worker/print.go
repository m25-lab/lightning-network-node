package worker

import (
	"fmt"
)

type PrintWorker struct {
}

func NewPrintWorker() (*PrintWorker, error) {
	return &PrintWorker{}, nil
}

func (worker PrintWorker) Handler() {
	fmt.Println("HEHE")
}
