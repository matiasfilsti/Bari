package registerdurhelper

import (
	"fmt"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

//type SiteResponses struct {
//	Site     string
//	SiteTime float64
//}

var requestDurations = prometheus.NewHistogramVec(prometheus.HistogramOpts{
	Name: "http_request_duration_seconds",
	Help: "A histogram of the HTTP request durations in seconds.",
	// Bucket configuration: the first bucket includes all requests finishing in 0.05 seconds, the last one includes all requests finishing in 10 seconds.
	//Buckets: []float64{0.005, 0.01, 0.025, 0.05, 0.1, 0.25, 0.5, 1, 1.5, 2, 3, 4, 5},
	Buckets: []float64{100.00, 200.00, 500.00, 800.00, 1000.00, 1200.00, 1400.00, 1600.00, 2000.00, 2300.00, 2600.00, 3000.00, 4000.00, 8000.00},
}, []string{"env"})

// function Register used in main function
func Register(sites []string) {

	prometheus.MustRegister(requestDurations)
	for {
		for _, site := range sites {
			go metricUpdate(site)
		}
		time.Sleep(30 * time.Second)
	}
}

func metricUpdate(site string) {
	tr := &http.Transport{
		MaxIdleConns:       1,
		IdleConnTimeout:    15 * time.Second,
		DisableCompression: true,
	}
	client := &http.Client{Transport: tr}
	startime := time.Now()
	resp, err := client.Get(site)
	if err != nil {
		fmt.Printf("Site: %v ERROR \n", site)
	} else {
		defer resp.Body.Close()
		responseTime := float64(time.Since(startime).Milliseconds())
		fmt.Printf("Site: %v -- ResponseTime: %v ms -- StatusCode: %v \n", site, responseTime, resp.StatusCode)
		requestDurations.WithLabelValues(fmt.Sprint(site)).Observe(responseTime)
	}
}
