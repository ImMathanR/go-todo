package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	api "github.com/immathanr/service-todo/api"
	"github.com/immathanr/service-todo/counter"
	"github.com/immathanr/service-todo/pens"
	//"bitbucket.org/greedygame/slice-north-star/northstar"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

// func init() {
// 	northstar.Init()
// }

const (
	REDIS_COUNTER_ADDR = "0.0.0.0:6379"
	REDIS_COUNTER_DB   = 3
)

func main() {
	fileWriter := pens.New()
	defer fileWriter.Close()

	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	redisCounterConfig := counter.NewCounterConfig()
	redisCounterConfig.Addr = REDIS_COUNTER_ADDR
	redisCounterConfig.DB = REDIS_COUNTER_DB
	redisCounterConfig.SweepTime = time.Minute * 1
	redisCounterCtx, _ := counter.New(redisCounterConfig)
	defer redisCounterCtx.Stop()

	var svc api.MetricsService
	svc = api.NewService(redisCounterCtx, fileWriter)
	svc = api.NewInstrumentingService(requestCount, countResult, requestLatency, svc)

	mux := http.NewServeMux()
	mux.Handle("/", api.MakeHandler(svc))
	mux.Handle("/metrics", promhttp.Handler())
	fmt.Println("Listening on 8080")
	log.Fatalln(http.ListenAndServe(":8080", mux))
}
