package threads

import "acs-sdk-go/threads/safe"

type limit interface{}

type Task struct {
	lt chan limit
}

// NewTask init Task
// lt chan
// concurrency int  source args int
func NewTask(concurrency int) *Task {
	return &Task{lt: make(chan limit, concurrency)}
}

// Schedule
// handler Schedule before init Task
func (t Task) Schedule(task func()) {
	t.lt <- task
	go func() {
		defer safe.Recovers(func() {
			<-t.lt
		})
		task()
	}()

}
