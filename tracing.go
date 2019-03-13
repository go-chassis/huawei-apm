package huaweiapm

import (
	"github.com/go-mesh/openlogging"
	"github.com/openzipkin-contrib/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	"sync"
	"time"
)

var batch = make([]*zipkincore.Span, 0)
var batchMutex = &sync.Mutex{}
var spanC = make(chan *zipkincore.Span)

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
	batch = batch[len(batch):]
	openlogging.Debug("report tracing success")
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
				break
			}
		}
	}

}
func WriteSpan(span *zipkincore.Span) {
	spanC <- span
}
