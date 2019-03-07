package api

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/go-kit/kit/endpoint"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	kModels "github.com/immathanr/service-todo/models"
)

func makeCrashEndPoint(msv MetricsService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(kModels.Crash)
		res := msv.Crash(req)
		return metricsResponse{res.Status}, nil
	}
}

func makeAnrEndPoint(msv MetricsService) endpoint.Endpoint {
	return func(_ context.Context, request interface{}) (interface{}, error) {
		req := request.(kModels.Crash)
		res := msv.Anr(req)
		return metricsResponse{res.Status}, nil
	}
}

func decodeCrashRequest(_ context.Context, r *http.Request) (interface{}, error) {
	decoder := json.NewDecoder(r.Body)
	var s statsRequest
	err := decoder.Decode(&s)

	if err != nil {
		panic(err)
	}

	var nonFatal = false
	if strings.HasPrefix(r.URL.EscapedPath(), "/anr") {
		nonFatal = true
	}

	header := &r.Header

	request := kModels.Crash{
		GameID:         header.Get("game_id"),
		AdvID:          header.Get("advid"),
		AI5:            header.Get("ai5"),
		SDKV:           header.Get("sdkv"),
		SDKN:           header.Get("sdkn"),
		AppN:           s.AppN,
		AppV:           s.AppV,
		AndroidVersion: s.AndroidVersion,
		CampaignID:     s.CampaignID,
		Model:          s.Model,
		Stacktrace:     s.Stacktrace,
		Platform:       s.Platform,
		CrashTs:        s.CrashTs,
		Tag:            s.Tag,
		NonFatal:       nonFatal,
		CreatedAt:      time.Now(),
		IP:             r.Header.Get("X-Forwarded-For"),
	}

	// // Find the countryID using northstar module
	// geoInfo := resolveGeo("", "", request.IP)
	// if geoInfo == nil {
	// 	fmt.Println("could not find out geo location for ip: ", request.IP)
	// } else {
	// 	request.CountryID = geoInfo.CountryId
	// }

	fmt.Println("Request: ", request.ToString())
	return request, nil
}

func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(w).Encode(response)
}

type metricsResponse struct {
	Status string `json:"status"`
}

type statsRequest struct {
	AppN           string `json:"appn"`
	AppV           string `json:"appv"`
	AndroidVersion string `json:"android_v"`
	CampaignID     string `json:"campaign_id"`
	Model          string `json:"model"`
	Stacktrace     string `json:"stacktrace"`
	Platform       string `json:"platform"`
	CrashTs        string `json:"crash_ts"`
	Tag            string `json:"tag"`
	NonFatal       bool   `json:"non_fatal"`
}

func MakeHandler(svc MetricsService) http.Handler {
	crashHandler := httptransport.NewServer(
		makeCrashEndPoint(svc),
		decodeCrashRequest,
		encodeResponse,
	)

	anrHandler := httptransport.NewServer(
		makeAnrEndPoint(svc),
		decodeCrashRequest,
		encodeResponse,
	)

	r := mux.NewRouter()

	r.Methods("POST").
		Path("/crash/stats").
		Handler(crashHandler)

	r.Methods("POST").
		Path("/anr/stats").
		Handler(anrHandler)

	return r
}

func isEmptyStr(s string) bool {
	if s == "" {
		return true
	}
	return false
}
