package http

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/SkyAPM/go2sky"
	"github.com/jsmzr/boot/tracer"
	agentv3 "skywalking.apache.org/repo/goapi/collect/language/agent/v3"
)

const componentIDGOHttpClient = 5005

type SkywalkingTransport struct {
	delegated http.RoundTripper
}

func (s *SkywalkingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	entryReq, ok := tracer.ThreadLocal.Get().(*http.Request)
	if !ok {
		return s.delegated.RoundTrip(req)
	}
	t := go2sky.GetGlobalTracer()
	// get context by user
	span, err := t.CreateExitSpan(entryReq.Context(), fmt.Sprintf("/%s%s", req.Method, req.URL.Path), req.Host, func(headerKey, headerValue string) error {
		req.Header.Set(headerKey, headerValue)
		return nil
	})
	if err != nil {
		return s.delegated.RoundTrip(req)
	}
	defer span.End()
	span.SetComponent(componentIDGOHttpClient)
	span.Tag(go2sky.TagHTTPMethod, req.Method)
	span.Tag(go2sky.TagURL, req.URL.String())
	span.SetSpanLayer(agentv3.SpanLayer_Http)
	res, err := s.delegated.RoundTrip(req)
	if err != nil {
		span.Error(time.Now(), err.Error())
		return res, err
	}
	span.Tag(go2sky.TagStatusCode, strconv.Itoa(res.StatusCode))
	if res.StatusCode >= http.StatusBadRequest {
		span.Error(time.Now(), "Errors on handling client")
	}
	return res, nil
}

func init() {
	http.DefaultTransport = &SkywalkingTransport{http.DefaultTransport}
}
