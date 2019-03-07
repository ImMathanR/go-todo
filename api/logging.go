package api

import (
	"time"

	"github.com/go-kit/kit/log"
	kModels "github.com/immathanr/service-todo/models"
)

type loggingMiddleWare struct {
	logger log.Logger
	next   MetricsService
}

// NewLoggingService returns a new instance of a logging Service.
func NewLoggingService(logger log.Logger, s MetricsService) MetricsService {

	return &loggingMiddleWare{logger, s}
}

func (mw loggingMiddleWare) Crash(crash kModels.Crash) (output Response) {
	err := ""
	if output.Error != nil {
		err = output.Error.Error()
	}
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "impression",
			"input", crash.ToString(),
			"output", output.ToString(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output = mw.next.Crash(crash)
	return
}

func (mw loggingMiddleWare) Anr(crash kModels.Crash) (output Response) {
	err := ""
	if output.Error != nil {
		err = output.Error.Error()
	}
	defer func(begin time.Time) {
		_ = mw.logger.Log(
			"method", "click",
			"input", crash.ToString(),
			"output", output.ToString(),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	output = mw.next.Anr(crash)
	return
}
