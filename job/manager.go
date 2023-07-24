package job

import (
	"log"
	"time"

	"github.com/m25-lab/lightning-network-node/client"
	"github.com/m25-lab/lightning-network-node/database/models"
	"github.com/m25-lab/lightning-network-node/job/worker"
	"github.com/m25-lab/lightning-network-node/node"
)

type Job struct {
	worker   worker.JobWorker
	duration int64
}

type Manager struct {
	JobMap map[string]Job
}

func New(node *node.LightningNode, client *client.Client) (manger *Manager, err error) {
	manger = new(Manager)
	manger.JobMap = make(map[string]Job)

	// printWK, err := worker.NewPrintWorker()
	// if err != nil {
	// 	log.Println("NewPrintWorker error: ", err.Error())
	// 	return
	// }
	// manger.JobMap["print"] = Job{worker: printWK, duration: 300}

	checkFindRouteWK, err := worker.NewCheckFindRoute(node.Repository, client)
	if err != nil {
		log.Println("NewCheckFindRoute error: ", err.Error())
		return
	}
	manger.JobMap[models.CheckFindRouteJobName] = Job{worker: checkFindRouteWK, duration: 10}

	return
}

func (manger *Manager) Run() {
	for _, jobWorker := range manger.JobMap {
		go func(job Job) {
			for true {
				time.Sleep(time.Second * time.Duration(job.duration))
				job.worker.Handler()
			}
		}(jobWorker)
	}
}
