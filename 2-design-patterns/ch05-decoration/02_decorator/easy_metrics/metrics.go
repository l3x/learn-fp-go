package easy_metrics

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"github.com/admobi/easy-metrics"
	"decorator"
)

var (
	avgResponseTime = metrics.NewGauge("avgResponseTime")
	requests        = metrics.NewCounter("requests")
	responseTime    = &timing{}
)

func Serve(addr string) error {
	r, err := metrics.NewTrackRegistry("Stats", 100, time.Second, false)
	if err != nil {
		decorator.Error.Println(err)
	}

	err = r.AddMetrics(requests, avgResponseTime)
	if err != nil {
		decorator.Error.Println(err)
	}

	http.HandleFunc("/", handler)
	return http.ListenAndServe(addr, nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	begin := time.Now()

	work()

	responseTime.Observe(time.Since(begin))
	requests.Inc()
}

func DisplayResults(addr string) error {
	decorator.Info.Printf("Go to http://%s/easy-metrics?show=Stats", addr)
	return nil
}

type timing struct {
	count int64
	sum   time.Duration
}

func (t *timing) Observe(d time.Duration) {
	t.count++
	t.sum += d
	avgResponseTime.Set(t.sum.Seconds() / float64(t.count))
}

func (t timing) String() string {
	avg := time.Duration(t.sum.Nanoseconds() / t.count)
	return fmt.Sprintf("\"%v\"", avg)
}


func work() {
	randInt := rand.Intn(5000)
	decorator.Debug.Printf("- randInt: %v", randInt)

	workTime := time.Duration(randInt) * time.Millisecond
	time.Sleep(workTime)
}
