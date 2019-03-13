package huaweiapm_test

import (
	"github.com/go-chassis/huawei-apm"
	"github.com/openzipkin-contrib/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestWriteSpan(t *testing.T) {
	os.Setenv("PAAS_POD_ID", "1")
	err := huaweiapm.Start(huaweiapm.Options{
		TracingBatchInterval: "2s",
		MonitoringGroup:      "app",
		ServiceName:          "service",
		ServiceType:          "go-chassis",
	})
	t.Log(err)
	assert.NoError(t, err)

	span := &zipkincore.Span{}
	t.Log(span)
	huaweiapm.WriteSpan(span)
}
