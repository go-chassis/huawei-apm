package huaweiapm

import (
	"fmt"
	"github.com/go-mesh/openlogging"
	"github.com/openzipkin-contrib/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	"sync"
	"time"
)

var batch = make([]*zipkincore.Span, 0)
var batchMutex = &sync.Mutex{}
var spanC = make(chan *zipkincore.Span)

var tracingRunning = true

func appendSpan(span *zipkincore.Span) int {
	batchMutex.Lock()
	defer batchMutex.Unlock()
	batch = append(batch, span)
	return len(batch)
}
func doReport() {
	batchMutex.Lock()
	defer batchMutex.Unlock()
	err := client.ReportTracing(batch)
	if err != nil {
		openlogging.Error("can not report tracing: " + err.Error())
		return
	}
	openlogging.Debug(fmt.Sprintf("report %d spans", len(batch)))
	batch = batch[len(batch):]

}
func startReportSpans() {
	t, _ := time.ParseDuration(opt.TracingBatchInterval)
	ticker := time.Tick(t)
	openlogging.Debug("start tracing")

	for {
		select {
		case <-ticker:
			doReport()
		case span := <-spanC:
			size := appendSpan(span)
			if size >= opt.TracingBatchSize {
				doReport()
			}
		case stop := <-StopTracing:
			if stop {
				openlogging.Info("tracing stopped")
				tracingRunning = false
				break
			}
		}
	}

}
func WriteSpan(span *zipkincore.Span) {
	if !tracingRunning {
		openlogging.Warn("lost span, huawei apm tracing is not running")
	}
	spanC <- span
}
