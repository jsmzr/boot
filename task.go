package boot

import (
	"fmt"

	"github.com/robfig/cron/v3"
)

type Task struct {
	Name    string
	Cron    string
	Handler func()
}

var tasks = make([]Task, 0)

func RegisterTask(name, cron string, handler func()) {
	tasks = append(tasks, Task{Name: name, Cron: cron, Handler: handler})
}

func runTask() error {
	if len(tasks) == 0 {
		Log("No task to run")
		return nil
	}
	cron := cron.New(cron.WithSeconds())
	for _, task := range tasks {
		Log(fmt.Sprintf("Add task [%s]: [%s]", task.Cron, task.Name))
		_, err := cron.AddFunc(task.Cron, task.Handler)
		if err != nil {
			return err
		}
	}
	cron.Start()
	Log("Task started")
	return nil
}
