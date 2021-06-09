package timewheel

// GoPool goroutine pool
type GoPool interface {
	Go(func())
}

// Job 定时任务接口
type Job interface {
	Run()
}

// JobFunc job function
type JobFunc func()

// Run implement Job interface
func (sf JobFunc) Run() {
	sf()
}

func wrapJob(job Job) {
	defer func() {
		_ = recover()
	}()

	job.Run()
}

type emptyJob struct{}

func (emptyJob) Run() {}
