package monitor

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/prometheus/client_golang/prometheus/promhttp"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
)

type RequestCounterUrlLabelMappingFunc func(c *gin.Context) string

type (
	Metric struct {
		MetricCollect prometheus.Collector
		Id            string
		Name          string
		Description   string
		Type          string
		Args          []string
	}

	Prometheus struct {
		requestCntVec             *prometheus.CounterVec
		requestDuration           *prometheus.HistogramVec
		requestSize               prometheus.Summary
		responseSize              prometheus.Summary
		router                    *gin.Engine
		listenAddress             string
		PromPg                    PromPushGateway
		Metrics                   []*Metric
		MetricsPath               string
		ReqCntUrlLabelMappingFunc RequestCounterUrlLabelMappingFunc

		UrlLabelFromContext string
	}

	PromPushGateway struct {
		PushIntervalSeconds time.Duration
		PushGatewayUrl      string
		MetricUrl           string
		JobName             string
	}
)

var (
	defaultMetricPath = "/metrics"

	requestCnt = &Metric{
		Id:          "requestCnt",
		Name:        "http_request_total",
		Description: "How many HTTP  requests processed, partitioned by status code and HTTP method",
		Type:        "counter_vec",
		Args:        []string{"code", "method", "handler", "host", "url"},
	}

	requestDuration = &Metric{
		Id:          "requestDuration",
		Name:        "http_request_duration_seconds",
		Description: "The HTTP request latencies inn seconds",
		Type:        "histogram_vec",
		Args:        []string{"code", "method", "url"},
	}

	requestSize = &Metric{
		Id:          "requestSize",
		Name:        "http_request_size_bytes",
		Description: "The HTTP request sizes in bytes",
		Type:        "summary",
	}

	responseSize = &Metric{
		Id:          "responseSize",
		Name:        "http_response_size_bytes",
		Description: "The HTTP response sizes in bytes",
		Type:        "summary",
	}

	standardMetrics = []*Metric{requestCnt, requestDuration, requestSize, responseSize}
)

func NewPrometheus(subsystem string, customMetrics ...[]*Metric) *Prometheus {
	if len(customMetrics) > 1 {
		panic("too many args. NewPrometheus( string, <optional []*Metric> ).")
	}

	var metrics []*Metric
	if len(customMetrics) == 1 {
		metrics = customMetrics[1]
	}

	metrics = append(metrics, standardMetrics...)

	p := &Prometheus{
		Metrics:     metrics,
		MetricsPath: defaultMetricPath,
		ReqCntUrlLabelMappingFunc: func(c *gin.Context) string {
			return c.Request.URL.Path
		},
	}

	return p
}

func (p *Prometheus) SetPushGateway(pushGatewayUrl, metricsUrl string, pushIntervalSeconds time.Duration) {
	p.PromPg.PushIntervalSeconds = pushIntervalSeconds
	p.PromPg.PushGatewayUrl = pushGatewayUrl
	p.PromPg.MetricUrl = metricsUrl
	p.startPushTicker()
}

func (p *Prometheus) SetPushGatewayJob(job string) {
	p.PromPg.JobName = job
}

func (p *Prometheus) SetListenAddr(addr string) {
	p.listenAddress = addr
	if len(p.listenAddress) != 0 {
		p.router = gin.Default()
	}
}

func (p *Prometheus) SetMetricsPath(e *gin.Engine) {
	if len(p.listenAddress) == 0 {
		e.GET(p.MetricsPath, prometheusHandler())
		return
	}

	p.router.GET(p.MetricsPath, prometheusHandler())
	p.runServer()
}

func (p *Prometheus) SetMetricsPathWithAuth(e *gin.Engine, accounts gin.Accounts) {
	if len(p.listenAddress) == 0 {
		e.GET(p.MetricsPath, gin.BasicAuth(accounts), prometheusHandler())
		return
	}

	p.router.GET(p.MetricsPath, gin.BasicAuth(accounts), prometheusHandler())
	p.runServer()
}

func (p *Prometheus) runServer() {
	if len(p.listenAddress) != 0 {
		go p.router.Run(p.listenAddress)
	}
}

func (p *Prometheus) getPushGatewayUrl() string {
	host, _ := os.Hostname()

	if len(p.PromPg.JobName) == 0 {
		p.PromPg.JobName = "gin"
	}

	return p.PromPg.PushGatewayUrl + "/metrics/job/" + p.PromPg.JobName + "/instance/" + host
}

func (p *Prometheus) getMetrics() []byte {
	response, _ := http.Get(p.PromPg.MetricUrl)
	defer response.Body.Close()

	b, _ := ioutil.ReadAll(response.Body)
	return b
}

func (p *Prometheus) startPushTicker() {
	t := time.NewTicker(time.Second * p.PromPg.PushIntervalSeconds)
	go func() {
		for range t.C {
			p.sedMetricsToPushGateway(p.getMetrics())
		}
	}()
}

func (p *Prometheus) sedMetricsToPushGateway(metrics []byte) {
	req, err := http.NewRequest("POST", p.getPushGatewayUrl(), bytes.NewBuffer(metrics))
	if err != nil {
		logrus.Errorln(err)
	}

	c := &http.Client{}
	if _, err := c.Do(req); err != nil {
		logrus.WithError(err).Errorln("Error sending to push gateway")
	}
}

func (p *Prometheus) registerMetrics(subsystem string) {
	for _, metric := range p.Metrics {
		m := NewMetric(metric, subsystem)
		if err := prometheus.Register(m); err != nil {
			logrus.WithError(err).Errorf("%s could not be registered in Prometheus", metric.Name)
		}

		switch metric {
		case requestCnt:
			p.requestCntVec = m.(*prometheus.CounterVec)
		case requestDuration:
			p.requestDuration = m.(*prometheus.HistogramVec)
		case requestSize:
			p.requestSize = m.(prometheus.Summary)
		case responseSize:
			p.responseSize = m.(prometheus.Summary)
		}

		metric.MetricCollect = m
	}
}

func (p *Prometheus) Use(e *gin.Engine) {
	e.Use(p.HandlerFunc())

}

func (p *Prometheus) HandlerFunc() gin.HandlerFunc {
	return func(context *gin.Context) {
		if context.Request.URL.Path == p.MetricsPath {
			context.Next()
			return
		}

		bt := time.Now()
		reqSize := computeApproximateRequestSize(context.Request)

		context.Next()

		status := strconv.Itoa(context.Writer.Status())
		elapsed := float64(time.Since(bt)) / float64(time.Second)
		resSize := context.Writer.Size()

		url := p.ReqCntUrlLabelMappingFunc(context)

		if len(p.UrlLabelFromContext) > 0 {
			u, found := context.Get(p.UrlLabelFromContext)
			if !found {
				u = "unknown"
			}

			url = u.(string)
		}

		p.requestDuration.WithLabelValues(status, context.Request.Method, url).Observe(elapsed)
		p.requestCntVec.WithLabelValues(status, context.Request.Method, context.HandlerName(), context.Request.Host, url).Inc()
		p.requestSize.Observe(float64(reqSize))
		p.responseSize.Observe(float64(resSize))
	}
}

func NewMetric(m *Metric, subsystem string) (collector prometheus.Collector) {
	opt := prometheus.Opts{
		Subsystem: subsystem,
		Name:      m.Name,
		Help:      m.Description,
	}

	switch m.Type {
	case "counter_vec":
		collector = prometheus.NewCounterVec(prometheus.CounterOpts(opt), m.Args)
	case "counter":
		collector = prometheus.NewCounter(prometheus.CounterOpts(opt))
	case "gauge_vec":
		collector = prometheus.NewGauge(prometheus.GaugeOpts(opt))
	case "gauge":
		collector = prometheus.NewGauge(prometheus.GaugeOpts(opt))
	case "histogram":
		collector = prometheus.NewHistogram(prometheus.HistogramOpts{Subsystem: subsystem, Name: m.Name, Help: m.Description})
	case "histogram_vec":
		collector = prometheus.NewHistogramVec(prometheus.HistogramOpts{Subsystem: subsystem, Name: m.Name, Help: m.Description}, m.Args)
	case "summary_vec":
		collector = prometheus.NewSummaryVec(prometheus.SummaryOpts{Subsystem: subsystem, Name: m.Name, Help: m.Description}, m.Args)
	case "summary":
		collector = prometheus.NewSummary(prometheus.SummaryOpts{Subsystem: subsystem, Name: m.Name, Help: m.Description})
	}

	return
}

func prometheusHandler() gin.HandlerFunc {
	h := promhttp.Handler()
	return func(context *gin.Context) {
		h.ServeHTTP(context.Writer, context.Request)
	}
}

func computeApproximateRequestSize(r *http.Request) (size int) {
	if r.URL != nil {
		size = len(r.URL.Path)
	}

	size += len(r.Method) + len(r.Proto)
	for name, values := range r.Header {
		size += len(name)
		for _, v := range values {
			size += len(v)
		}
	}
	size += len(r.Host)

	if r.ContentLength != -1 {
		size += int(r.ContentLength)
	}

	return
}
