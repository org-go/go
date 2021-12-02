package threads

type Worker struct {
	job     func()
	workers int
}

// NewWorker init Worker
// job func
// workers int  lock Num group
func NewWorker(job func(), workers int) Worker {

	return Worker{job: job, workers: workers}

}

// Schedule handler sync.WaitGroup instruct
// if need sync.waitGroup func before init Worker struct
func (w Worker) Schedule() {

	g := NewGroup()
	for i := 0; i < w.workers; i++ {
		g.SafeRun(w.job)
	}
	g.Wait()

}

func GWorker(cfs ...func()) {

	group := NewGroup()
	for _, cf := range cfs {
		group.SafeRun(cf)
	}
	group.Wait()

}
