package api

import (
	"bufio"
	"fmt"
	"github.com/apache/thrift/lib/go/thrift"
	"github.com/go-chassis/huawei-apm/pkg/fifo"
	"github.com/go-chassis/huawei-apm/thrift/gen-go/apm"
	"github.com/go-mesh/openlogging"
	"github.com/openzipkin-contrib/zipkin-go-opentracing/thrift/gen-go/zipkincore"
	"sync"
)

var once sync.Once
var defaultService APM

//every API should be called after a certain interval
type APM interface {
	//ReportDiscoveryInfo send info to APM
	ReportDiscoveryInfo(info *apm.TDiscoveryInfo) error //one time per 5 min
	ReportKPI(message []*apm.TKpiMessage) error         //one time per 1 min, source service can be several, so it is a slice
	ReportTracing(span []*zipkincore.Span) error        //one time per 1 min
}

type DefaultAPM struct {
	writer *bufio.Writer
}

func GetAPMClient(app, serviceName string) (APM, error) {
	var err error
	once.Do(func() {
		defaultAPM := &DefaultAPM{}
		defaultAPM.writer, err = fifo.NewWriter(app, serviceName)
		if err != nil {
			openlogging.Error("can not create writer:" + err.Error())
		}
		defaultService = defaultAPM
	})
	return defaultService, err

}
func (da *DefaultAPM) ReportDiscoveryInfo(info *apm.TDiscoveryInfo) error {
	t := thrift.NewTMemoryBuffer()
	p := thrift.NewTBinaryProtocolTransport(t)
	if err := info.Write(p); err != nil {
		openlogging.Error("can not serialize discovery: " + err.Error())
		return err
	}
	n, err := da.writer.Write(t.Buffer.Bytes())
	if err != nil {
		openlogging.Error("can not report discovery: " + err.Error())
		return err
	}
	err = da.writer.Flush()
	if err != nil {
		openlogging.Error("can not flush discovery: " + err.Error())
		return err
	}
	openlogging.Debug(fmt.Sprintf("write inventory size %d to fifo", n))
	return nil
}
func (da *DefaultAPM) ReportKPI(messages []*apm.TKpiMessage) error {
	for _, m := range messages {
		t := thrift.NewTMemoryBuffer()
		p := thrift.NewTBinaryProtocolTransport(t)
		if err := m.Write(p); err != nil {
			openlogging.Error("can not serialize kpi: " + err.Error())
			return err
		}
		openlogging.Debug(t.String())
		_, err := da.writer.Write(t.Buffer.Bytes())
		if err != nil {
			openlogging.Error("can not report kpi: " + err.Error())
			return err
		}
		err = da.writer.Flush()
		if err != nil {
			openlogging.Error("can not flush kpi: " + err.Error())
			return err
		}
	}
	return nil
}
func (da *DefaultAPM) ReportTracing(spans []*zipkincore.Span) error {
	t := thrift.NewTMemoryBuffer()
	p := thrift.NewTBinaryProtocolTransport(t)
	if err := p.WriteListBegin(thrift.STRUCT, len(spans)); err != nil {
		openlogging.Error("can not serialize spans: " + err.Error())
		return err
	}
	for _, s := range spans {
		if err := s.Write(p); err != nil {
			openlogging.Error("can not serialize spans: " + err.Error())
			return err
		}
	}
	if err := p.WriteListEnd(); err != nil {
		openlogging.Error("can not serialize spans: " + err.Error())
		return err
	}
	openlogging.Debug(t.String())
	_, err := da.writer.Write(t.Buffer.Bytes())
	if err != nil {
		openlogging.Error("can not report spans: " + err.Error())
		return err
	}
	err = da.writer.Flush()
	if err != nil {
		openlogging.Error("can not flush spans: " + err.Error())
		return err
	}
	return nil
}
