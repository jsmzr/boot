package logrus

import (
	"net/http"

	"github.com/SkyAPM/go2sky"
	"github.com/jsmzr/boot/tracer"
	"github.com/sirupsen/logrus"
)

type SkywalkingLogrusHook struct {
}

func (s *SkywalkingLogrusHook) Levels() []logrus.Level {
	return logrus.AllLevels
}

func (s *SkywalkingLogrusHook) Fire(e *logrus.Entry) error {
	req, ok := tracer.ThreadLocal.Get().(*http.Request)
	if ok && req != nil {
		e.Data["tid"] = go2sky.TraceID(req.Context())
	}
	return nil
}

func init() {
	logrus.AddHook(&SkywalkingLogrusHook{})
}
