// Package analytics handles metric collection and processing.
package analytics

import (
	"time"
)

// Collector handles analytic collection.
type Collector interface {
	MeasureDuration(metric string, duration time.Duration)
}

// MeasureDuration is a helper function to use with defer that will measure elapsed time.
func MeasureDuration(obs Collector, metric string) func() {
	start := time.Now()

	return func() {
		elapsed := time.Since(start)

		obs.MeasureDuration(metric, elapsed)
	}
}
