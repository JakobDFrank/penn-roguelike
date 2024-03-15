package analytics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
	"sync"
	"time"
)

// Prometheus is an implementation of analytics.Collector using the Prometheus library.
type Prometheus struct {
	histograms sync.Map
}

// MeasureDuration will collect the given metric and duration into a histogram
func (p *Prometheus) MeasureDuration(metric string, duration time.Duration) {

	hist, found := p.histograms.Load(metric)

	if !found {
		newHist := promauto.NewHistogram(prometheus.HistogramOpts{
			Name:    metric,
			Buckets: prometheus.DefBuckets,
		})

		hist2, loaded := p.histograms.LoadOrStore(metric, newHist)

		if loaded {
			hist = newHist
		} else {
			hist = hist2
		}
	}

	histogram := hist.(prometheus.Histogram)

	histogram.Observe(duration.Seconds())
}

// NewPrometheus returns an instance of Prometheus.
func NewPrometheus() (*Prometheus, error) {
	http.Handle("/metrics", promhttp.Handler())

	go func() {
		if err := http.ListenAndServe(":2112", nil); err != nil {
			// JF: todo - panic in debug
			log.Fatal(err)
		}
	}()

	return &Prometheus{}, nil
}

var _ Collector = (*Prometheus)(nil)
