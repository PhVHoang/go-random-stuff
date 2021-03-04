var (
	MaxWorkers = os.Getenv("MAX_WORKERS")
	MaxQueue   = os.Getenv("MAX_QUEUE")
)

// Jog represents the job to be run
type Job struct {
	Payload Payload
}

// A buffered channel that we can send work requests on.
var JobQueue chan Job

// Worker represents the worker that executes the job
type Worker struct {
	WorkerPool chan chan Job
	JobChannel chan Job
	quit       chan bool
}

func NewWorker(workerPool chan chan Job) Worker {
	return Worker{
		WorkerPool: workerPool,
		JobChannel: make(chan Job),
		quit: make(chan bool)
	}
}

// Start method starts the run loop for the worker, listening for a quit channel 
// in case we need to stop it
func (w Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobChannel
			select {
			case job := <-w.JobChannel:
				if err := job.Payload.UploadToS3(); err != nil {
					log.Errorf("Error uploading to S3: %s", err.Error())
				}
			case <-w.quit:
				// we have received a signal to stop
				return
			}
		}
	}()
}

// Stop signals the worker to stop listening for work requests
func (w Worker) Stop() {
	go func() {
		w.quit <- true
	}()
}