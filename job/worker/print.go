package worker

import "log"

type PrintWorker struct {
}

func NewPrintWorker() (*PrintWorker, error) {
	return &PrintWorker{}, nil
}

func (worker PrintWorker) Handler() {
	log.Println("5 minutes has passed")
}
