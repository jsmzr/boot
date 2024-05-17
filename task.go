package boot

import (
	"fmt"

	"github.com/robfig/cron/v3"
	"github.com/spf13/viper"
)

type Task struct {
	Name    string
	Handler func()
}

var tasks = make([]Task, 0)

func RegisterTask(name string, handler func()) {
	tasks = append(tasks, Task{Name: name, Handler: handler})
}

func runTask() error {
	if len(tasks) == 0 {
		Log("No task to run")
		return nil
	}
	cron := cron.New(cron.WithSeconds())
	for _, task := range tasks {
		configCron := viper.GetString("task.cron." + task.Name)
		Log(fmt.Sprintf("Add task [%s]: [%s]", configCron, task.Name))
		_, err := cron.AddFunc(configCron, task.Handler)
		if err != nil {
			return err
		}
	}
	cron.Start()
	Log("Task started")
	return nil
}
