package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"

	kModels "github.com/immathanr/service-todo/models"
	pens "github.com/immathanr/service-todo/pens"
)

const (
	KeyTimeFormat = "2006010215"
)

type Response struct {
	Status string `json:"status"`
	Error  error  `json:"error"`
}

// CrashService interface which has all the methods
type MetricsService interface {
	Crash(kModels.Crash) Response
	Anr(kModels.Crash) Response
}

type metricsService struct {
	fileWriter *pens.Pens
}

// ToString serializes Response model to json
func (r *Response) ToString() string {
	output, _ := json.Marshal(r)
	return string(output)
}

func (ms metricsService) Crash(crash kModels.Crash) Response {
	ms.fileWriter.Write("CRASH", crash.ToString())
	err := validator(crash)
	if err != nil {
		fmt.Println("Params not available: ", err.Error())
		return Response{"Ok", err}
	}
	crash.CreatedAt = time.Now()

	return Response{"Ok", nil}
}

func (ms metricsService) Anr(crash kModels.Crash) Response {
	ms.fileWriter.Write("ANR", crash.ToString())
	err := validator(crash)
	if err != nil {
		fmt.Println("Params not available: ", err.Error())
		return Response{"Ok", err}
	}
	crash.CreatedAt = time.Now()

	return Response{"Ok", nil}
}

func validator(crash kModels.Crash) error {
	var err error
	if crash.GameID == "" {
		err = errors.New("Game id not available")
	} else if crash.SDKV == "" {
		err = errors.New("SDK version not available")
	} else if crash.SDKN == "" {
		err = errors.New("SDK number not available")
	} else if crash.Platform == "" {
		err = errors.New("Platform not available")
	}
	return err
}
