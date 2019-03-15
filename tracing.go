package huaweiapm

import (
	"fmt"
	"github.com/go-mesh/openlogging"
	"github.com/openzipkin-contrib/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	"sync"
	"time"
)

type TracingReporter struct {
	batch          []*zipkincore.Span
	batchMutex     *sync.Mutex
	spanC          chan *zipkincore.Span
	batchInterval  string
	batchSize      int
	tracingRunning bool
}

func NewTracingReporter(batchInterval string, batchSize int) *TracingReporter {
	if batchInterval == "" {
		batchInterval = DefaultTracingBatchInterval
	}
	if batchSize == 0 {
		batchSize = DefaultTracingBatchSize
	}
	tr = &TracingReporter{
		batch:         make([]*zipkincore.Span, 0),
		batchMutex:    &sync.Mutex{},
		spanC:         make(chan *zipkincore.Span),
		batchSize:     batchSize,
		batchInterval: batchInterval,
	}
	return tr
}
func (tr *TracingReporter) appendSpan(span *zipkincore.Span) int {
	tr.batchMutex.Lock()
	defer tr.batchMutex.Unlock()
	tr.batch = append(tr.batch, span)
	return len(tr.batch)
}
func (tr *TracingReporter) doReport() {
	tr.batchMutex.Lock()
	defer tr.batchMutex.Unlock()
	err := client.ReportTracing(tr.batch)
	if err != nil {
		openlogging.Error("can not report tracing: " + err.Error())
		return
	}
	openlogging.Debug(fmt.Sprintf("report %d spans", len(tr.batch)))
	tr.batch = tr.batch[len(tr.batch):]

}
func (tr *TracingReporter) StartReportSpans() {
	t, _ := time.ParseDuration(tr.batchInterval)
	ticker := time.Tick(t)
	openlogging.Debug("start tracing")
	tr.tracingRunning = true
	for {

		select {
		case <-ticker:
			tr.doReport()
		case span := <-tr.spanC:
			size := tr.appendSpan(span)
			if size >= tr.batchSize {
				tr.doReport()
			}
		case stop := <-StopTracing:
			if stop {
				openlogging.Warn("stopped reporting spans, huawei apm tracing function lost")
				tr.tracingRunning = false
				break
			}
		}
	}

}
func (tr *TracingReporter) WriteSpan(span *zipkincore.Span) {
	if !tr.tracingRunning {
		return
	}
	tr.spanC <- span
}
