package api

import (
	"fmt"
	"time"

	"github.com/go-kit/kit/metrics"
	kModels "github.com/immathanr/service-todo/models"
)

type instrumentingMiddleware struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	countResult    metrics.Histogram
	next           MetricsService
}

// NewInstrumentingService returns an instance of an instrumenting Service.
func NewInstrumentingService(
	counter metrics.Counter,
	successCounter metrics.Histogram,
	latency metrics.Histogram,
	s MetricsService) MetricsService {
	return &instrumentingMiddleware{
		requestCount:   counter,
		countResult:    successCounter,
		requestLatency: latency,
		next:           s,
	}
}

func (mw instrumentingMiddleware) Crash(crash kModels.Crash) (response Response) {
	defer func(begin time.Time) {
		lvs := []string{"method", "crash", "error", fmt.Sprint(response.Error != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	response = mw.next.Crash(crash)
	return
}

func (mw instrumentingMiddleware) Anr(crash kModels.Crash) (response Response) {
	defer func(begin time.Time) {
		lvs := []string{"method", "anr", "error", fmt.Sprint(response.Error != nil)}
		mw.requestCount.With(lvs...).Add(1)
		mw.requestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	response = mw.next.Anr(crash)
	return
}
