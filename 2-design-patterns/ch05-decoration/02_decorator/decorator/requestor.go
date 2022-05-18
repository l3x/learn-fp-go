package decorator

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type response struct {
	duration      time.Duration
	err           error
}

type Job struct {
	Client       Client
	NumRequests  int
	Request      *http.Request
	IntervalSecs int
	responseChan chan *response
}

func (b *Job) displayProgress(stopChan chan struct{}) {
	var prevResponseCount int
	for {
		select {
		case <-time.Tick(time.Millisecond * 500):
			responseCount := len(b.responseChan)
			if prevResponseCount < responseCount {
				prevResponseCount = responseCount
				Debug.Printf("> %d requests done.", responseCount)
			}
		case <-stopChan:
			return
		}
	}
}

func (j *Job) Run() {
	j.responseChan = make(chan *response, j.NumRequests)
	stopChan := make(chan struct{})
	go j.displayProgress(stopChan)

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-interruptChan
		stopChan <- struct{}{}
		close(j.responseChan)
		os.Exit(130)
	}()

	var wg sync.WaitGroup
	intervalSecs := time.Duration(j.IntervalSecs)
	requestsPerformed := 0
	for range time.Tick(intervalSecs * time.Second)  {
		wg.Add(1)
		go func() {
			client := j.Client
			j.makeRequest(client)
			wg.Done()
		}()
		requestsPerformed++
		if requestsPerformed >= j.NumRequests {
			break
		}
	}
	wg.Wait()
	stopChan <- struct{}{}
	Debug.Printf("All requests done.")
	close(j.responseChan)
}

func (j *Job) makeRequest(c Client) {
	Debug.Printf("makeRequest: ")
	start := time.Now()
	resp, err := c.Do(j.Request)
	if err == nil {
		io.Copy(ioutil.Discard, resp.Body)
		resp.Body.Close()
	}
	t := time.Now()
	finish := t.Sub(start)
	j.responseChan <- &response{
		duration:   finish,
		err:        err,
	}
}
