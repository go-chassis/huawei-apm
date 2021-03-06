package huaweiapm_test

import (
	"github.com/go-chassis/huawei-apm"
	"github.com/openzipkin-contrib/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestWriteSpan(t *testing.T) {
	os.Setenv("PAAS_POD_ID", "1")
	err := huaweiapm.Start(huaweiapm.Options{
		MonitoringGroup: "app",
		ServiceName:     "service",
		ServiceType:     "go-chassis",
	})
	t.Log(err)
	assert.NoError(t, err)

	span := &zipkincore.Span{}
	t.Log(span)
	time.Sleep(1 * time.Second)
	r := huaweiapm.NewTracingReporter("1s", 1)
	go r.StartReportSpans()
	r.WriteSpan(span)
}
