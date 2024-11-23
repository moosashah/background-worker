package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Job struct {
	ID     int
	Status string
}

var jobQueue = make(chan Job, 100)
var jobID = 1
var workerPool = 5
var mu sync.Mutex

func startWorker() {
	for job := range jobQueue {
		processJob(job)
	}
}

func processJob(job Job) {
	time.Sleep(5 * time.Second)
	fmt.Printf("Job ID %d processed: %s\n", job.ID, job.Status)
}

// HTTP handler function to enqueue job
func enqueueJobHandler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	job := Job{ID: jobID, Status: "Queued"}
	jobID++
	mu.Unlock()

	select {
	case jobQueue <- job:
		fmt.Fprintf(w, "Job ID %d enqueued\n", job.ID)
	default:
		fmt.Fprintf(w, "Job queue full. Job ID %d could not be enqueued\n", job.ID)
	}
}

// Function to check the number of jobs remaining in the queue
func remainingJobsHandler(w http.ResponseWriter, r *http.Request) {
	remainingJobs := len(jobQueue)
	fmt.Fprintf(w, "Jobs remaining in queue: %d\n", remainingJobs)
}

func startWorkerPool(workerCount int) {
	for i := 0; i < workerCount; i++ {
		go startWorker() // Start a worker goroutine
	}
}

func main() {
	// Start the background worker pool with desired number of simultaneous workers
	startWorkerPool(workerPool)

	// Set up HTTP server
	http.HandleFunc("/enqueue-job", enqueueJobHandler)
	http.HandleFunc("/remaining", remainingJobsHandler)

	// Start the server
	fmt.Printf("Starting server on port 8080\n")
	http.ListenAndServe(":8080", nil)
}
