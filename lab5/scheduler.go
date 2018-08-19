package lab5

import (
	"time"
)

// The different scheduling policies for this lab:
// FIFO: Schedules and executes the jobs in the order they are given.
// 		 In other words: The first job in the list should
//		 be the first to be completed
// SJF: Schedules the jobs based on the estimated time for the task.
// 		As the name Shortest Job First suggests, the tasks with
//		the lowest estimate should run first.
// RR: Round robin scheduling executes a job for a fixed amount of time.
// 	   This time is known as the quantum. Once the given quantum has elapsed,
//	   the next job is scheduled with the same quantum.
const (
	FIFO = iota
	SJF
	RR
)

type Job struct {
	id        int
	task      func(time.Duration)
	start     time.Time
	estimated time.Duration
	scheduled time.Duration
	remaining time.Duration
}

type Result struct {
	job     Job
	latency time.Duration
}

type Scheduler struct {
	Jobs    chan Job
	Results chan Result
}

// NewJob creates a new job in a much simpler way by initializing all
// Duration types to the given estimate and adds the start time for the job.
func NewJob(id int, task func(time.Duration), estimate time.Duration) Job {
	if task == nil {
		task = time.Sleep
	}
	return Job{id, task, time.Now(), estimate, estimate, estimate}
}

//NewScheduler creates a new scheduler that should be able to handle at least 500 jobs
// at once.
func NewScheduler() *Scheduler {
	// TODO(student): Initialize scheduler. Should the channels be buffered or not? Why?
	return &Scheduler{}
}

// Schedule is responsible for scheduling the given jobs according to a policy.
// The jobs should be put on the channel in the order dictated by the policy.
func (s *Scheduler) Schedule(jobs []Job, policy int, quantum time.Duration) {

}

// CreateWorkerPool creates a given number of worker goroutines
func (s *Scheduler) CreateWorkerPool(nrOfWorkers int) {

}

// Responsible for executing all jobs that appear on the channel
// by making the goroutine sleep for the scheduled time
func (s *Scheduler) worker() {

}
